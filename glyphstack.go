package paasaathai

import (
	"fmt"
	"sync"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// Represents one vertical stack of characters that fit into one horizontal
// character box, according to Unicode glyphs. It is generic for Unicode, not
// Thai-specific. But, there are Thai-specific methods for investigating it.
type GlyphStack struct {
	Runes []rune
}

// Implements the fmt.Stringer interface
func (s *GlyphStack) String() string {
	return string(s.Runes)
}

func (s *GlyphStack) Repr() string {
	return fmt.Sprintf("<GlyphStack %s %s>", string(s.Runes), StringToRuneNames(string(s.Runes)))
}

func (s GlyphStack) IsThai() bool {
	return RuneIsThai(s.Runes[0])
}

/*
func (s GlyphStack) StartsWithConsonant() bool {
	return RuneIsConsonant(s.Runes[0])
}
*/

func (s GlyphStack) HasUpperPositionVowel() bool {
	for _, r := range s.Runes {
		if RuneIsUpperPositionVowel(r) {
			return true
		}
	}
	return false
}

type GlyphStackParser struct {
	Chan chan *GlyphStack
	Wg   sync.WaitGroup
}

func ParseGlyphStacks(input string) []*GlyphStack {
	var parser GlyphStackParser
	parser.GoParse(input)

	gstacks := make([]*GlyphStack, 0, len(input))
	for g := range parser.Chan {
		gstacks = append(gstacks, g)
	}

	parser.Wg.Wait()
	return gstacks
}

func (s *GlyphStackParser) GoParse(input string) {
	s.Chan = make(chan *GlyphStack)

	normalizedInput := norm.NFD.String(input)
	s.Wg.Add(1)
	go s.parse(normalizedInput)
}

func (s *GlyphStackParser) parse(input string) {
	defer close(s.Chan)
	defer s.Wg.Done()

	// Check the string (array of bytes for the UTF-8 encoding)
	for i := 0; i < len(input); {
		// The Unicode library notation of "boundaries" doesn't handle Thai
		// the way we need it to. Implement it ourselves.
		r1, r1sz := utf8.DecodeRuneInString(input[i:])

		// How many bytes have we decoded
		d := r1sz

		// Need 3 bytes to encode a Thai glyph in UTF-8; do we have
		// enough for another codepoint?
		if RuneIsConsonant(r1) && len(input)-(i+d) >= 3 {
			r2, r2sz := utf8.DecodeRuneInString(input[i+d:])
			if RuneIsUpperPosition(r2) || RuneIsLowerPositionVowel(r2) {
				d += r2sz
			}

			// An upper or lower vowel can still take a tone mark
			if (RuneIsLowerPositionVowel(r2) || RuneIsUpperPositionVowel(r2)) && len(input)-(i+d) >= 3 {
				r3, r3sz := utf8.DecodeRuneInString(input[i+d:])
				if RuneIsToneMark(r3) {
					d += r3sz
				}
			}
		}

		s.Chan <- &GlyphStack{
			Runes: []rune(input[i : i+d]),
		}
		i += d
	}
}
