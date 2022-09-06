package paasaathai

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestCCCase01(c *C) {
	// 6. "กรรม", [ gamM], deed; kamma; karma; sin; bad karma earned
	rnames := []rune{THAI_CHARACTER_KO_KAI, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_MO_MA}
	input := string(rnames)

	gstacks := ParseGlyphStacks(input)
	c.Assert(gstacks, HasLen, 4)
	/*
		ccs := ParseCharacterClustersFromGlyphStacks(gstacks)
		c.Assert(ccs, HasLen, 1)

		var cc *CharacterCluster
		cc = ccs[0]
		c.Check(cc.IsThai, Equals, true)
		c.Check(cc.IsValidThai, Equals, true)
		c.Check(cc.FirstConsonant.Runes[0], Equals, THAI_CHARACTER_KO_KAI)
		c.Check(cc.Tail, HasLen, 3)
		c.Check(cc.Tail[0].Runes[0], Equals, THAI_CHARACTER_RO_RUA)
		c.Check(cc.Tail[1].Runes[0], Equals, THAI_CHARACTER_RO_RUA)
		c.Check(cc.Tail[2].Runes[0], Equals, THAI_CHARACTER_MO_MA)
	*/
}

// input := "พิสูจนไดคะ"

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
