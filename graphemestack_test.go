package paasaathai

import (
	"fmt"

	. "gopkg.in/check.v1"
)

// Find 3 individual Thai characters
func (s *MySuite) TestGraphemeStack01(c *C) {

	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_CHO_CHAN
	// [used with another verb to indicate possibility] may; might
	input := "อาจ"
	c.Assert(len([]rune(input)), Equals, 3)

	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 3)

	c.Check(gstacks[0].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gstacks[1].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gstacks[2].Main, Equals, THAI_CHARACTER_CHO_CHAN)
}

// Find a mix of Latin and Thai characters
func (s *MySuite) TestGraphemeStack02(c *C) {

	// a, THAI_CHARACTER_O_ANG, b, THAI_CHARACTER_SARA_AA, c, THAI_CHARACTER_CHO_CHAN
	input := "aอbาcจ"
	c.Assert(len([]rune(input)), Equals, 6)

	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 6)

	c.Check(gstacks[0].Main, Equals, 'a')
	c.Check(gstacks[1].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gstacks[2].Main, Equals, 'b')
	c.Check(gstacks[3].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gstacks[4].Main, Equals, 'c')
	c.Check(gstacks[5].Main, Equals, THAI_CHARACTER_CHO_CHAN)
}

// Handle a mistake in Thai orthography, but allowed in Unicode.
// It's a vowel above the front vowel which is wrong. It may not
// display correctly in your editor.
func (s *MySuite) TestGraphemeStack03(c *C) {

	// THAI_CHARACTER_SARA_E, THAI_CHARACTER_SARA_E,
	// THAI_CHARACTER_SARA_II, THAI_CHARACTER_YO_YAK,
	// THAI_CHARACTER_WO_WAEN
	input := "เเียว"

	c.Assert(len([]rune(input)), Equals, 5)

	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 4)

	c.Check(gstacks[0].Main, Equals, THAI_CHARACTER_SARA_AE)
	c.Check(gstacks[1].DiacriticVowel, Equals, THAI_CHARACTER_SARA_II)
	c.Check(gstacks[2].Main, Equals, THAI_CHARACTER_YO_YAK)
	c.Check(gstacks[3].Main, Equals, THAI_CHARACTER_WO_WAEN)

}

// Handle a 2-codepoint stack
func (s *MySuite) TestGraphemeStack04(c *C) {
	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_MAI_THO, THAI_CHARACTER_SARA_AA,
	// THAI_CHARACTER_NGO_NGU
	// "to claim; claim that"
	input := "อ้าง"

	c.Assert(len([]rune(input)), Equals, 4)

	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 3)

	c.Check(gstacks[0].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gstacks[0].UpperDiacritic, Equals, THAI_CHARACTER_MAI_THO)
	c.Check(gstacks[1].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gstacks[2].Main, Equals, THAI_CHARACTER_NGO_NGU)
}

// Handle a 3-codepoint stack
func (s *MySuite) TestGraphemeStack05(c *C) {
	// This may not display properly in your editor
	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_YO_YAK, THAI_CHARACTER_SARA_UU,
	// THAI_CHARACTER_MAI_EK
	// "is (located at); to reside; to live (at); stay; exist at a
	// particular point in time"
	input := "อยู่"
	c.Assert(len([]rune(input)), Equals, 4)

	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 2)

	c.Check(gstacks[0].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gstacks[1].Main, Equals, THAI_CHARACTER_YO_YAK)
	c.Check(gstacks[1].DiacriticVowel, Equals, THAI_CHARACTER_SARA_UU)
	c.Check(gstacks[1].UpperDiacritic, Equals, THAI_CHARACTER_MAI_EK)
}

// Handle a SARA_AM
func (s *MySuite) TestGraphemeStack06(c *C) {
	// This may not display properly in your editor
	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_BO_BAIMAI,
	// THAI_CHARACTER_NO_NU, THAI_CHARACTER_MAI_THO, THAI_CHARACTER_SARA_AM

	// to bathe; take a bath; take a shower; swim
	input := "อาบน้ำ"

	c.Assert(len([]rune(input)), Equals, 6)

	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 5)

	c.Check(gstacks[0].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gstacks[1].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gstacks[2].Main, Equals, THAI_CHARACTER_BO_BAIMAI)
	c.Check(gstacks[3].Main, Equals, THAI_CHARACTER_NO_NU)
	c.Check(gstacks[3].UpperDiacritic, Equals, THAI_CHARACTER_MAI_THO)
	c.Check(gstacks[4].Main, Equals, THAI_CHARACTER_SARA_AM)

}

// Handle "ก็"
func (s *MySuite) TestGraphemeStack07(c *C) {
	// THAI_CHARACTER_KO_KAI, THAI_CHARACTER_MAITAIKHU
	input := "ก็"
	c.Assert(len([]rune(input)), Equals, 2)

	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 1)

	c.Check(gstacks[0].Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gstacks[0].UpperDiacritic, Equals, THAI_CHARACTER_MAITAIKHU)
}

// Handle "อึ"
func (s *MySuite) TestGraphemeStack08(c *C) {
	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_SARA_UE
	input := "อึ"
	c.Assert(len([]rune(input)), Equals, 2)

	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 1)

	c.Check(gstacks[0].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gstacks[0].DiacriticVowel, Equals, THAI_CHARACTER_SARA_UE)
}

// Handle sara e sara e -> sara ae correction
func (s *MySuite) TestGraphemeStack09(c *C) {
	input := string([]rune{THAI_CHARACTER_SARA_E, THAI_CHARACTER_SARA_E, THAI_CHARACTER_TO_TAO})
	c.Assert(len([]rune(input)), Equals, 3)

	gstacks := ParseGraphemeStacks(input)
	fmt.Printf("got: %+v\n", gstacks)
	c.Assert(len(gstacks), Equals, 2)

	c.Check(gstacks[0].Main, Equals, THAI_CHARACTER_SARA_AE)
	c.Check(gstacks[1].Main, Equals, THAI_CHARACTER_TO_TAO)
}
