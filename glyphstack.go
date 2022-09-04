package paasaathai

import (
	"fmt"
	"sync"

	"golang.org/x/text/unicode/norm"
)

// Represents one vertical stack of characters that fit into one horizontal
// character box, according to Unicode glyphs. It is generic for Unicode, not
// Thai-specific.
type GlyphStack struct {
	Runes []rune
}

// Implements the io.Reader interface
/*
func (s *GlyphStack) Read(p []byte) (n int, err error) {

}
*/

// Implements the fmt.Stringer interface
func (s *GlyphStack) String() string {
	return string(s.Runes)
}

func (s *GlyphStack) Repr() string {
	return fmt.Sprintf("<GlyphStack %s>", string(s.Runes))
}

type GlyphStackParser struct {
	GlyphChan chan *GlyphStack
	Wg        sync.WaitGroup
}

func ParseGlyphStacks(input string) []*GlyphStack {
	var parser GlyphStackParser
	parser.GoParse(input)

	//gstacks := make([]*GlyphStack, 0, len(input))
	gstacks := make([]*GlyphStack, 0)
	for g := range parser.GlyphChan {
		gstacks = append(gstacks, g)
	}

	parser.Wg.Wait()
	return gstacks
}

func (s *GlyphStackParser) GoParse(input string) {
	s.GlyphChan = make(chan *GlyphStack)

	normalizedInput := norm.NFD.String(input)
	s.Wg.Add(1)
	go s.parse(normalizedInput)
}

func (s *GlyphStackParser) parse(input string) {
	defer close(s.GlyphChan)
	defer s.Wg.Done()

	for i := 0; i < len(input); {
		d := norm.NFC.NextBoundaryInString(input[i:], true)
		s.GlyphChan <- &GlyphStack{
			Runes: []rune(input[i : i+d]),
		}
		i += d
	}
}

/*
// Feed runes into a goroutine that is accumulating runes
// to build up GlyphStacks. That in turn feeds into
// another goroutine to accumulate the GlyphStacks
func ParseGlyphStacks(input string) []*GlyphStack {

	var parser glyphStackParser

	parser.Init()

	gstacks := make([]*GlyphStack, 0)

	var gstack *GlyphStack

	runes := []rune(input)

	start := -1
	var i int
	for i = 0; i < len(runes); i++ {
	retryRune:
		r := runes[i]
		//fmt.Printf("[%d] (%d) is Thai: %v\n", i, r, RuneIsThai(r))
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
*/
