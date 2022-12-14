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
func (s *MySuite) TestClusterTFBp243Short05a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เตะ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_TO_TAO)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
}

func (s *MySuite) TestClusterTFBp243Long05a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เอง"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_NGO_NGU)
}

// LO LING by itself
func (s *MySuite) TestClusterTFBp243Short06a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "และ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_AE)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
}

// HO HIP and LO LING
func (s *MySuite) TestClusterTFBp243Short06b(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "แหล่ะ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_AE)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_HO_HIP)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[0].Tail[0].UpperDiacritic, Equals, THAI_CHARACTER_MAI_EK)
	c.Check(gcs[0].Tail[1].Main, Equals, THAI_CHARACTER_SARA_A)
}

// TO_TAO doesn't glide with LO LING
func (s *MySuite) TestClusterTFBp243Short06c(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "แต่ละ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_AE)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_TO_TAO)
	c.Check(gcs[0].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_EK)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[1].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
}

// PHO PHUNG and LO LING
func (s *MySuite) TestClusterTFBp243Short06d(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "ตบแผละ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 3)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_TO_TAO)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_BO_BAIMAI)
	c.Check(gcs[2].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_AE)
	c.Check(gcs[2].FirstConsonant.Main, Equals, THAI_CHARACTER_PHO_PHUNG)
	c.Check(gcs[2].Tail[0].Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[2].Tail[1].Main, Equals, THAI_CHARACTER_SARA_A)
}

// KO KAI and LO LING
func (s *MySuite) TestClusterTFBp243Short06e(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "แกละ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_AE)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[0].Tail[1].Main, Equals, THAI_CHARACTER_SARA_A)
}

func (s *MySuite) TestClusterTFBp243Long06a(c *C) {

	var gcp GStackClusterParser
	gcp.Initialize()
	input := "แก"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_AE)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
}

// PO PLA and LO LING
func (s *MySuite) TestClusterTFBp243Long06b(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "แปลง"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	/*
		fmt.Printf("0: %+v\n", gcs[0])
		fmt.Printf("1: %+v\n", gcs[1])
		fmt.Printf("2: %+v\n", gcs[2])
	*/
	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_AE)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_PO_PLA)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_NGO_NGU)
}

func (s *MySuite) TestClusterTFBp243Short07a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "โต๊ะ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	/*
		fmt.Printf("0: %+v\n", gcs[0])
		fmt.Printf("1: %+v\n", gcs[1])
		fmt.Printf("2: %+v\n", gcs[2])
	*/
	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_O)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_TO_TAO)
	c.Check(gcs[0].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_TRI)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
}

/*
???
func (s *MySuite) TestClusterTFBp243Short07b(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "โต๊ระ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	fmt.Printf("0: %+v\n", gcs[0])
	fmt.Printf("1: %+v\n", gcs[1])

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_O)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_TO_TAO)
	c.Check(gcs[0].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_TRI)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_RO_RUA)
	c.Check(gcs[0].Tail[1].Main, Equals, THAI_CHARACTER_SARA_A)
}
*/

/*
???
func (s *MySuite) TestClusterTFBp243Long07a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "โผล๊ะ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_O)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_PHO_PHUNG)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[1].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_TRI)
	c.Check(gcs[1].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
}
*/

func (s *MySuite) TestClusterTFBp243Short08a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เปาะ"

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_PO_PLA)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gcs[0].Tail[1].Main, Equals, THAI_CHARACTER_SARA_A)
}

func (s *MySuite) TestClusterTFBp243Long08a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "ก่อน"

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gcs[0].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_EK)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_NO_NU)
}

func (s *MySuite) TestClusterTFBp243Short09a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()

	input := "ลัวะ"

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[0].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_MAI_HAN_AKAT)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_WO_WAEN)
	c.Check(gcs[0].Tail[1].Main, Equals, THAI_CHARACTER_SARA_A)
}

func (s *MySuite) TestClusterTFBp243Long09a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "กลัว"

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[1].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_MAI_HAN_AKAT)
	c.Check(gcs[1].Tail[0].Main, Equals, THAI_CHARACTER_WO_WAEN)
}

func (s *MySuite) TestClusterTFBp243Short10a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เปรี๊ยะ"
	//THAI_CHARACTER_SARA_E, THAI_CHARACTER_PO_PLA, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_SARA_II, THAI_CHARACTER_MAI_TRI, THAI_CHARACTER_YO_YAK, THAI_CHARACTER_SARA_A
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_PO_PLA)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_RO_RUA)
	c.Check(gcs[0].Tail[0].DiacriticVowel, Equals, THAI_CHARACTER_SARA_II)
	c.Check(gcs[0].Tail[1].Main, Equals, THAI_CHARACTER_YO_YAK)
	c.Check(gcs[0].Tail[2].Main, Equals, THAI_CHARACTER_SARA_A)
}

func (s *MySuite) TestClusterTFBp243Long10a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เงียบ"

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_NGO_NGU)
	c.Check(gcs[0].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_II)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_YO_YAK)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_BO_BAIMAI)
}

/*
TODO - can't find an example
func (s *MySuite) TestClusterTFBp243Short11a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เมือง"
	//	THAI_CHARACTER_SARA_E, THAI_CHARACTER_MO_MA, THAI_CHARACTER_SARA_UEE, THAI_CHARACTER_O_ANG, THAI_CHARACTER_NGO_NGU

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_MO_MA)
	c.Check(gcs[0].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_UEE)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_NGO_NGU)
}
*/

func (s *MySuite) TestClusterTFBp243Long11a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เมือง"
	//	THAI_CHARACTER_SARA_E, THAI_CHARACTER_MO_MA, THAI_CHARACTER_SARA_UEE, THAI_CHARACTER_O_ANG, THAI_CHARACTER_NGO_NGU

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_MO_MA)
	c.Check(gcs[0].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_UEE)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_NGO_NGU)
}

func (s *MySuite) TestClusterTFBp243Short12a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เยอะ"
	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_YO_YAK)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gcs[0].Tail[1].Main, Equals, THAI_CHARACTER_SARA_A)
}

func (s *MySuite) TestClusterTFBp243Long12a(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เฟ้อ"

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_FO_FAN)
	c.Check(gcs[0].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_THO)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_O_ANG)
}

func (s *MySuite) TestClusterTFBp250MedialEr(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "เพลิง"
	//        THAI_CHARACTER_SARA_E, THAI_CHARACTER_PHO_PHAN, THAI_CHARACTER_LO_LING, THAI_CHARACTER_SARA_I, THAI_CHARACTER_NGO_NGU

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)

	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_PHO_PHAN)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[0].Tail[0].DiacriticVowel, Equals, THAI_CHARACTER_SARA_I)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_NGO_NGU)
}

func (s *MySuite) TestClusterSaraAm(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "กระนั้น"
	//		THAI_CHARACTER_KO_KAI, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_SARA_A, THAI_CHARACTER_NO_NU, THAI_CHARACTER_MAI_HAN_AKAT, THAI_CHARACTER_MAI_THO, THAI_CHARACTER_NO_NU

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 3)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_RO_RUA)
	c.Check(gcs[1].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
	c.Check(gcs[2].FirstConsonant.Main, Equals, THAI_CHARACTER_NO_NU)
	c.Check(gcs[2].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_MAI_HAN_AKAT)
	c.Check(gcs[2].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_MAI_THO)
	c.Check(gcs[2].Tail[0].Main, Equals, THAI_CHARACTER_NO_NU)

}

func (s *MySuite) TestClusterMaiHanAkat(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "กระทำ"
	//THAI_CHARACTER_KO_KAI, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_SARA_A, THAI_CHARACTER_THO_THAHAN, THAI_CHARACTER_SARA_AM

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 3)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_RO_RUA)
	c.Check(gcs[1].Tail[0].Main, Equals, THAI_CHARACTER_SARA_A)
	c.Check(gcs[2].FirstConsonant.Main, Equals, THAI_CHARACTER_THO_THAHAN)
	c.Check(gcs[2].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AM)

}

func (s *MySuite) TestClusterPunctuation(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "กรุงเทพฯ"
	//	THAI_CHARACTER_KO_KAI, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_SARA_U, THAI_CHARACTER_NGO_NGU, THAI_CHARACTER_SARA_E, THAI_CHARACTER_THO_THAHAN, THAI_CHARACTER_PHO_PHAN, THAI_CHARACTER_PAIYANNOI

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 6)

	c.Check(gcs[5].SingleMidSign.Main, Equals, THAI_CHARACTER_PAIYANNOI)
}

func (s *MySuite) TestClusterSaraAi(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "กำไร"
	//THAI_CHARACTER_KO_KAI, THAI_CHARACTER_SARA_AM, THAI_CHARACTER_SARA_AI_MAIMALAI, THAI_CHARACTER_RO_RUA

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 2)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AM)
	c.Check(gcs[1].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_AI_MAIMALAI)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_RO_RUA)
}

func (s *MySuite) TestClusterSaraOGlide(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()

	input := "โคคาโคลา"
	//	THAI_CHARACTER_SARA_O, THAI_CHARACTER_KHO_KHWAI, THAI_CHARACTER_KHO_KHWAI, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_SARA_O, THAI_CHARACTER_KHO_KHWAI, THAI_CHARACTER_LO_LING, THAI_CHARACTER_SARA_AA

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 4)
}

func (s *MySuite) TestClusterSandwichAo(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()

	input := "เกลา"
	//	THAI_CHARACTER_SARA_E, THAI_CHARACTER_KO_KAI, THAI_CHARACTER_LO_LING, THAI_CHARACTER_SARA_AA

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_LO_LING)
	c.Check(gcs[0].Tail[1].Main, Equals, THAI_CHARACTER_SARA_AA)
}

func (s *MySuite) TestClusterChaphaw(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()

	input := "เฉพาะ"
	//	THAI_CHARACTER_SARA_E, THAI_CHARACTER_CHO_CHING, THAI_CHARACTER_PHO_PHAN, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_SARA_A

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_CHO_CHING)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_PHO_PHAN)
	c.Check(gcs[0].Tail[1].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gcs[0].Tail[2].Main, Equals, THAI_CHARACTER_SARA_A)
}

func (s *MySuite) TestClusterOAngConsonant(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()

	input := "ราชอาณาจักรไทย"
	//	THAI_CHARACTER_RO_RUA, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_CHO_CHANG, THAI_CHARACTER_O_ANG, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_NO_NEN, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_CHO_CHAN, THAI_CHARACTER_MAI_HAN_AKAT, THAI_CHARACTER_KO_KAI, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_SARA_AI_MAIMALAI, THAI_CHARACTER_THO_THAHAN, THAI_CHARACTER_YO_YAK

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 8)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_RO_RUA)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_CHO_CHANG)
	c.Check(gcs[2].FirstConsonant.Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gcs[2].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gcs[3].FirstConsonant.Main, Equals, THAI_CHARACTER_NO_NEN)
	c.Check(gcs[3].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gcs[4].FirstConsonant.Main, Equals, THAI_CHARACTER_CHO_CHAN)
	c.Check(gcs[4].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_MAI_HAN_AKAT)
	c.Check(gcs[4].Tail[0].Main, Equals, THAI_CHARACTER_KO_KAI)
	c.Check(gcs[5].FirstConsonant.Main, Equals, THAI_CHARACTER_RO_RUA)
	c.Check(gcs[6].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_AI_MAIMALAI)
	c.Check(gcs[6].FirstConsonant.Main, Equals, THAI_CHARACTER_THO_THAHAN)
	c.Check(gcs[7].FirstConsonant.Main, Equals, THAI_CHARACTER_YO_YAK)
}

func (s *MySuite) TestClusterSaraEOAngVowel(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "พีเคอาร์อาร์พี"
	//	article/article_00072.txt line 113 -> THAI_CHARACTER_PHO_PHAN, THAI_CHARACTER_SARA_II, THAI_CHARACTER_SARA_E, THAI_CHARACTER_KHO_KHWAI, THAI_CHARACTER_O_ANG, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_THANTHAKHAT, THAI_CHARACTER_O_ANG, THAI_CHARACTER_SARA_AA, THAI_CHARACTER_RO_RUA, THAI_CHARACTER_THANTHAKHAT, THAI_CHARACTER_PHO_PHAN, THAI_CHARACTER_SARA_II

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 7)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_PHO_PHAN)
	c.Check(gcs[0].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_II)
	c.Check(gcs[1].FrontVowel.Main, Equals, THAI_CHARACTER_SARA_E)
	c.Check(gcs[1].FirstConsonant.Main, Equals, THAI_CHARACTER_KHO_KHWAI)
	c.Check(gcs[2].FirstConsonant.Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gcs[2].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gcs[3].FirstConsonant.Main, Equals, THAI_CHARACTER_RO_RUA)
	c.Check(gcs[3].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_THANTHAKHAT)
	c.Check(gcs[4].FirstConsonant.Main, Equals, THAI_CHARACTER_O_ANG)
	c.Check(gcs[4].Tail[0].Main, Equals, THAI_CHARACTER_SARA_AA)
	c.Check(gcs[5].FirstConsonant.Main, Equals, THAI_CHARACTER_RO_RUA)
	c.Check(gcs[5].FirstConsonant.UpperDiacritic, Equals, THAI_CHARACTER_THANTHAKHAT)
	c.Check(gcs[6].FirstConsonant.Main, Equals, THAI_CHARACTER_PHO_PHAN)
	c.Check(gcs[6].FirstConsonant.DiacriticVowel, Equals, THAI_CHARACTER_SARA_II)
}

func (s *MySuite) TestClusterSanskritRa(c *C) {
	var gcp GStackClusterParser
	gcp.Initialize()
	input := "ฤๅ"
	//        article/article_00177.txt line 2 -> THAI_CHARACTER_RU, THAI_CHARACTER_LAKKHANGYAO

	gs := ParseGraphemeStacks(input)
	gcs := gcp.ParseGraphemeStacks(gs)
	c.Assert(len(gcs), Equals, 1)

	c.Check(gcs[0].FirstConsonant.Main, Equals, THAI_CHARACTER_RU)
	c.Check(gcs[0].Tail[0].Main, Equals, THAI_CHARACTER_LAKKHANGYAO)
}

// TODO โต๊ระ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00061.txt) line 445 item 5

// sara o
// TODO 9905. โชวะ in data/best/article.zip(article/article_00084.txt) line 64 item 28
// TODO 11141. เชโกสโลวะเกีย in data/best/article.zip(article/article_00111.txt) line 38 item 94
// TODO โต๊ระ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00061.txt) line 445 item 5
// TODO 27401. แล็กโทบะซิลลัส in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00107.txt) line 572 item 18
// TODO 35100. สโลวะเกีย in data/best/news.zip(news/news_00056.txt) line 102 item 74
// TODO 42315. ทาวีโกวะ in data/best/novel.zip(novel/novel_00014.txt) line 61 item 24
// TODO 44014. โผล๊ะ in data/best/novel.zip(novel/novel_00021.txt) line 255 item 62

// best ae..a; 40612, 43914? 44003

// 1109. เกาะติด in data/best/article.zip(article/article_00002.txt) line 100 item 81
// 1351. เพาะปลูก in data/best/article.zip(article/article_00003.txt) line 30 item 222
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
*/
