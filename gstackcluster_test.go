package paasaathai

import (
	. "gopkg.in/check.v1"
)

/*
 The vowel orthograpy patterns from "Thai For Beginners" p. 243
 There are 12 patterns, #1 - #12, each with a short and long form
*/

func (s *MySuite) TestClusterTFBp243Short01a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "ละ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
}

func (s *MySuite) TestClusterTFBp243Short01b(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "สละ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_SO_SUA)

	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[1].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
}

func (s *MySuite) TestClusterTFBp243Long01a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "หา"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_HO_HIP)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AA)
}

func (s *MySuite) TestClusterTFBp243Long01b(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "อ่าน"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gcs[0].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_EK)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AA)

	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_NO_NU)
}

func (s *MySuite) TestClusterTFBp243Short02a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "สิบ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_SO_SUA)
	c.Check(gcs[0].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_I)

	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_BO_BAIMAI)
}

func (s *MySuite) TestClusterTFBp243Long02a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "หนี"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_HO_HIP)

	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_NO_NU)
	c.Check(gcs[1].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_II)
}

func (s *MySuite) TestClusterTFBp243Short03a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "ตึก"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_TO_TAO)
	c.Check(gcs[0].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_UE)

	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
}

func (s *MySuite) TestClusterTFBp243Long03a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "หรือ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_HO_HIP)

	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_RO_RUA)
	c.Check(gcs[1].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_UEE)
	c.Check(gcs[1].Tail[0].Main, Equals, THAI_CHARACTER_O_ANG)
}

func (s *MySuite) TestClusterTFBp243Short04a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "พุธ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_PHO_PHAN)
	c.Check(gcs[0].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_U)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_THO_THONG)
}

func (s *MySuite) TestClusterTFBp243Long04a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "หมู"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_HO_HIP)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_MO_MA)
	c.Check(gcs[1].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_UU)
}

/*
func (s *MySuite) TestClusterTFBp243Short06a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
input :=  "แหล่ะ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_SO_SUA)

	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[1].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
}
*/

/*
func (s *MySuite) TestCluster01(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()

	input := "เก่า"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gcs[0].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_EK)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AA)
}

func (s *MySuite) TestCluster02(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()

	input := "เบื่อ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_BO_BAIMAI)
	c.Check(gcs[0].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_UEE)
	c.Check(gcs[0].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_EK)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_O_ANG)
}
8?
//	input := "เมตตา"
/*
func (s *MySuite) TestCCCase01(c *C) {
	// 6. "กรรม", [ gamM], deed; kamma; karma; sin; bad karma earned
	rnames := []rune{THAI_CHARACTER_KO_KAI, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_MO_MA}
	input := string(rnames)

	gstacks := ParseGraphemeStacks(input)
	c.Assert(gstacks, HasLen, 4)
*/
/*
	ccs := ParseGStackClustersFromGraphemeStacks(gstacks)
	c.Assert(ccs, HasLen, 1)

	var cc *GStackCluster
	cc = ccs[0]
	c.Check(cc.IsThai, Equals, true)
	c.Check(cc.IsValidThai, Equals, true)
	c.Check(cc.FirstConsonant.Runes[0], Equals, THAI_CHARACTER_KO_KAI)
	c.Check(cc.Tail, HasLen, 3)
	c.Check(cc.Tail[0].Runes[0], Equals, THAI_CHARACTER_RO_RUA)
	c.Check(cc.Tail[1].Runes[0], Equals, THAI_CHARACTER_RO_RUA)
	c.Check(cc.Tail[2].Runes[0], Equals, THAI_CHARACTER_MO_MA)
*/
//}

// input := "พิสูจนไดคะ"

// "ก็"
/*
func (s *MySuite) TestCC01(c *C) {

	// THAI_CHARACTER_KO_KAI, THAI_CHARACTER_MAITAIKHU
	input := "ก็"
	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 1)

	ccs := ParseGStackClustersFromGraphemeStacks(gstacks)
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

	gstacks := ParseGraphemeStacks(input)
	c.Assert(len(gstacks), Equals, 1)

	ccs := ParseGStackClustersFromGraphemeStacks(gstacks)
	c.Assert(len(ccs), Equals, 1)

	cc := ccs[0]
	c.Check(cc.IsThai, Equals, true)
	c.Check(cc.IsValidThai, Equals, true)
	c.Check(len(cc.FirstConsonant.Runes), Equals, 2)
	c.Check(cc.FirstConsonant.HasUpperPositionVowel(), Equals, true)
}
*/
