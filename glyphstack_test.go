package paasaathai

import (
	. "gopkg.in/check.v1"
)

const (
	// THAI_CHARACTER_SARA_E, THAI_CHARACTER_SARA_E, THAI_CHARACTER_SARA_II, THAI_CHARACTER_YO_YAK, THAI_CHARACTER_WO_WAE
	ThaiError1 = "เเียว"
)

// Find 3 individual Thai characters
func (s *MySuite) TestGlyphStack1(c *C) {

	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_CHO_CHAN
	// [used with another verb to indicate possibility] may; might
	input := "อาจ"
	c.Assert(len([]rune(input)), Equals, 3)

	gstacks := ParseGlyphStacks(input)
	c.Assert(len(gstacks), Equals, 3)

	c.Check(len(gstacks[0].Runes), Equals, 1)
	c.Check(len(gstacks[1].Runes), Equals, 1)
	c.Check(len(gstacks[2].Runes), Equals, 1)
}

// Find a mix of Latin and Thai characters
func (s *MySuite) TestGlyphStack2(c *C) {

	// a, THAI_CHARACTER_O_ANG, b, THAI_CHARACTER_SARA_AA, c, THAI_CHARACTER_CHO_CHAN
	input := "aอbาcจ"
	c.Assert(len([]rune(input)), Equals, 6)

	gstacks := ParseGlyphStacks(input)
	c.Assert(len(gstacks), Equals, 6)

	c.Check(len(gstacks[0].Runes), Equals, 1)
	c.Check(len(gstacks[1].Runes), Equals, 1)
	c.Check(len(gstacks[2].Runes), Equals, 1)
	c.Check(len(gstacks[3].Runes), Equals, 1)
	c.Check(len(gstacks[4].Runes), Equals, 1)
	c.Check(len(gstacks[5].Runes), Equals, 1)
}

// Handle a mistake in Thai orthography, but allowed in Unicode.
// It's a vowel above the front vowel which is wrong. It may not
// display correctly in your editor.
func (s *MySuite) TestGlyphStack3(c *C) {
	input := ThaiError1
	c.Assert(len([]rune(input)), Equals, 5)

	gstacks := ParseGlyphStacks(input)
	c.Assert(len(gstacks), Equals, 5)

	c.Check(len(gstacks[0].Runes), Equals, 1)
	c.Check(len(gstacks[1].Runes), Equals, 1)
	c.Check(len(gstacks[2].Runes), Equals, 1)
	c.Check(len(gstacks[3].Runes), Equals, 1)
	c.Check(len(gstacks[4].Runes), Equals, 1)

}

// Handle a 2-codepoint stack
func (s *MySuite) TestGlyphStack5(c *C) {
	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_MAI_THO, THAI_CHARACTER_SARA_AA,
	// THAI_CHARACTER_NGO_NGU
	// "to claim; claim that"
	input := "อ้าง"

	c.Assert(len([]rune(input)), Equals, 4)

	gstacks := ParseGlyphStacks(input)
	c.Assert(len(gstacks), Equals, 3)

	c.Check(len(gstacks[0].Runes), Equals, 2)
	c.Check(len(gstacks[1].Runes), Equals, 1)
	c.Check(len(gstacks[2].Runes), Equals, 1)

	c.Check(gstacks[0].Runes[0], Equals, THAI_CHARACTER_O_ANG)
	c.Check(gstacks[0].Runes[1], Equals, THAI_CHARACTER_MAI_THO)
}

// Handle a 3-codepoint stack
func (s *MySuite) TestGlyphStack6(c *C) {
	// This may not display properly in your editor
	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_YO_YAK, THAI_CHARACTER_SARA_UU,
	// THAI_CHARACTER_MAI_EK
	// "is (located at); to reside; to live (at); stay; exist at a
	// particular point in time"
	input := "อยู่"
	c.Assert(len([]rune(input)), Equals, 4)

	gstacks := ParseGlyphStacks(input)
	c.Assert(len(gstacks), Equals, 2)

	c.Check(len(gstacks[0].Runes), Equals, 1)
	c.Check(len(gstacks[1].Runes), Equals, 3)
}

// Handle a SARA_AM
func (s *MySuite) TestGlyphStack7(c *C) {
	// This may not display properly in your editor
	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_BO_BAIMAI,
	// THAI_CHARACTER_NO_NU, THAI_CHARACTER_MAI_THO, THAI_CHARACTER_SARA_AM

	// to bathe; take a bath; take a shower; swim
	input := "อาบน้ำ"

	c.Assert(len([]rune(input)), Equals, 6)

	gstacks := ParseGlyphStacks(input)
	c.Assert(len(gstacks), Equals, 5)

	c.Check(len(gstacks[0].Runes), Equals, 1)
	c.Check(len(gstacks[1].Runes), Equals, 1)
	c.Check(len(gstacks[2].Runes), Equals, 1)
	c.Check(len(gstacks[3].Runes), Equals, 2)
	c.Check(len(gstacks[4].Runes), Equals, 1)
}
