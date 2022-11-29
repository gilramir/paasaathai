package paasaathai

import (
	"errors"
	"fmt"
	"sync"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

var DiacriticVowelWithoutConsonantError = errors.New("An upper/lower diacritic vowel should be connected to a consonant")
var DiacriticWithoutConsonantError = errors.New("An upper diacritic should be connected to a consonant")

// For Thai code points, represents one vertical stack of Unicode codepoints
// that fit into one horizontal user-perceived character box.
// It can hold a single non-Thai code point instead of Thai code points.
// TODO - this is going to change over time as I decide how I want to store the
// info.
type GraphemeStack struct {
	// The original UTF-8
	Text string

	// This one is always set, for Thai or non-Thai
	Main rune

	// For Thai, these might be set
	DiacriticVowel rune
	UpperDiacritic rune
}

// Implements the fmt.Stringer interface
func (s GraphemeStack) String() string {
	return s.Text
}

func (s GraphemeStack) Repr() string {
	labels := ""
	if s.Main != 0 {
		labels += fmt.Sprintf("MAIN=%s", RuneToName(s.Main))
	}
	if s.DiacriticVowel != 0 {
		if labels != "" {
			labels += " "
		}
		labels += fmt.Sprintf("DV=%s", RuneToName(s.DiacriticVowel))
	}
	if s.UpperDiacritic != 0 {
		if labels != "" {
			labels += " "
		}
		labels += fmt.Sprintf("UD=%s", RuneToName(s.UpperDiacritic))
	}

	return fmt.Sprintf("<GraphemeStack %s %s>", s.Text, labels)
}

func (s GraphemeStack) IsThai() bool {
	return RuneIsThai(s.Main)
}

func (s GraphemeStack) IsValidThai() bool {
	return s.Main != 0 && RuneIsThai(s.Main)
}

/*
func (s GraphemeStack) StartsWithConsonant() bool {
	return RuneIsConsonant(s.Runes[0])
}

func (s GraphemeStack) HasUpperPositionVowel() bool {
	for _, r := range s.Runes {
		if RuneIsUpperPositionVowel(r) {
			return true
		}
	}
	return false
}
*/

func MustParseSingleGraphemeStack(input string) GraphemeStack {

	gstacks := ParseGraphemeStacks(input)
	if len(gstacks) != 1 {
		panic(fmt.Sprintf("The input string resulted in %d GraphemeStacks", len(gstacks)))
	}
	return gstacks[0]
}

type GraphemeStackParser struct {
	Chan chan GraphemeStack
	Wg   sync.WaitGroup
}

func ParseGraphemeStacks(input string) []GraphemeStack {
	var parser GraphemeStackParser
	parser.GoParse(input)

	gstacks := make([]GraphemeStack, 0, len(input))
	for g := range parser.Chan {
		gstacks = append(gstacks, g)
	}

	parser.Wg.Wait()
	return gstacks
}

func (s *GraphemeStackParser) GoParse(input string) {
	s.Chan = make(chan GraphemeStack)

	normalizedInput := norm.NFC.String(input)
	s.Wg.Add(1)
	go s.parse(normalizedInput)
}

// TODO
// should convert e e to ae
//33026. เเละ in data/best/news.zip(news/news_00038.txt) line 15 item 75
//	THAI_CHARACTER_SARA_E, THAI_CHARACTER_SARA_E, THAI_CHARACTER_LO_LING, THAI_CHARACTER_SARA_A

func (s *GraphemeStackParser) parse(input string) {
	defer close(s.Chan)
	defer s.Wg.Done()

	// Check the string (array of bytes for the UTF-8 encoding)
	for i := 0; i < len(input); {
		// The Unicode library notation of "boundaries" doesn't handle Thai
		// the way we need it to. Implement it ourselves.
		r1, r1sz := utf8.DecodeRuneInString(input[i:])

		// How many UTF-8 bytes have we decoded?
		decodedBytes := r1sz

		gs := GraphemeStack{}

		invalidThai := false
		if RuneIsUpperPositionVowel(r1) || RuneIsLowerPositionVowel(r1) {
			gs.DiacriticVowel = r1
			invalidThai = true
		} else if RuneIsToneMark(r1) || RuneIsUpperPositionSign(r1) {
			gs.UpperDiacritic = r1
			invalidThai = true
		} else {
			gs.Main = r1
		}

		// Not Thai? Next!
		if !RuneIsThai(r1) || invalidThai {
			gs.Text = input[i : i+decodedBytes]
			s.Chan <- gs
			i += decodedBytes
			continue
		}

		// If this Thai rune could have a diacritic on it, check.
		// A Thai code point needs 3 bytes to be encoded in UTF-8; do we have
		// enough for another code point?
		if RuneIsConsonant(r1) && len(input)-(i+decodedBytes) >= 3 {
			r2, r2sz := utf8.DecodeRuneInString(input[i+decodedBytes:])
			// Not Thai? Next!
			if !RuneIsThai(r2) {
				gs.Text = input[i : i+decodedBytes]
				s.Chan <- gs
				i += decodedBytes
				continue
			}

			if RuneIsUpperPositionVowel(r2) || RuneIsLowerPositionVowel(r2) {
				decodedBytes += r2sz
				gs.DiacriticVowel = r2
			} else if RuneIsToneMark(r2) || RuneIsUpperPositionSign(r2) {
				decodedBytes += r2sz
				gs.UpperDiacritic = r2
				// This GraphemeStack is only made of 2 code
				// points, because nothing can follow the
				// UpperDiacritic
				gs.Text = input[i : i+decodedBytes]
				s.Chan <- gs
				i += decodedBytes
				continue
			} else {
				// This GraphemeStack is only made of one code
				// point, because r2 cannot be stacked
				// over/under r1
				gs.Text = input[i : i+decodedBytes]
				s.Chan <- gs
				i += decodedBytes
				continue
			}

			// We have 2 code points. Is there a third?
			// An upper or lower vowel can still take a tone mark
			// or other upper diacritic
			if (RuneIsLowerPositionVowel(r2) || RuneIsUpperPositionVowel(r2)) && len(input)-(i+decodedBytes) >= 3 {
				r3, r3sz := utf8.DecodeRuneInString(input[i+decodedBytes:])
				if RuneIsToneMark(r3) || RuneIsUpperPositionSign(r3) {
					decodedBytes += r3sz
					gs.UpperDiacritic = r3
				}
			}
		} else if r1 == THAI_CHARACTER_SARA_E && len(input)-(i+decodedBytes) >= 3 {

			r2, r2sz := utf8.DecodeRuneInString(input[i+decodedBytes:])
			// Not Thai? Next!
			if !RuneIsThai(r2) {
				gs.Text = input[i : i+decodedBytes]
				s.Chan <- gs
				i += decodedBytes
				continue
			} else if r2 == THAI_CHARACTER_SARA_E {
				// correct spelling mistakes; 2 sara e's == 1 sara ae
				decodedBytes += r2sz
				gs.Main = THAI_CHARACTER_SARA_AE
				gs.Text = string(THAI_CHARACTER_SARA_AE)
				s.Chan <- gs
				i += decodedBytes
				continue
			} else {
				gs.Text = input[i : i+decodedBytes]
				s.Chan <- gs
				i += decodedBytes
				continue
			}
		}

		// At this point we have 2 or 3 code points.
		gs.Text = input[i : i+decodedBytes]
		s.Chan <- gs
		i += decodedBytes
	}
}
