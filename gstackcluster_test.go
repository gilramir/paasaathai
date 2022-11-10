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

// long tests

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
