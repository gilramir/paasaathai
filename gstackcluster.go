package paasaathai

// Based on ideas from, but not the implemention defined in,
// "Character Cluster Based Thai Information Retrieval"
// By Thanaruk Theeramunkong, Virach Sornlertlamvanich,
// Thanasan Tanhermhong, and Wirat Chinnan
//
// https://www.researchgate.net/publication/2853284_Character_Cluster_Based_Thai_Information_Retrieval
// and more specifically,
// https://www.researchgate.net/profile/Virach-Sornlertlamvanich/publication/2853284_Character_Cluster_Based_Thai_Information_Retrieval/links/02e7e514db194bcb1f000000/Character-Cluster-Based-Thai-Information-Retrieval.pdf

import (
	"fmt"

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

	re_gaaw              *objregexp.Regexp[GraphemeStack]
	re_single_consonant  *objregexp.Regexp[GraphemeStack]
	re_final_pos_short_1 *objregexp.Regexp[GraphemeStack]
	re_final_pos_long_1  *objregexp.Regexp[GraphemeStack]
	re_sara_am           *objregexp.Regexp[GraphemeStack]
}

func (s *GStackClusterParser) Initialize() {
	s.compiler.Initialize()

	// regex classes
	s.compiler.MakeClass("gaaw",
		func(gs GraphemeStack) bool {
			return gs.String() == "ก็"
		})

	s.compiler.MakeClass("consonant",
		func(gs GraphemeStack) bool {
			return RuneIsConsonant(gs.Main)
		})

	s.compiler.MakeClass("consonant-no-diacritic-vowel",
		func(gs GraphemeStack) bool {
			return RuneIsConsonant(gs.Main) && gs.DiacriticVowel == 0
		})

	s.compiler.MakeClass("sliding-consonant",
		func(gs GraphemeStack) bool {
			return GlidingConsonants.Has(gs.Main)
		})

	// regex identity classes
	s.compiler.AddIdentity("sara_a",
		MustParseSingleGraphemeStack(
			string(THAI_CHARACTER_SARA_A)))

	s.compiler.AddIdentity("sara_e",
		MustParseSingleGraphemeStack(
			string(THAI_CHARACTER_SARA_E)))

	s.compiler.AddIdentity("sara_ae",
		MustParseSingleGraphemeStack(
			string(THAI_CHARACTER_SARA_AE)))

	s.compiler.AddIdentity("sara_o",
		MustParseSingleGraphemeStack(
			string(THAI_CHARACTER_SARA_O)))

	s.compiler.AddIdentity("sara_aa",
		MustParseSingleGraphemeStack(
			string(THAI_CHARACTER_SARA_AA)))

	s.compiler.AddIdentity("sara_am",
		MustParseSingleGraphemeStack(
			string(THAI_CHARACTER_SARA_AM)))

	s.compiler.AddIdentity("o_ang",
		MustParseSingleGraphemeStack(
			string(THAI_CHARACTER_O_ANG)))

	s.compiler.Finalize()

	var text string

	// Final position vowels; due to how they are spelled,
	// we know they must be final position
	// Thai for Beginners, pg. 250
	// Thai for Beginners, pg. 243

	// Short vowel sara a patterns
	text = "([:sara_e:] | [:sara_ae:] | [:sara_o:])? " +
		"([:consonant-no-diacritic-vowel:]) ([:sliding-consonant:])? " +
		"([:sara_a:])"
	s.re_final_pos_short_1 = s.compiler.MustCompile(text)

	// Long vowel patterns; these must be checked after the short vowels
	text = "([:consonant-no-diacritic-vowel:]) ([:sliding-consonant:])? " +
		"([:sara_aa:] | [:o_ang:])"
	s.re_final_pos_long_1 = s.compiler.MustCompile(text)

	// SARA_AM... can this be combined with re_final_pos_long_1 ?
	text = "([:consonant-no-diacritic-vowel:]) ([:sara_am:])"
	s.re_sara_am = s.compiler.MustCompile(text)

	// Consonant Vowel patterns not found from the re_final_pos*
	// regexes
	text = "[:consonant:]"
	s.re_gaaw = s.compiler.MustCompile(text)

	//
	// Thai for Beginners, pg. 243

	text = "[:gaaw:]"
	s.re_gaaw = s.compiler.MustCompile(text)

	text = "[:consonant:]"
	s.re_single_consonant = s.compiler.MustCompile("[:consonant:]")
}

func (s *GStackClusterParser) assertRegisterLength(reg objregexp.Range, length int) {
	if reg.Length() != length {
		panic(fmt.Sprintf("Register expected to have length %d. Got: %+v",
			length, reg))
	}
}

func (s *GStackClusterParser) ParseGraphemeStacks(input []GraphemeStack) []GStackCluster {

	// Estimate the space by dividing the number of UTF-8 bytes
	// by 3 (each Thai code point needs 3), and then taking 2/3 of that
	estimatedAllocation := len(input) * 2 / 3
	clusters := make([]GStackCluster, 0, estimatedAllocation)

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

		// check short vowels which indicate the end of the cluster
		// (short in vowel length; longer in # of GraphemeStacks)
		if s.ck_final_pos_short_1(input, i, &length, &c) ||
			// check long vowels which indicate the end of the
			// cluster
			s.ck_final_pos_long_1(input, i, &length, &c) ||

			// C + SARA_AM
			s.ck_sara_am(input, i, &length, &c) ||

			// Just a gaaw?
			s.ck_gaaw(input, i, &length, &c) ||

			// Just a single consonant?
			s.ck_single_consonant(input, i, &length, &c) {

			clusters = append(clusters, c)
			i += length
			continue
		}

		// No success.
		msg := fmt.Sprintf("Can't handle gstack at pos %d: %s", i, input[i].Repr())
		panic(msg)
	}

	return clusters
}

// Check for final position short vowel patterns
// "([:sara_e:] | [:sara_ae:] | [:sara_o:])? " +
// "([:consonant-no-diacritic-vowel:] ([:sliding-consonant:]?) " +
// "([:sara_a:])"
func (s *GStackClusterParser) ck_final_pos_short_1(input []GraphemeStack, i int, length *int, c *GStackCluster) bool {

	m := s.re_final_pos_short_1.MatchAt(input, i)
	if !m.Success {
		return false
	}
	*c = makeCluster(input[i : i+m.Length()])
	reg1 := m.Register(1)
	if !reg1.Empty() {
		s.assertRegisterLength(reg1, 1)
		c.FrontVowel = input[reg1.Start]
	}

	reg2 := m.Register(2)
	c.FirstConsonant = input[reg2.Start]

	reg3 := m.Register(3)
	if !reg3.Empty() {
		c.Tail = append(c.Tail, input[reg3.Start:reg3.End]...)
	}
	reg4 := m.Register(4)
	c.Tail = append(c.Tail, input[reg4.Start:reg4.End]...)

	*length = m.Length()
	return true
}

// check for final position long vowel patterns
//"([:consonant-no-diacritic-vowel:]) ([:sliding-consonant:])? " +
//		"([:sara_aa:] | [:o_ang:])"
func (s *GStackClusterParser) ck_final_pos_long_1(input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
	m := s.re_final_pos_long_1.MatchAt(input, i)
	if !m.Success {
		return false
	}
	*c = makeCluster(input[i : i+m.Length()])
	reg1 := m.Register(1)
	c.FirstConsonant = input[reg1.Start]

	if m.HasRegister(2) {
		reg2 := m.Register(2)
		c.Tail = append(c.Tail, input[reg2.Start:reg2.End]...)
	}
	reg3 := m.Register(3)
	c.Tail = append(c.Tail, input[reg3.Start:reg3.End]...)

	*length = m.Length()
	return true
}

// check for C + SARA_AM
// "([:consonant-no-diacritic-vowel:]) ([:sara_am:])"
func (s *GStackClusterParser) ck_sara_am(input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
	m := s.re_sara_am.MatchAt(input, i)
	if !m.Success {
		return false
	}
	*c = makeCluster(input[i : i+2])
	reg1 := m.Register(1)
	c.FirstConsonant = input[reg1.Start]

	reg2 := m.Register(2)
	c.Tail = append(c.Tail, input[reg2.Start])

	*length = m.Length()
	return true
}

// Just a gaww?
func (s *GStackClusterParser) ck_gaaw(input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
	m := s.re_gaaw.MatchAt(input, i)
	if !m.Success {
		return false
	}
	*c = makeCluster(input[i : i+1])
	c.FirstConsonant = input[i]
	*length = m.Length()
	return true
}

// Do we at least have a consonant?
func (s *GStackClusterParser) ck_single_consonant(input []GraphemeStack, i int, length *int, c *GStackCluster) bool {
	m := s.re_single_consonant.MatchAt(input, i)
	if !m.Success {
		return false
	}
	*c = makeCluster(input[i : i+1])
	c.FirstConsonant = input[i]
	*length = m.Length()
	return true
}

var ConsonantsAllowedAtFront = NewSetFromSlice([]rune{
	'ก', 'ข', 'ค', 'ง', 'จ',
	'ฉ', 'ช', 'ซ', 'ญ', 'ด',
	'ต', 'ถ', 'ท', 'ธ', 'น',
	'บ', 'ป', 'ผ', 'ฝ', 'พ',
	'ฟ', 'ภ', 'ม', 'ย', 'ร',
	'ล', 'ว', 'ศ', 'ส', 'ห',
	'อ',
})

// The initial O_ANG changes a subsequent
// consonant into being a MID class consonant.
var ConsonantsAllowedAfterOAng = NewSetFromSlice([]rune{
	'ก', 'ง', 'ด', 'ต', 'ธ',
	'น', 'ม', 'ย', 'ร', 'ว',
})

// The initial HO_HIP changes a subsequent
// consonant into being a HIGH class consonant.
var ConsonantsAllowedAfterHoHip = NewSetFromSlice([]rune{
	'ก', 'ญ', 'ด', 'น', 'ม',
	'ย', 'ร', 'ล', 'ว',
})

// Consonants allowed after other consonants to make a single
// sound
var GlidingConsonants = NewSetFromSlice([]rune{
	'ล',
})
