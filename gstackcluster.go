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

const (
	ReasonSaraEInvalidCombo  = 1
	ReasonSaraAeInvalidCombo = 2
	ReasonSoloDiacritic      = 3
)

type GStackCluster struct {
	// The UTF-8 string in this cluster
	Text string

	// Is it comprised completely of Thai code points?
	IsThai bool

	// Did it contain a valid sequence of Thai code points?
	IsValidThai bool

	// If invalid, this is the reason
	InvalidReason int

	// These are the well-known parts of a cluster
	// Not always set, but if so, this is the vowel that
	// goes in front.
	// TODO - FrontVowel can be a rune
	FrontVowel     GraphemeStack
	FirstConsonant GraphemeStack

	// If the cluster is for punctuation instead of
	// consonants and vowels, this is set
	// TODO - SingleMidsign can be a rune
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
			return fmt.Sprintf("<CC Invalid-Thai: %s Reason: %d>", s.Text, s.InvalidReason)
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
/*
var ConsonantsAllowedAfterOAng = NewSetFromSlice([]rune{
	'ย',
})
*/

// The initial HO_HIP changes a subsequent
// consonant into being a HIGH class consonant.
var LowConsonantsAllowedAfterHoHip = NewSetFromSlice([]rune{
	'ญ', 'ง', 'น', 'ม',
	'ย', 'ร', 'ล', 'ว',
})

// Consonants that can glide with lo ling
var ConsonantsAllowedBeforeGlidingLoLing = NewSetFromSlice([]rune{
	THAI_CHARACTER_KO_KAI,
	THAI_CHARACTER_KHO_KHWAI,
	THAI_CHARACTER_PHO_PHUNG,
	THAI_CHARACTER_PHO_PHAN,
	THAI_CHARACTER_PO_PLA,
})

// 2 TODO check yo ying 40156. เผียะ in data/best/novel.zip(novel/novel_00001.txt) line 331 item 4
// 2 TODO check not gliding 38296. หมู่บ้านการเคหะร่มเกล้า in data/best/news.zip(news/news_00080.txt) line 20 item 121
// 2 TODO check not gliding 47608. เพคะ in data/best/novel.zip(novel/novel_00071.txt) line 52 item 3

// Consonants that can glide with ro rua
// 1 แกระ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00026.txt) line 177 item 18
// 1 แคระแกร็น in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00003.txt) line 344 item 6
// 1 ว่านหางนกยูงแคระ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00030.txt) line 133 item 38
// 2 แประ in data/best/novel.zip(novel/novel_00021.txt) line 238 item 37
// 2 พระเถรานุเถระ in data/best/article.zip(article/article_00013.txt) line 27 item 25
// 2 TODO 33902. แม่ระมาด in data/best/news.zip(news/news_00045.txt) line 91 item 33
// 2 TODO 40612. แสยะ in data/best/novel.zip(novel/novel_00004.txt) line 11 item 22
// 2 TODO 11829. รัฐอิสลามอเสระ in data/best/article.zip(article/article_00123.txt) line 64 item 58
var ConsonantsAllowedBeforeGlidingRoRua = NewSetFromSlice([]rune{
	THAI_CHARACTER_KO_KAI,
	THAI_CHARACTER_KHO_KHWAI,
	THAI_CHARACTER_PO_PLA,
	THAI_CHARACTER_PHO_PHAN,
	THAI_CHARACTER_THO_THUNG,
})

// Consonants that can glide with wo waen
// 1 แขวะ in data/best/news.zip(news/news_00008.txt) line 93 item 2
var ConsonantsAllowedBeforeGlidingWoWaen = NewSetFromSlice([]rune{
	THAI_CHARACTER_KHO_KHAI,
})

//var ConsonantsAllowedAfterGlidingSoSuea = NewSetFromSlice([]rune{

// TODO chepa!
// 1754. เฉพาะเจาะจง in data/best/article.zip(article/article_00004.txt) line 13 item 51

//2128. เยอะ in data/best/article.zip(article/article_00005.txt) line 38 item 139
// 3806. พระเถรานุเถระ in data/best/article.zip(article/article_00013.txt) line 27 item 25

// 2 19191. ปาละเสนะ in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00027.txt) line 254 item 17

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

	s.compiler.MakeClass("tone mark",
		func(gs GraphemeStack) bool {
			return RuneIsToneMark(gs.UpperDiacritic)
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

	s.compiler.MakeClass("front position vowel",
		func(gs GraphemeStack) bool {
			return RuneIsFrontPositionVowel(gs.Main)
		})

	s.compiler.MakeClass("mid position vowel",
		func(gs GraphemeStack) bool {
			return RuneIsMidPositionVowel(gs.Main)
		})

	s.compiler.MakeClass("mid position sign",
		func(gs GraphemeStack) bool {
			return RuneIsMidPositionSign(gs.Main)
		})

	s.compiler.MakeClass("upper position sign",
		func(gs GraphemeStack) bool {
			return RuneIsUpperPositionSign(gs.UpperDiacritic)
		})

	s.compiler.MakeClass("digit",
		func(gs GraphemeStack) bool {
			return RuneIsDigit(gs.Main)
		})

	// regex identity classes for:
	// digits, non-diacritic vowels, currency, and other mid-position signs
	// for consonant, prefix with "bare "
	for fullName, thaiRune := range ThaiNameToRune {
		var name string
		if fullName[0:11] == "THAI_DIGIT_" {
			name = fullName[11:]
		} else if fullName[0:15] == "THAI_CHARACTER_" {
			name = fullName[15:]
			if RuneIsConsonant(thaiRune) || RuneIsDiacritic(thaiRune) {
				name = "bare " + name
			}
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

	// regex iidentity class for all consonants and diacritics
	// (anything that can be stacked)
	// This checks if the stack *contains* the consonant or diacritic,
	// allowing other runes in the stack to exist.
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
	r_sandwich_ia.CompileWith(&s.compiler)
	r_sandwich_ueea_er.CompileWith(&s.compiler)
	r_special_o_ang.CompileWith(&s.compiler)
	r_short_o_ang.CompileWith(&s.compiler)
	r_sara_a_aa.CompileWith(&s.compiler)
	r_sara_uee.CompileWith(&s.compiler)
	r_sara_ai.CompileWith(&s.compiler)
	r_medial_er.CompileWith(&s.compiler)
	r_ua.CompileWith(&s.compiler)
	r_sara_am.CompileWith(&s.compiler)
	r_mai_han_akat.CompileWith(&s.compiler)
	r_sandwich_ao.CompileWith(&s.compiler)
	r_single_diacritic_vowel.CompileWith(&s.compiler)
	r_single_consonant.CompileWith(&s.compiler)
	r_punctuation_or_digit.CompileWith(&s.compiler)
	r_sanskrit.CompileWith(&s.compiler)
	r_error_solo_diacritic.CompileWith(&s.compiler)
	r_error_sara_e_ae.CompileWith(&s.compiler)
	r_error_short_o_ang.CompileWith(&s.compiler)
	r_error_solo_final_vowel.CompileWith(&s.compiler)
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
		r_special_o_ang,    // must come before short_o_ang
		r_short_o_ang,      // must come before maybe_sandwich_sara_a
		r_sandwich_ia,      // must come before maybe_sandwich_sara_a
		r_sandwich_ueea_er, // must come before maybe_sandwich_sara_a
		r_medial_er,        // must come before maybe_sandwich_sara_a
		r_sandwich_ao,      // must come before maybe_sandwich_sara_a
		r_sara_ai,
		r_maybe_sandwich_sara_a,
		r_sara_a_aa,
		r_sara_uee,
		r_ua,
		r_sara_am,
		r_mai_han_akat,
		r_single_diacritic_vowel, // this comes after other vowels
		r_sanskrit,               // this must come before single_consonant
		r_single_consonant,       // this needs to be the last consonant rule
		r_punctuation_or_digit,
		r_error_solo_diacritic,
		r_error_solo_final_vowel,
		r_error_sara_e_ae,   // this must come after maybe_sandwich_sara_a
		r_error_short_o_ang, // this must come after maybe_sandwich_sara_a
	}

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
				fmt.Printf("matched: %s @i=%d length=%d %s\n",
					rule.name, i, length, c.Repr())
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

var r_sara_am = TccRule{
	name: "sara_am",
	rs:   "([:consonant:] [:sara am:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])

		reg1 := m.Group(1)
		c.FirstConsonant = input[reg1.Start]
		c.Tail = append(c.Tail, input[reg1.Start+1])

		*length = m.Length()
		return true
	},
}

var r_mai_han_akat = TccRule{
	name: "mai_han_akat",
	rs:   "([:consonant: && :mai han akat:]) ([:consonant:])",
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

var r_sara_uee = TccRule{
	name: "sara_uee",
	// we have to distinguish the O ANG from a subsequent
	// O ANG that is followed by a vowel
	rs: "([:consonant: && (:sara uee: || !:diacritic vowel:)])" +
		"([:o ang: && !:diacritic vowel: && !:tone mark:])" +
		"([!:mid position vowel:] | $)",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		reg1 := m.Group(1)
		reg2 := m.Group(2)
		totalLength := reg1.Length() + reg2.Length()
		*c = makeCluster(input[i : i+totalLength])

		c.FirstConsonant = input[reg1.Start]

		c.Tail = append(c.Tail, input[reg2.Start])

		*length = totalLength
		return true
	},
}

// short and long ua
var r_ua = TccRule{
	name: "ua",
	rs:   "([:consonant: && :mai han akat:]) ([:wo waen:] [:sara a:]?)",
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
		c.Tail = append(c.Tail, input[reg2.Start:reg2.End]...)

		*length = m.Length()
		return true
	},
}

var r_maybe_sandwich_sara_a = TccRule{
	name: "maybe_sandwich_sara_a",
	rs: "([:sara e:] | [:sara ae:] | [:sara o:]) " +
		"(?P<consonant>" +
		// BEGIN possible consonants allowed between sandwich vowels
		"([:consonant: && !:diacritic vowel:]) | " +
		"([:bare ho hip:] [:low consonant after ho hip: && !:diacritic vowel:]) | " +
		"([:consonant before gliding lo ling: && !:diacritic vowel:] [:lo ling: && !:diacritic vowel:]) |" +
		"([:consonant before gliding ro rua: && !:diacritic vowel:] [:ro rua: && !:diacritic vowel:])   |" +
		"([:consonant before gliding wo waen: && !:diacritic vowel:] [:wo waen: && !:diacritic vowel:])" +
		// END   possible consonants allowed between sandwich vowels
		")" +
		"(?P<final>$|[(!:mid position vowel:)||:sara a:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}

		totalLength := m.Length()

		finalReg := m.GroupName("final")
		xtra := finalReg.Length()
		hasSaraA := false
		if xtra > 0 {
			if input[finalReg.Start].Main != THAI_CHARACTER_SARA_A {
				totalLength--
			} else {
				hasSaraA = true
			}
		}

		*c = makeCluster(input[i : i+totalLength])

		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		regc := m.GroupName("consonant")
		c.FirstConsonant = input[regc.Start]
		if regc.Length() > 1 {
			c.Tail = append(c.Tail, input[regc.Start+1:regc.End]...)
		}

		if hasSaraA {
			c.Tail = append(c.Tail, input[finalReg.Start])
		}

		*length = totalLength
		return true
	},
}

var r_special_o_ang = TccRule{
	name: "special_o_ang",
	rs: "([:sara e:])" +
		"(?P<consonant>" +
		// BEGIN special
		// special words:
		"([:bare cho ching:][:bare pho phan:])" + // "เฉพาะ"
		// END   special
		")" +
		"(?P<tail>[:sara aa:][:sara a:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		regc := m.GroupName("consonant")
		c.FirstConsonant = input[regc.Start]
		if regc.Length() > 1 {
			c.Tail = append(c.Tail, input[regc.Start+1:regc.End]...)
		}

		regt := m.GroupName("tail")
		c.Tail = append(c.Tail, input[regt.Start:regt.End]...)

		*length = m.Length()
		return true
	},
}

var r_short_o_ang = TccRule{
	name: "short_o_ang",
	rs: "([:sara e:])" +
		"(?P<consonant>" +
		// BEGIN possible consonants allowed between sandwich vowels
		"([:consonant: && !:diacritic vowel:]) | " +
		"([:bare ho hip:] [:low consonant after ho hip: && !:diacritic vowel:]) | " +
		"([:consonant before gliding lo ling: && !:diacritic vowel:] [:lo ling: && !:diacritic vowel:]) |" +
		"([:consonant before gliding ro rua: && !:diacritic vowel:] [:ro rua: && !:diacritic vowel:]) |" +
		"([:consonant before gliding wo waen: && !:diacritic vowel:] [:wo waen: && !:diacritic vowel:]) " +
		// special words:
		"([:bare cho ching:][:bare pho phan:])" + // "เฉพาะ"
		// END   possible consonants allowed between sandwich vowels
		")" +
		"(?P<tail>[:sara aa:][:sara a:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		regc := m.GroupName("consonant")
		c.FirstConsonant = input[regc.Start]
		if regc.Length() > 1 {
			c.Tail = append(c.Tail, input[regc.Start+1:regc.End]...)
		}

		regt := m.GroupName("tail")
		c.Tail = append(c.Tail, input[regt.Start:regt.End]...)

		*length = m.Length()
		return true
	},
}

var r_sandwich_ia = TccRule{
	name: "sandwich_ia",
	rs: "([:sara e:]) " +
		"(?P<consonant>" +
		// BEGIN possible consonants allowed between sandwich vowels
		"([:consonant: && :sara ii:]) | " +
		"([:bare ho hip:] [:low consonant after ho hip: && :sara ii:]) | " +
		"([:consonant before gliding lo ling: && !:diacritic vowel:] [:lo ling: && :sara ii:]) |" +
		"([:consonant before gliding ro rua: && !:diacritic vowel:] [:ro rua: && :sara ii:]) |" +
		"([:consonant before gliding wo waen: && !:diacritic vowel:] [:wo waen: && :sara ii:]) " +
		// END   possible consonants allowed between sandwich vowels
		")" +
		"(?P<tail>[:yo yak:] [:sara a:]?)",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		regc := m.GroupName("consonant")
		c.FirstConsonant = input[regc.Start]
		if regc.Length() > 1 {
			c.Tail = append(c.Tail, input[regc.Start+1:regc.End]...)
		}

		regt := m.GroupName("tail")
		c.Tail = append(c.Tail, input[regt.Start:regt.End]...)

		*length = m.Length()
		return true
	},
}

var r_sandwich_ueea_er = TccRule{
	name: "sandwich_ueea_er",
	rs: "([:sara e:]) " +
		"(?P<consonant>" +
		// BEGIN possible consonants allowed between sandwich vowels
		"([:consonant: && (:sara uee: || !:diacritic vowel:)]) | " +
		"([:bare ho hip:] [:low consonant after ho hip: && (:sara uee: || !:diacritic vowel:)]) | " +
		"([:consonant before gliding lo ling: && !:diacritic vowel:] [:lo ling: && (:sara uee: || !:diacritic vowel:)]) |" +
		"([:consonant before gliding ro rua: && !:diacritic vowel:] [:ro rua: && (:sara uee: || !:diacritic vowel:)]) |" +
		"([:consonant before gliding wo waen: && !:diacritic vowel:] [:wo waen: && (:sara uee: || !:diacritic vowel:)]) " +
		// END   possible consonants allowed between sandwich vowels
		")" +
		//		"(?P<tail>[:o ang:] [:sara a:]?)",
		"(?P<o ang>[:o ang:]) (?P<final>$|[(!:mid position vowel:)||:sara a:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		totalLength := m.Length()

		finalReg := m.GroupName("final")
		xtra := finalReg.Length()
		hasSaraA := false
		if xtra > 0 {
			if input[finalReg.Start].Main != THAI_CHARACTER_SARA_A {
				totalLength--
			} else {
				hasSaraA = true

			}
		}

		*c = makeCluster(input[i : i+totalLength])

		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		regc := m.GroupName("consonant")
		c.FirstConsonant = input[regc.Start]
		if regc.Length() > 1 {
			c.Tail = append(c.Tail, input[regc.Start+1:regc.End]...)
		}

		rego := m.GroupName("o ang")
		c.Tail = append(c.Tail, input[rego.Start])

		if hasSaraA {
			c.Tail = append(c.Tail, input[finalReg.Start])
		}

		*length = totalLength
		return true
	},
}

var r_sara_ai = TccRule{
	name: "sara_ai",
	rs: "([:sara ai maimuan:] | [:sara ai maimalai:]) " +
		"(?P<consonant>" +
		// BEGIN possible consonants allowed between sandwich vowels
		"([:consonant: && (:sara uee: || !:diacritic vowel:)]) | " +
		"([:bare ho hip:] [:low consonant after ho hip: && (:sara uee: || !:diacritic vowel:)]) | " +
		"([:consonant before gliding lo ling: && !:diacritic vowel:] [:lo ling: && (:sara uee: || !:diacritic vowel:)]) (?P<xtra1>$|[!:mid position vowel:]) |" +
		"([:consonant before gliding ro rua: && !:diacritic vowel:] [:ro rua: && (:sara uee: || !:diacritic vowel:)])   (?P<xtra2>$|[!:mid position vowel:]) |" +
		"([:consonant before gliding wo waen: && !:diacritic vowel:] [:wo waen: && (:sara uee: || !:diacritic vowel:)]) (?P<xtra3>$|[!:mid position vowel:])  " +
		// END   possible consonants allowed between sandwich vowels
		")",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		var xtra int
		for _, name := range []string{"xtra1", "xtra2", "xtra3"} {
			xtraReg := m.GroupName(name)
			xtra = xtraReg.Length()
			if xtra > 0 {
				break
			}
		}

		totalLength := m.Length()
		totalLength -= xtra

		*c = makeCluster(input[i : i+totalLength])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		regc := m.GroupName("consonant")
		c.FirstConsonant = input[regc.Start]
		if regc.Length() > 1 {
			tailEnd := regc.End - xtra
			c.Tail = append(c.Tail, input[regc.Start+1:tailEnd]...)
		}

		*length = totalLength
		return true
	},
}

var r_medial_er = TccRule{
	name: "medial_er",
	rs: "([:sara e:]) " +
		"(?P<consonant>" +
		// BEGIN possible consonants allowed between sandwich vowels
		"([:consonant: && :sara i:]) |" +
		"([:bare ho hip:] [:low consonant after ho hip: && :sara i:]) | " +
		"([:consonant before gliding lo ling: && !:diacritic vowel:] [:lo ling: && :sara i:]) |" +
		"([:consonant before gliding ro rua: && !:diacritic vowel:] [:ro rua: && :sara i:]) |" +
		"([:consonant before gliding wo waen: && !:diacritic vowel:] [:wo waen: && :sara i:]) " +
		// END   possible consonants allowed between sandwich vowels
		")",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		regc := m.GroupName("consonant")
		c.FirstConsonant = input[regc.Start]
		if regc.Length() > 1 {
			c.Tail = append(c.Tail, input[regc.Start+1:regc.End]...)
		}

		*length = m.Length()
		return true
	},
}

var r_sandwich_ao = TccRule{
	name: "sandwich_ao",
	rs: "([:sara e:])" +
		"(?P<consonant>" +
		// BEGIN possible consonants allowed between sandwich vowels
		"([:consonant: && !:diacritic vowel:]) | " +
		"([:bare ho hip:] [:low consonant after ho hip: && !:diacritic vowel:]) | " +
		"([:consonant before gliding lo ling: && !:diacritic vowel:] [:lo ling: && !:diacritic vowel:]) |" +
		"([:consonant before gliding ro rua: && !:diacritic vowel:] [:ro rua: && !:diacritic vowel:]) |" +
		"([:consonant before gliding wo waen: && !:diacritic vowel:] [:wo waen: && !:diacritic vowel:]) " +
		// END   possible consonants allowed between sandwich vowels
		")" +
		"(?P<tail>[:sara aa:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		regc := m.GroupName("consonant")
		c.FirstConsonant = input[regc.Start]
		if regc.Length() > 1 {
			c.Tail = append(c.Tail, input[regc.Start+1:regc.End]...)
		}

		regt := m.GroupName("tail")
		c.Tail = append(c.Tail, input[regt.Start:regt.End]...)

		*length = m.Length()
		return true
	},
}

// Single diacritic vowel
var r_single_diacritic_vowel = TccRule{
	name: "single_diacritic_vowel",
	rs: "([:consonant: && " +
		"(:sara i: || :sara ii: || :sara uee: || :sara ue: || :sara u: || :sara uu:) ])",
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

//24319. แอนะล็อก in data/best/encyclopedia.zip(encyclopedia/encyclopedia_00076.txt) line 133 item 10
//	THAI_CHARACTER_SARA_AE, THAI_CHARACTER_O_ANG, THAI_CHARACTER_NO_NU, THAI_CHARACTER_SARA_A, THAI_CHARACTER_LO_LING, THAI_CHARACTER_MAITAIKHU, THAI_CHARACTER_O_ANG, THAI_CHARACTER_KO_KAI

// Do we at least have a bare consonant?
var r_single_consonant = TccRule{
	name: "consonant",
	rs:   "[:consonant: && !:diacritic vowel:]",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		c.FirstConsonant = input[i]
		*length = m.Length()
		return true
	},
}

var r_sanskrit = TccRule{
	name: "sanskrit",
	rs:   "[:ru:] [:lakkhangyao:]",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		c.FirstConsonant = input[i]
		c.Tail = append(c.Tail, input[i+1])
		*length = m.Length()
		return true
	},
}

var r_punctuation_or_digit = TccRule{
	name: "punctuation_or_digit",
	rs:   "[:mid position sign:] | [:digit:]",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		c.SingleMidSign = input[i]
		*length = m.Length()
		return true
	},
}

var r_error_solo_diacritic = TccRule{
	name: "solo_diacritic",
	rs:   "[:diacritic vowel:] | [:tone mark:] | [:upper position sign:]",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		c.SingleMidSign = input[i]
		c.IsValidThai = false
		c.InvalidReason = ReasonSoloDiacritic
		*length = m.Length()
		return true
	},
}

var r_error_sara_e_ae = TccRule{
	name: "error_sara_e_ae",
	rs:   "([:sara e:]|[:sara ae:]) [:consonant: && :diacritic vowel:]",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		c.FrontVowel = input[i]
		c.FirstConsonant = input[i+1]
		c.IsValidThai = false
		c.InvalidReason = ReasonSaraEInvalidCombo
		*length = m.Length()
		return true
	},
}

var r_error_solo_final_vowel = TccRule{
	name: "solo_final_vowel",
	rs:   "[:sara aa:] | [:sara a:] | [:sara am:]",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		c.SingleMidSign = input[i]
		c.IsValidThai = false
		c.InvalidReason = ReasonSoloDiacritic
		*length = m.Length()
		return true
	},
}

// some people mispell it with a sara ae at the beginning
var r_error_short_o_ang = TccRule{
	name: "error_short_o_ang",
	rs: "([:sara ae:])" +
		"(?P<consonant>" +
		// BEGIN possible consonants allowed between sandwich vowels
		"([:consonant: && !:diacritic vowel:]) | " +
		"([:bare ho hip:] [:low consonant after ho hip: && !:diacritic vowel:]) | " +
		"([:consonant before gliding lo ling: && !:diacritic vowel:] [:lo ling: && !:diacritic vowel:]) |" +
		"([:consonant before gliding ro rua: && !:diacritic vowel:] [:ro rua: && !:diacritic vowel:]) |" +
		"([:consonant before gliding wo waen: && !:diacritic vowel:] [:wo waen: && !:diacritic vowel:]) " +
		// special words:
		"([:bare cho ching:][:bare pho phan:])" + // "เฉพาะ"
		// END   possible consonants allowed between sandwich vowels
		")" +
		"(?P<tail>[:sara aa:][:sara a:])",
	ck: func(s *TccRule, input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
		m := s.regex.MatchAt(input, i)
		if !m.Success {
			return false
		}
		*c = makeCluster(input[i : i+m.Length()])
		reg1 := m.Group(1)
		c.FrontVowel = input[reg1.Start]

		regc := m.GroupName("consonant")
		c.FirstConsonant = input[regc.Start]
		if regc.Length() > 1 {
			c.Tail = append(c.Tail, input[regc.Start+1:regc.End]...)
		}

		regt := m.GroupName("tail")
		c.Tail = append(c.Tail, input[regt.Start:regt.End]...)

		c.IsValidThai = false
		c.InvalidReason = ReasonSaraAeInvalidCombo

		*length = m.Length()
		return true
	},
}
