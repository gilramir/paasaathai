package paasaathai

// Based on ideas from, but not the implemention defined in,
// "Character Cluster Based Thai Information Retrieval"
// By Thanaruk Theeramunkong, Virach Sornlertlamvanich,
// Thanasan Tanhermhong, and Wirat Chinnan
//
// https://www.researchgate.net/publication/2853284_Character_Cluster_Based_Thai_Information_Retrieval
// and more specifically,
// https://www.researchgate.net/profile/Virach-Sornlertlamvanich/publication/2853284_Character_Cluster_Based_Thai_Information_Retrieval/links/02e7e514db194bcb1f000000/Character-Cluster-Based-Thai-Information-Retrieval.pdf
//
// Also based on rules explained in "Thai for Beginners" by Benjawan Poomsan
// Becker, ISBN 1-887521-00 3

import (
	"fmt"
	"strings"

	"github.com/gilramir/objregexp"
)

type GStackCluster struct {
	// The UTF-8 string in this cluster
	Text string

	// Is it comprised completely of Thai code points?
	IsThai bool

	// Did it contain a valid sequence of Thai code points?
	IsValidThai bool

	// If invalid, this is the reason
	InvalidReason string

	// These are the well-known parts of a cluster
	// Not always set, but if so, this is the vowel that
	// goes in front.
	FrontVowel     GraphemeStack
	FirstConsonant GraphemeStack

	// If the cluster is for punctuation instead of
	// consonants and vowels, this is set
	SingleMidSign GraphemeStack

	// The Tail might contain consontants and vowels,
	// only consonants, or only vowels.
	Tail []GraphemeStack

	// The name of the rule that created this cluster.
	MatchingRule string
}

func (s *GStackCluster) Repr() string {
	if s.IsThai {
		if s.IsValidThai {
			if s.SingleMidSign.Main != 0 {
				return fmt.Sprintf("<CC Thai %s SS:%s", s.Text, s.SingleMidSign.Repr())
			} else {
				result := fmt.Sprintf("<CC Thai %s", s.Text)
				if s.FrontVowel.Main != 0 {
					result += fmt.Sprintf(" FV:%s", s.FrontVowel.Repr())
				}
				if s.FirstConsonant.Main != 0 {
					result += fmt.Sprintf(" FC:%s", s.FirstConsonant.Repr())
				}
				if len(s.Tail) > 0 {
					for ti, tgs := range s.Tail {
						result += fmt.Sprintf(" T%d:%s", ti, tgs.Repr())
					}
				}
				return result + ">"
			}
		} else {
			return fmt.Sprintf("<CC Invalid-Thai: %s Reason: %s>", s.Text, s.InvalidReason)
		}
	} else {
		return fmt.Sprintf("<CC Not-Thai: %s>", s.Text)
	}
}

func makeCluster(input []GraphemeStack) GStackCluster {
	isThai := true
	isValidThai := true
	text := ""
	for _, gs := range input {
		text += gs.Text
		if gs.IsThai() {
			if !gs.IsValidThai() {
				isValidThai = false
			}
		} else {
			isThai = false
		}
	}
	tcc := GStackCluster{
		Text:        text,
		IsThai:      isThai,
		IsValidThai: isValidThai,
	}
	// More TODO
	return tcc
}

type GStackClusterParser struct {
	compiler objregexp.Compiler[GraphemeStack]
}

/* The vowel orthograpy patterns from "Thai For Beginners" p. 243
	short				long
	pattern	regex			pattern regex
#1	-ะ	final_pos_short_1	-า	final_pos_long_1
#2	-ิอิ				-อ๊
#3					uu_ua
#4
#5		final_pos_short_1		final_pos_long_2
#6		final_pos_short_1		final_pos_long_2
#7		final_pos_short_1		final_pos_long_2
#8		eu_o_ao				final_pos_long_1
#9		ua				ua
#10		ia				ia
#11		uu_ua				uu_ua
#12		eu_o_ao				eu_o_ao

And from p. 244

#1		c_sara_am
#2		final_pos_long_2
#3		final_pos_long_2
#4		eu_o_ao
#5		eei

And from p. 250

	final				medial
	pattern	regex			pattern regex
#1
#2
#3
#4
#5
#6
#7

*/

type TccRule struct {
	name  string
	rs    string
	ck    func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool
	regex *objregexp.Regexp[GraphemeStack]
}

func (s *TccRule) CompileWith(compiler *objregexp.Compiler[GraphemeStack]) {
	s.regex = compiler.MustCompile(s.rs)
}

/*
var ConsonantsAllowedAtFront = NewSetFromSlice([]rune{
	'ก', 'ข', 'ค', 'ง', 'จ',
	'ฉ', 'ช', 'ซ', 'ญ', 'ด',
	'ต', 'ถ', 'ท', 'ธ', 'น',
	'บ', 'ป', 'ผ', 'ฝ', 'พ',
	'ฟ', 'ภ', 'ม', 'ย', 'ร',
	'ล', 'ว', 'ศ', 'ส', 'ห',
	'อ',
})
*/

// The initial O_ANG changes a subsequent
// consonant into being a MID class consonant.
var ConsonantsAllowedAfterOAng = NewSetFromSlice([]rune{
	'ย',
})

// The initial HO_HIP changes a subsequent
// consonant into being a HIGH class consonant.
var LowConsonantsAllowedAfterHoHip = NewSetFromSlice([]rune{
	'ญ', 'ง', 'น', 'ม',
	'ย', 'ร', 'ล', 'ว',
})

// TODO check yo ying 40156. เผียะ in data/best/novel.zip(novel/novel_00001.txt) line 331 item 4

// Consonants that can glide with lo ling
// แต่ละ in data/best/article.zip(article/article_00001.txt) line 14 item 40
// แคระแกร็น in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00003.txt) line 344 item 6
// ตบแผละ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00023.txt) line 25 item 11
// แกละ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00023.txt) line 326 item 2
// TODO check not gliding 38296. หมู่บ้านการเคหะร่มเกล้า in data/best/news.zip(news/news_00080.txt) line 20 item 121
// TODO check not gliding 47608. เพคะ in data/best/novel.zip(novel/novel_00071.txt) line 52 item 3

var ConsonantsAllowedBeforeGlidingLoLing = NewSetFromSlice([]rune{
	THAI_CHARACTER_TO_TAO,
	THAI_CHARACTER_KHO_KHWAI,
	THAI_CHARACTER_PHO_PHUNG,
	THAI_CHARACTER_KO_KAI,
})

// Consonants that can glide with ro rua
// แกระ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00026.txt) line 177 item 18
// ว่านหางนกยูงแคระ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00030.txt) line 133 item 38
// TODO 33902. แม่ระมาด in data/best/news.zip(news/news_00045.txt) line 91 item 33
// TODO 40612. แสยะ in data/best/novel.zip(novel/novel_00004.txt) line 11 item 22
// TODO 11829. รัฐอิสลามอเสระ in data/best/article.zip(article/article_00123.txt) line 64 item 58
// แประ in data/best/novel.zip(novel/novel_00021.txt) line 238 item 37
// พระเถรานุเถระ in data/best/article.zip(article/article_00013.txt) line 27 item 25
var ConsonantsAllowedBeforeGlidingRoRua = NewSetFromSlice([]rune{
	THAI_CHARACTER_KO_KAI,
	THAI_CHARACTER_KHO_KHWAI,
	THAI_CHARACTER_PO_PLA,
	THAI_CHARACTER_THO_THUNG,
})

// Consonants that can glide with wo waen
// แขวะ in data/best/news.zip(news/news_00008.txt) line 93 item 2
var ConsonantsAllowedBeforeGlidingWoWaen = NewSetFromSlice([]rune{
	THAI_CHARACTER_KHO_KHAI,
})

// TODO chepa!
// 1754. เฉพาะเจาะจง in data/best/article.zip(article/article_00004.txt) line 13 item 51

//2128. เยอะ in data/best/article.zip(article/article_00005.txt) line 38 item 139
// 3806. พระเถรานุเถระ in data/best/article.zip(article/article_00013.txt) line 27 item 25

//19191. ปาละเสนะ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00027.txt) line 254 item 17

// Consonants allowed after other consonants to make a single
// sound
// But...
// "เฉพาะ"
// "โบราณ"
// "เพราะ"
// "กว่า"
/*
var GlidingConsonants = NewSetFromSlice([]rune{
	'ล',
	THAI_CHARACTER_PHO_PHAN,
	THAI_CHARACTER_RO_RUA,
	THAI_CHARACTER_WO_WAEN,
})
*/

func curriedHasClass(r rune) func(GraphemeStack) bool {
	return func(g GraphemeStack) bool {
		return g.Main == r || g.DiacriticVowel == r || g.UpperDiacritic == r
	}
}

func (s *GStackClusterParser) Initialize() {
	s.compiler.Initialize()

	// regex classes
	/*
		s.compiler.MakeClass("gaaw",
			func(gs GraphemeStack) bool {
				return gs.String() == "ก็"
			})
	*/
	s.compiler.MakeClass("consonant",
		func(gs GraphemeStack) bool {
			return RuneIsConsonant(gs.Main)
		})
	s.compiler.MakeClass("diacritic vowel",
		func(gs GraphemeStack) bool {
			return gs.DiacriticVowel != 0
		})

	s.compiler.MakeClass("low consonant after ho hip",
		func(gs GraphemeStack) bool {
			return LowConsonantsAllowedAfterHoHip.Has(gs.Main)
		})

	s.compiler.MakeClass("consonant before gliding lo ling",
		func(gs GraphemeStack) bool {
			return ConsonantsAllowedBeforeGlidingLoLing.Has(gs.Main)
		})

	s.compiler.MakeClass("consonant before gliding ro rua",
		func(gs GraphemeStack) bool {
			return ConsonantsAllowedBeforeGlidingRoRua.Has(gs.Main)
		})

	s.compiler.MakeClass("consonant before gliding wo waen",
		func(gs GraphemeStack) bool {
			return ConsonantsAllowedBeforeGlidingWoWaen.Has(gs.Main)
		})

	/*
		s.compiler.MakeClass("has-sara-i",
			func(gs GraphemeStack) bool {
				return gs.DiacriticVowel == THAI_CHARACTER_SARA_I
			})

		s.compiler.MakeClass("has-sara-ii",
			func(gs GraphemeStack) bool {
				return gs.DiacriticVowel == THAI_CHARACTER_SARA_II
			})

		s.compiler.MakeClass("has-sara-uee",
			func(gs GraphemeStack) bool {
				return gs.DiacriticVowel == THAI_CHARACTER_SARA_UEE
			})

		s.compiler.MakeClass("has-mai-han-akat",
			func(gs GraphemeStack) bool {
				return gs.DiacriticVowel == THAI_CHARACTER_MAI_HAN_AKAT
			})

		s.compiler.MakeClass("has-maithaku",
			func(gs GraphemeStack) bool {
				return RuneIsConsonant(gs.Main) && gs.DiacriticVowel == THAI_CHARACTER_MAITAIKHU
			})
	*/
	/*
		s.compiler.MakeClass("sliding-consonant",
			func(gs GraphemeStack) bool {
				return GlidingConsonants.Has(gs.Main)
			})
	*/
	s.compiler.MakeClass("front position vowel",
		func(gs GraphemeStack) bool {
			return RuneIsFrontPositionVowel(gs.Main)
		})

	s.compiler.MakeClass("mid position vowel",
		func(gs GraphemeStack) bool {
			return RuneIsMidPositionVowel(gs.Main)
		})

	// regex identity classes for:
	// digits, non-diacritic vowels, currency, and other mid-position signs
	for fullName, thaiRune := range ThaiNameToRune {
		var name string
		if fullName[0:11] == "THAI_DIGIT_" {
			name = fullName[11:]
		} else if fullName[0:15] == "THAI_CHARACTER_" {
			if RuneIsConsonant(thaiRune) || RuneIsDiacritic(thaiRune) {
				continue
			}
			name = fullName[15:]
		} else if fullName[0:21] == "THAI_CURRENCY_SYMBOL_" {
			name = fullName[21:]
		} else {
			panic(fmt.Sprintf("Didn't expect code point name %s", fullName))
		}

		name = strings.ToLower(name)
		name = strings.ReplaceAll(name, "_", " ")
		s.compiler.AddIdentity(name,
			MustParseSingleGraphemeStack(string(thaiRune)))
	}

	// has-* tests for all consonants and diacritics
	// (anything that can be stacked)
	for fullName, thaiRune := range ThaiNameToRune {
		if !RuneIsConsonant(thaiRune) && !RuneIsDiacritic(thaiRune) {
			continue
		}

		var name string
		if fullName[0:15] == "THAI_CHARACTER_" {
			name = fullName[15:]
		} else {
			panic(fmt.Sprintf("Didn't expect code point name %s", fullName))
		}

		name = strings.ToLower(name)
		name = strings.ReplaceAll(name, "_", " ")
		s.compiler.MakeClass(name, curriedHasClass(thaiRune))
	}

	s.compiler.Finalize()

	r_maybe_sandwich_sara_a.CompileWith(&s.compiler)
	r_sara_a_aa.CompileWith(&s.compiler)
	r_sara_uee.CompileWith(&s.compiler)
	/*
		r_final_pos_short_1.CompileWith(&s.compiler)
		r_final_pos_short_3.CompileWith(&s.compiler)
		r_front_o.CompileWith(&s.compiler)
		r_final_pos_long_1.CompileWith(&s.compiler)
		r_final_pos_long_2_o_ang.CompileWith(&s.compiler)
		r_final_pos_long_2.CompileWith(&s.compiler)
		r_final_pos_eei.CompileWith(&s.compiler)
		r_eu_o_ao.CompileWith(&s.compiler)
		r_medial_maithaiku.CompileWith(&s.compiler)
		r_medial_sara_i.CompileWith(&s.compiler)
		r_ia.CompileWith(&s.compiler)
		r_uu_ua.CompileWith(&s.compiler)
		r_c_sara_am.CompileWith(&s.compiler)
		r_mai_han_akat.CompileWith(&s.compiler)
		r_ua.CompileWith(&s.compiler)
		r_medial_ua.CompileWith(&s.compiler)
		r_rr.CompileWith(&s.compiler)
		r_standalone_symbol.CompileWith(&s.compiler)
		r_gaaw.CompileWith(&s.compiler)
	*/
	r_single_diacritic_vowel.CompileWith(&s.compiler)
	r_single_consonant.CompileWith(&s.compiler)
}

func assertGroupLength(reg objregexp.Range, length int) {
	if reg.Length() != length {
		panic(fmt.Sprintf("Group expected to have length %d. Got: %+v",
			length, reg))
	}
}

func (s *GStackClusterParser) CompileRegex(text string) (*objregexp.Regexp[GraphemeStack], error) {
	return s.compiler.Compile(text)
}

func (s *GStackClusterParser) ParseGraphemeStacks(input []GraphemeStack) []GStackCluster {

	// Estimate the space by dividing the number of UTF-8 bytes
	// by 3 (each Thai code point needs 3), and then taking 2/3 of that
	estimatedAllocation := len(input) * 2 / 3
	clusters := make([]GStackCluster, 0, estimatedAllocation)

	rules := []TccRule{
		r_maybe_sandwich_sara_a, // this comes after sandwich a_aa
		r_sara_a_aa,
		r_sara_uee,
		r_single_diacritic_vowel, // this comes after other vowels
		r_single_consonant,       // this needs to be the last rule
	}
	/*
		rules := []TccRule{
			r_final_pos_short_1,
			r_medial_maithaiku,
			r_medial_sara_i,
			r_final_pos_short_3,
			r_final_pos_long_1,
			r_eu_o_ao,                // must come before long_2
			r_final_pos_eei,          // eei must come before long_2
			r_front_o,                // must come before long_2
			r_final_pos_long_2_o_ang, // must come before long_2
			r_final_pos_long_2,
			r_ia,
			r_ua,
			r_medial_ua,
			r_rr,
			r_uu_ua,
			r_c_sara_am,
			r_mai_han_akat,
			r_standalone_symbol,
			r_gaaw,
			r_single_consonant,
		}
	*/

next_input:
	for i := 0; i < len(input); {

		// Is this not Thai?
		if !input[i].IsThai() {
			c := makeCluster(input[i : i+1])
			c.FirstConsonant = input[i]
			clusters = append(clusters, c)
			i++
			continue
		}

		var c GStackCluster
		var length int

		// Run the regexes, in order
		for _, rule := range rules {
			matched := rule.ck(&rule, input, i, &length, &c)
			if matched {
				c.MatchingRule = rule.name
				fmt.Printf("matched: %s %s %s\n", rule.name, rule.rs, c.Repr())
				clusters = append(clusters, c)
				i += length
				continue next_input
			}
		}

		// No success.
		msg := fmt.Sprintf("Can't handle gstack at pos %d: %s", i, input[i].Repr())
		panic(msg)
	}

	return clusters
}

// Short vowel sara a or aa pattern
var r_sara_a_aa = TccRule{
	name: "sara_a_aa",
	rs: "([:consonant: && !:diacritic vowel:])" +
		"([:sara a:] | [:sara aa:])",
	/*
		rs: "([:sara_e:] | [:sara_ae:] | [:sara_o:])? " +
			"([:consonant: && !:diacritic-vowel:]) ([:sliding-consonant:])? " +
			"([:sara_a:])",
	*/
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		/*
			reg1 := m.Group(1)
			if !reg1.Empty() {
				assertGroupLength(reg1, 1)
				c.FrontVowel = input[reg1.Start]
			}
		*/

		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]

		reg2 := m.Group(2)
		c.Tail = append(c.Tail, input[reg2.Start])

		*length = m.Length()
		return true
	},
}

// Single diacritic vowel
var r_single_diacritic_vowel = TccRule{
	name: "single_diacritic_vowel",
	rs: "([:consonant: && " +
		"(:sara i: || :sara ii: || :sara ue: || :sara u: || :sara uu:) ])",
	/*
		rs: "([:sara_e:] | [:sara_ae:] | [:sara_o:])? " +
			"([:consonant: && !:diacritic-vowel:]) ([:sliding-consonant:])? " +
			"([:sara_a:])",
	*/
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		/*
			reg1 := m.Group(1)
			if !reg1.Empty() {
				assertGroupLength(reg1, 1)
				c.FrontVowel = input[reg1.Start]
			}
		*/

		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]

		*length = m.Length()
		return true
	},
}

var r_sara_uee = TccRule{
	name: "sara_uee",
	rs:   "([:consonant: && :sara uee:]) ([:o ang:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		/*
			reg1 := m.Group(1)
			if !reg1.Empty() {
				assertGroupLength(reg1, 1)
				c.FrontVowel = input[reg1.Start]
			}
		*/

		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]

		reg2 := m.Group(2)
		c.Tail = append(c.Tail, input[reg2.Start])

		*length = m.Length()
		return true
	},
}

// TODO - instead of no diacritic, allow this
// TODO โต๊ระ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00061.txt) line 445 item 5
var r_maybe_sandwich_sara_a = TccRule{
	name: "maybe_sandwich_sara_a",
	rs: "([:sara e:] | [:sara ae:] | [:sara o:]) " +
		"(" +
		// BEGIN possible consonants allowed between sandwich vowels
		"([:consonant: && !:diacritic vowel:]) | " +
		"([:ho hip: && !:diacritic vowel:] [:low consonant after ho hip: && !:diacritic vowel:]) | " +
		"([:consonant before gliding lo ling: && !:diacritic vowel:] [:lo ling: && !:diacritic vowel:]) " +
		"([:consonant before gliding ro rua: && !:diacritic vowel:] [:ro rua: && !:diacritic vowel:]) " +
		"([:consonant before gliding wo waen: && !:diacritic vowel:] [:wo waen: && !:diacritic vowel:]) " +
		// END   possible consonants allowed between sandwich vowels
		")" +
		"([:sara a:]?)",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		fmt.Printf("reg1: %v\n", reg1)
		c.FrontVowel = input[reg1.Start]

		reg2 := m.Group(2)
		fmt.Printf("reg2: %v\n", reg1)
		c.FirstConsonant = input[reg2.Start]
		if reg2.Length() > 1 {
			c.Tail = append(c.Tail, input[reg2.Start+1:reg2.End]...)
		}

		if m.HasGroup(8) {
			reg3 := m.Group(8)
			c.Tail = append(c.Tail, input[reg3.Start])
		}

		*length = m.Length()
		return true
	},
}

//24319. แอนะล็อก in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00076.txt) line 133 item 10
//	THAI_CHARACTER_SARA_AE, THAI_CHARACTER_O_ANG, THAI_CHARACTER_NO_NU, THAI_CHARACTER_SARA_A, THAI_CHARACTER_LO_LING, THAI_CHARACTER_MAITAIKHU, THAI_CHARACTER_O_ANG, THAI_CHARACTER_KO_KAI

/*
// Short vowel sara a patterns
var r_final_pos_short_1 = TccRule{
	name: "final_pos_short_1",
	rs: "([:sara_e:] | [:sara_ae:] | [:sara_o:])? " +
		"([:consonant: && !:diacritic-vowel:]) ([:sliding-consonant:])? " +
		"([:sara_a:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		if !reg1.Empty() {
			assertGroupLength(reg1, 1)
			c.FrontVowel = input[reg1.Start]
		}

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		reg3 := m.Group(3)
		if !reg3.Empty() {
			c.Tail = append(c.Tail, input[reg3.Start:reg3.End]...)
		}
		reg4 := m.Group(4)
		c.Tail = append(c.Tail, input[reg4.Start:reg4.End]...)

		*length = m.Length()
		return true
	},
}

// Short vowel a/oh patterns
var r_final_pos_short_3 = TccRule{
	name: "final_pos_short_3",
	rs: "([:sara_e:]) " +
		"([:consonant: && !:diacritic-vowel:]) ([:sliding-consonant:])? " +
		"([:sara_aa:]|[:o_ang:]) ([:sara_a:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		reg3 := m.Group(3)
		if !reg3.Empty() {
			c.Tail = append(c.Tail, input[reg3.Start:reg3.End]...)
		}
		reg4 := m.Group(4)
		c.Tail = append(c.Tail, input[reg4.Start:reg4.End]...)

		reg5 := m.Group(5)
		c.Tail = append(c.Tail, input[reg5.Start])

		*length = m.Length()
		return true
	},
}

// Long vowel patterns; these must be checked after the short vowels
var r_final_pos_long_1 = TccRule{
	name: "final_pos_long_1",
	rs: "([:consonant: && !:diacritic-vowel:]) ([:sliding-consonant:])? " +
		"([:sara_aa:] | [:o_ang:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]

		if m.HasGroup(2) {
			reg2 := m.Group(2)
			c.Tail = append(c.Tail, input[reg2.Start:reg2.End]...)
		}
		reg3 := m.Group(3)
		c.Tail = append(c.Tail, input[reg3.Start:reg3.End]...)

		*length = m.Length()
		return true
	},
}

// "โบราณ"
var r_front_o = TccRule{
	//	rs: "([:sara_o:]) " +
	//		"([:consonant: && !:diacritic-vowel:])",
	name: "front_o",
	rs: "([:front-position-vowel:]) " +
		"([:consonant: && !:o_ang: && !:diacritic-vowel:]) " +
		"([:sliding-consonant: && :diacritic-vowel:] | " +
		"([:sliding-consonant:][:mid-position-vowel:]))",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}

		reg1 := m.Group(1)
		reg2 := m.Group(2)

		// We don't use the full length
		partialLen := reg1.Length() + reg2.Length()
		*c = makeCluster(input[i : i+partialLen])

		c.FirstConsonant = input[reg2.Start]

		*length = partialLen
		return true
	},
}

var r_final_pos_long_2_o_ang = TccRule{
	// ":o_ang:" does already mean "!:diactricit-vowel:"
	name: "final_pos_long_2_o_ang",
	rs: "([:front-position-vowel:]) " +
		"([:o_ang:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])

		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		*length = m.Length()
		return true
	},
}

// Similar to re_final_pos_short_1, but no sara_a at the end
// We allow all front-position vowels here
var r_final_pos_long_2 = TccRule{
	name: "pos_long_2",
	rs: "([:front-position-vowel:]) " +
		"([:consonant: && !:o_ang: && !:diacritic-vowel:]) ([:sliding-consonant:])?",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])

		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		if m.HasGroup(3) {
			reg3 := m.Group(3)
			c.Tail = append(c.Tail, input[reg3.Start])
		}

		*length = m.Length()
		return true
	},
}

// sara e + C + (eu,o,ao) vowel sandwich
var r_eu_o_ao = TccRule{
	name: "eu_o_ao",
	rs: "([:sara_e:]) " +
		"([:consonant: && !:diacritic-vowel:]) ([:sliding-consonant:])?" +
		"([:sara_aa:] | [:o_ang:] [:sara_a:]?)",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])

		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		if m.HasGroup(3) {
			reg3 := m.Group(3)
			c.Tail = append(c.Tail, input[reg3.Start])
		}
		reg4 := m.Group(4)
		c.Tail = append(c.Tail, input[reg4.Start])

		*length = m.Length()
		return true
	},
}

// sara e + C + e vowel sandwich
var r_final_pos_eei = TccRule{
	name: "final_pos_eei",
	rs: "([:sara_e:]) " +
		"([:consonant: && !:diacritic-vowel:]) ([:sliding-consonant:])?" +
		"([:yo_yak:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])

		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		if m.HasGroup(3) {
			reg3 := m.Group(3)
			c.Tail = append(c.Tail, input[reg3.Start])
		}
		reg4 := m.Group(4)
		c.Tail = append(c.Tail, input[reg4.Start])

		*length = m.Length()
		return true
	},
}

// sara e ? = C + uu/ua
var r_uu_ua = TccRule{
	name: "uu_ua",
	rs: "([:sara_e:]?) " +
		"([:consonant: && :has-sara-uee:]) ([:sliding-consonant:]?)" +
		"([:o_ang:] [:sara_a:]?)",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])

		if m.HasGroup(1) {
			reg1 := m.Group(1)
			c.FrontVowel = input[reg1.Start]
		}

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		if m.HasGroup(3) {
			reg3 := m.Group(3)
			c.Tail = append(c.Tail, input[reg3.Start])
		}
		reg4 := m.Group(4)
		c.Tail = append(c.Tail, input[reg4.Start])

		*length = m.Length()
		return true
	},
}

// Medial position maithaku patterns
var r_medial_maithaiku = TccRule{
	name: "medial_maithaiku",
	rs: "([:sara_e:] | [:sara_ae:] ) " +
		"([:consonant: && :has-maithaku:]) ([:sliding-consonant:])? " +
		"([:consonant: && !:diacritic-vowel:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		reg3 := m.Group(3)
		if !reg3.Empty() {
			c.Tail = append(c.Tail, input[reg3.Start:reg3.End]...)
		}
		reg4 := m.Group(4)
		c.Tail = append(c.Tail, input[reg4.Start:reg4.End]...)

		*length = m.Length()
		return true
	},
}

// Medial position sara i patterns
var r_medial_sara_i = TccRule{
	name: "medial_sara_i",
	rs: "([:sara_e:] ) " +
		"([:consonant: && :has-sara-i:]) ([:sliding-consonant:])? " +
		"([:consonant: && !:diacritic-vowel:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		reg3 := m.Group(3)
		if !reg3.Empty() {
			c.Tail = append(c.Tail, input[reg3.Start:reg3.End]...)
		}
		reg4 := m.Group(4)
		c.Tail = append(c.Tail, input[reg4.Start:reg4.End]...)

		*length = m.Length()
		return true
	},
}

// long or short ia
var r_ia = TccRule{
	name: "r_ia",
	rs: "([:sara_e:] ) " +
		"([:consonant: && :has-sara-ii:]) ([:sliding-consonant:])? " +
		"([:yo_yak:] [:sara_a:]?)",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		reg2 := m.Group(2)
		c.FirstConsonant = input[reg2.Start]

		reg3 := m.Group(3)
		if !reg3.Empty() {
			c.Tail = append(c.Tail, input[reg3.Start:reg3.End]...)
		}
		reg4 := m.Group(4)
		c.Tail = append(c.Tail, input[reg4.Start:reg4.End]...)

		*length = m.Length()
		return true
	},
}

// long or short ua
var r_ua = TccRule{
	name: "r_ua",
	rs: "([:consonant: && :has-mai-han-akat:]) ([:sliding-consonant:]? " +
		"[:wo_waen:] [:sara_a:]?)",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]

		reg2 := m.Group(2)
		c.Tail = append(c.Tail, input[reg2.Start:reg2.End]...)

		*length = m.Length()
		return true
	},
}

// medial_ua
var r_medial_ua = TccRule{
	name: "medial_ua",
	rs: "([:consonant: && !:diacritic-vowel:]) ([:sliding-consonant:]? " +
		"[:wo_waen:] [:consonant:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]

		reg2 := m.Group(2)
		c.Tail = append(c.Tail, input[reg2.Start:reg2.End]...)

		*length = m.Length()
		return true
	},
}

// rr
// "บทบรรณาธิการ"
var r_rr = TccRule{
	name: "rr",
	rs:   "([:consonant:]) ([:ro_rua:][:ro_rua:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]

		reg2 := m.Group(2)
		c.Tail = append(c.Tail, input[reg2.Start:reg2.End]...)

		*length = m.Length()
		return true
	},
}

// C + SARA_AM... can this be combined with re_final_pos_long_1 ?
var r_c_sara_am = TccRule{
	name: "c_sara_am",
	rs:   "([:consonant: && !:diacritic-vowel:]) ([:sara_am:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+2])
		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]

		reg2 := m.Group(2)
		c.Tail = append(c.Tail, input[reg2.Start])

		*length = m.Length()
		return true
	},
}

var r_mai_han_akat = TccRule{
	name: "mai_han_akat",
	rs:   "([:consonant: && :has-mai-han-akat:]) ([:consonant:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])

		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]

		reg2 := m.Group(2)
		c.Tail = append(c.Tail, input[reg2.Start])

		*length = m.Length()
		return true
	},
}

// Standalone symbol
var r_standalone_symbol = TccRule{
	name: "standalone_symbol",
	rs:   "[:paiyannoi:]",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+1])
		c.FirstConsonant = input[i]
		*length = m.Length()
		return true
	},
}

// From BNF in "Character Cluster Based Thai Information Retrieval"
// But do we really need this?
var r_gaaw = TccRule{
	name: "gaaw",
	rs:   "[:gaaw:]",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+1])
		c.FirstConsonant = input[i]
		*length = m.Length()
		return true
	},
}
*/

// Do we at least have a bare consonant?
var r_single_consonant = TccRule{
	name: "consonant",
	rs:   "[:consonant: && !:diacritic vowel:]",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+1])
		c.FirstConsonant = input[i]
		*length = m.Length()
		return true
	},
}
