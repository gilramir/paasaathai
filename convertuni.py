#!/usr/bin/python3

# Convert unicode.txt, from Unicode Thai code points,
# https://www.unicode.org/charts/PDF/U0E00.pdf
# to Golang constants

from argparse import ArgumentParser
from collections import namedtuple
import re
import sys

# input:
# 0E07 ง THAI CHARACTER NGO NGU
re_def = re.compile(r"^\s*(?P<hex>0E[0-9A-F]{2}) (?P<thai>\S+) (?P<desc>THAI"
    " (CHARACTER|DIGIT|CURRENCY) .*)$")

TCP = namedtuple("TCP", ("hex", "thai", "description"))

# output:
#	/* ฆ */ KHO_RAKHANG = 0x0E06

def main():
    parser = ArgumentParser()
    parser.add_argument("input_file")
    parser.add_argument("output_file", nargs="?")

    args = parser.parse_args()

    wfh = sys.stdout
    if args.output_file:
        wfh = open(args.output_file, "wt")

    print("""package paasaathai

// From Unicode Thai code points
// https://www.unicode.org/charts/PDF/U0E00.pdf
const (""", file=wfh)

    codepoints = []

    # The const section
    # /* ก */ THAI_CHARACTER_KO_KAI = rune(0x0E01)
    with open(args.input_file, "rt", encoding="utf-8") as rfh:
        for line in rfh.readlines():
            m = re_def.match(line)
            if m:
                h = m.group("hex")
                t = m.group("thai")
                d = m.group("desc").replace(" ", "_").replace("-", "_")
                d = d.rstrip("\n")
                print(f"\t/* {t} */ {d} = rune(0x{h})", file=wfh)
                cp = TCP(h, t, d)
                codepoints.append(cp)
            elif "THAI" in line:
                print("no match:", line)

    print(""")

var RuneToThaiName = map[rune]string{""", file=wfh)

    # The string map
    # 0x0E01:/* ก */ "THAI_CHARACTER_KO_KAI",
    for cp in codepoints:
        print(f"\t0x{cp.hex}:/* {cp.thai} */ \"{cp.description}\",",
            file=wfh)

    print("""}

var ThaiNameToRune = map[string]rune{""", file=wfh)

    # The string map
    # /* ก */ "THAI_CHARACTER_KO_KAI" : 0x0E01,
    for cp in codepoints:
        print(f"\t/* {cp.thai} */ \"{cp.description}\": 0x{cp.hex},",
            file=wfh)

    print("}", file=wfh)

if __name__ == "__main__":
    main()
