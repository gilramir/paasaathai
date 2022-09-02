package paasaathai

import "fmt"

type GlyphStack struct {
	Text string

	// IsThai is true, IsValid is true:  all the Thai fields are filled in, as well as Text,
	// IsThai is true, IsValid is false: Text is filled in with 1 or more
	// Thai characters
	// IsThai is false, IsValid is always false, and Text is filld in with
	// no Thai characters:
	IsThai bool

	// Is this a valid GlyphStack, or just one that carries
	// an invalid Text
	IsValid bool

	// Tone
	Upper2Level rune

	// Vowel, Tone, or Karan
	Upper1Level rune

	// Vowel, consonant, or special
	MiddleLevel rune

	LowerVowel rune
}

// Implements the io.Reader interface
/*
func (s *GlyphStack) Read(p []byte) (n int, err error) {

}
*/

// Implements the fmt.Stringer interface
func (s *GlyphStack) String() string {
	return s.Text
}

func (s *GlyphStack) Repr() string {
	if s.IsThai {
		return fmt.Sprintf("(%s)", s.Text)
	} else {
		return fmt.Sprintf("(%+v)", s)
	}
}

func ParseGlyphStacks(input string) []*GlyphStack {

	gstacks := make([]*GlyphStack, 0)

	var gstack *GlyphStack

	runes := []rune(input)

	start := -1
	var i int
	for i = 0; i < len(runes); i++ {
	retryRune:
		r := runes[i]
		if gstack == nil {
			start = i
			gstack = &GlyphStack{
				IsThai: RuneIsThai(r),
			}
			continue
		} else if gstack.IsThai {
			// for now
			if RuneIsThai(r) {
				continue
			} else {
				gstack.Text = string(runes[start:i])
				gstacks = append(gstacks, gstack)
				gstack = nil
				goto retryRune
			}
		} else {
			// gstack is not Thai

			if RuneIsThai(r) {
				gstack.Text = string(runes[start:i])
				gstacks = append(gstacks, gstack)
				gstack = nil
				goto retryRune
			} else {
				// rune is not thai
				continue
			}
		}
	}
	if gstack != nil {
		gstack.Text = string(runes[start:i])
		gstacks = append(gstacks, gstack)
		gstack = nil
	}
	return gstacks
}
