package paasaathai

/*
import (
	. "gopkg.in/check.v1"
)
*/

// "ก็"
/*
func (s *MySuite) TestCC01(c *C) {

	// THAI_CHARACTER_KO_KAI, THAI_CHARACTER_MAITAIKHU
	input := "ก็"
	gstacks := ParseGlyphStacks(input)
	c.Assert(len(gstacks), Equals, 1)

	ccs := ParseCharacterClustersFromGlyphStacks(gstacks)
	c.Assert(len(ccs), Equals, 1)

	cc := ccs[0]
	c.Check(cc.IsThai, Equals, true)
	c.Check(cc.IsValidThai, Equals, true)
	c.Check(len(cc.FirstConsonant.Runes), Equals, 2)
	c.Check(cc.FirstConsonant.HasUpperPositionVowel(), Equals, true)
}

func (s *MySuite) TestCC02(c *C) {

	// THAI_CHARACTER_O_ANG, THAI_CHARACTER_SARA_UE
	input := "อึ"
	//fmt.Println(StringToRuneNames(input))

	gstacks := ParseGlyphStacks(input)
	c.Assert(len(gstacks), Equals, 1)

	ccs := ParseCharacterClustersFromGlyphStacks(gstacks)
	c.Assert(len(ccs), Equals, 1)

	cc := ccs[0]
	c.Check(cc.IsThai, Equals, true)
	c.Check(cc.IsValidThai, Equals, true)
	c.Check(len(cc.FirstConsonant.Runes), Equals, 2)
	c.Check(cc.FirstConsonant.HasUpperPositionVowel(), Equals, true)
}
*/
