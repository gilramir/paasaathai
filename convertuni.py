#!/usr/bin/python3

# Convert unicode.txt, from Unicode Thai code points,
# https://www.unicode.org/charts/PDF/U0E00.pdf
# to Golang constants

from argparse import ArgumentParser
import re
import sys

# input:
# 0E07 ง THAI CHARACTER NGO NGU
re_def = re.compile(r"^\s*(?P<hex>0E[0-9A-F]{2}) (?P<thai>\S+) (?P<desc>THAI"
    " (CHARACTER|DIGIT|CURRENCY) .*)$")

# output:
#	/* ฆ */ KHO_RAKHANG = 0x0E06

def main():
    parser = ArgumentParser()
    parser.add_argument("input_file")
    parser.add_argument("output_file", nargs="?")

    args = parser.parse_args()

    with open(args.input_file, "rt", encoding="utf-8") as rfh:
        if args.output_file:
            wfh = open(args.output_file, "wt")
        else:
            wfh = sys.stdout
        for line in rfh.readlines():
            m = re_def.match(line)
            if m:
                h = m.group("hex")
                t = m.group("thai")
                d = m.group("desc").replace(" ", "_").replace("-", "_")
                print(f"\t/* {t} */ {d} = 0x{h}", file=wfh)
            elif "THAI" in line:
                print("no match:", line)


if __name__ == "__main__":
    main()
