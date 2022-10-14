package paasaathai

// Based on ideas from, but not the implemention defined in,
// "Character Cluster Based Thai Information Retrieval"
// By Thanaruk Theeramunkong, Virach Sornlertlamvanich,
// Thanasan Tanhermhong, and Wirat Chinnan

import (
	"fmt"
)

var ConsonantsAllowedAtFront = []rune{
	'ก', 'ข', 'ค', 'ง', 'จ',
	'ฉ', 'ช', 'ซ', 'ญ', 'ด',
	'ต', 'ถ', 'ท', 'ธ', 'น',
	'บ', 'ป', 'ผ', 'ฝ', 'พ',
	'ฟ', 'ภ', 'ม', 'ย', 'ร',
	'ล', 'ว', 'ศ', 'ส', 'ห',
	'อ',
}
var ConsonantsAllowedAfterOAng = []rune{
	'ก', 'ง', 'ด', 'ต', 'ธ',
	'น', 'ม', 'ย', 'ร', 'ว',
}

var ConsonantsAllowedAfterHoHip = []rune{
	'ก', 'ญ', 'ด', 'น', 'ม',
	'ย', 'ร', 'ล', 'ว',
}

type GStackCluster struct {
	Text string

	IsThai        bool
	IsValidThai   bool
	InvalidReason string

	SingleMidSign  GraphemeStack
	FrontVowel     GraphemeStack
	FirstConsonant GraphemeStack
	// The Tail might contain consontants and vowels,
	// only consonants, or only vowels.
	Tail []GraphemeStack
}

func (s *GStackCluster) Repr() string {
	if s.IsThai {
		if s.IsValidThai {
			if len(s.SingleMidSign.Runes) > 0 {
				return fmt.Sprintf("<CC Thai %s SS:%s", s.Text, s.SingleMidSign.Repr())
			} else {
				result := fmt.Sprintf("<CC Thai %s", s.Text)
				if len(s.FrontVowel.Runes) > 0 {
					result += fmt.Sprintf(" FV:%s", s.FrontVowel.Repr())
				}
				if len(s.FirstConsonant.Runes) > 0 {
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

/*
// A consonant that can appear in the first position
var FirstConsonant = &ObjParserClass[*GraphemeStack]{
	"FirstConsonant",
	func(g *GraphemeStack) bool {
		gr := g.Runes[0]
		if !RuneIsConsonant(gr) {
			return false
		}
		// TODO - make this a set
		for _, r := range ConsonantsAllowedAtFront {
			if gr == r {
				return true
			}
		}
		return false
	},
}

var RuleCRRC = &ObjParserPattern[*CC]{
    "CRRC",
    "FirstConsonant RoRua{2} FinalConsonant"
}

var parser = NewObjParser[*GraphemeStack]()

func init() {
	parser.Initialize()
	parser.Register(FirstConsonant)
	parser.AddRules(`
	LEGAL = CC

	CC = FirstConsonant RoRua{2} FinalConsonant
	{
		RR(
`)
}

*/

/*
func (s *GStackCluster) HasFrontVowel() bool {
	return len(s.FrontVowel.Runes) > 0
}

type GStackClusterParser struct {
	Chan chan *GStackCluster
	Wg   sync.WaitGroup

	input []*GraphemeStack
	pos   int
	cc    *GStackCluster
}
*/
/*
func ParseGStackClusters(input string) []*GStackCluster {
	glyphStacks := ParseGraphemeStacks(input)
	return ParseGStackClustersFromGraphemeStacks(glyphStacks)
}

func ParseGStackClustersFromGraphemeStacks(input []*GraphemeStack) []*GStackCluster {
	var parser GStackClusterParser
	parser.GoParse(input)

	// 4 is just a guess right now; it has not been checked
	estimatedAllocation := len(input) / 4
	if estimatedAllocation < 1 {
		estimatedAllocation = 1
	}

	clusters := make([]*GStackCluster, 0, estimatedAllocation)
	for c := range parser.Chan {
		clusters = append(clusters, c)
	}

	parser.Wg.Wait()
	return clusters
}

func (s *GStackClusterParser) GoParse(input []*GraphemeStack) {
	s.Chan = make(chan *GStackCluster)
	s.input = input
	s.pos = 0
	s.cc = nil

	s.Wg.Add(1)

	go s.parse(input)
}

type stateFunc func() stateFunc

func (s *GStackClusterParser) parse(input []*GraphemeStack) {
	defer close(s.Chan)
	defer s.Wg.Done()

	for stateFunc := s.parseNew; stateFunc != nil; {
		stateFunc()
	}

	if s.cc != nil {
		panic(fmt.Sprintf("end of parse but cc is %s", cc.Repr()))
	}
*/
/*
		// In the middle of a parse?
		switch ccState {
		case ccNew:
			panic(fmt.Sprintf("At end, cc is nil, but ccState is %v", ccState))
		case ccExpectingFirstConsonant:
			cc.IsValidThai = false
		default:
			cc.IsValidThai = true
		}
		s.Chan <- cc
		cc = nil
	}
*/
/*
}

// Start a new GStackCluster
func (s *GStackClusterParser) parseNew() stateFunc {
	if s.cc != nil {
		s.Chan <- cc
		cc = nil
	}

	// End of input?
	if s.pos >= len(input) {
		return nil
	}

	s.cc = new(GStackCluster)
	gs := input[s.pos]

	if gs.IsThai() {
		return s.parseNewThai
	} else {
		return s.parseNewNonThai
	}
}

// The first glyph is not Thai
func (s *GStackClusterParser) parseNewNonThai() stateFunc {
	gs := input[s.pos]
	s.ccText = string(gs.Runes)
	s.cc.IsThai = false
	s.pos++

	return s.parseNew
}

// The first glyph is Thai
func (s *GStackClusterParser) parseNewThai() stateFunc {
	// This is a Thai GraphemeStack
	s.cc.IsThai = true

	gs := input[s.pos]
	// Is this a front vowel?
	if len(gs.Runes) == 1 {
		return s.parseNewThaiFirstGlyph1Rune
	} else {
		return s.parseNewThaiFirstGlyphManyRunes
	}
}

func (s *GStackClusterParser) invalidateOverLength() bool {
	if s.pos > len(s.input) {
		s.cc.InvalidThai = true
		return true
	}
	return false
}

// If the first GraphemeStack has 1 runes in it, then it could be anything.
func (s *GStackClusterParser) parseNewThaiFirstGlyph1Rune() stateFunc {
	gs := input[s.pos]
	if RuneIsFrontPositionVowel(gs.Runes[0]) {
		s.cc.FrontVowel = *gs
		s.pos++
		// We *must* have more input
		if s.invalidateOverLength() {
			s.cc.InvalidReason = fmt.Sprintf("End of string after a front vowel: %s", gs.Repr())
			return nil
		}
		return s.parseAfterFrontVowel
	} else if RuneIsConsonant(gs.Runes[0]) {
		s.cc.FirstConsonant = *gs
		self.pos++
		// We *must* have more input
		if s.invalidateOverLength() {
			s.cc.InvalidReason = fmt.Sprintf("End of string after a front consonant with no upper/lower vowel: %s", gs.Repr())
			return nil
		}
		// Switch on the first consonant rune
		switch gs.Runes[0][0] {
		case THAI_CHARACTER_HO_HIP:
			return s.parseAfterFirstConsonantHoHip
		case THAI_CHARACTER_O_ANG:
			return s.parseAfterFirstConsonantOAng
		default:
			return s.parseAfterFirstConsonant
		}
	} else if RuneIsMidPositionSign(gs.Runes[0]) {
		s.cc.SingleMidSign = *gs
		s.cc.IsValidThai = true
		s.pos++
		return s.parseNew
	} else {
		cc.IsValidThai = false
		cc.InvalidReason = fmt.Sprintf("First glyph is not a front-vowel, consonant, or mid-sign: %s", gs.Repr())
		s.pos++
		return s.parseNew
	}
}

// If the first GraphemeStack has multiple runes in it, then it must be a
// consonant
func (s *GStackClusterParser) parseNewThaiFirstGlyphManyRunes() stateFunc {
	gs := input[s.pos]
	// More than one rune in the GraphemeStack?
	// A GraphemeStack with >1 rune at the beinning of a
	// GStackCluster *must* start with at
	// consonant
	if RuneIsConsonant(gs.Runes[0]) {
		s.cc.FirstConsonant = *gs
		// The next glyphs will be vowels
		// and/or tone marks. The GraphemeStack
		// parser assures it.
		s.pos++
		// Don't check for HO_HIP or O_ANG here; if they are seen here
		// then they have a vowel or tone mark on them, so they are not
		// the class-changing versions of themselves; they're just
		// consonants.
		return s.parseAfterFirstConsonantExpectVowel
	} else {
		cc.IsValidThai = false
		cc.InvalidReason = fmt.Sprintf("First glyph has many runes but no consonant: %s", gs.Repr())
		s.pos++
		return s.parseNew
	}
}

// After a front vowel, we expect a consonant
func (s *GStackClusterParser) parseAfterFrontVowel() stateFunc {
	gs := input[s.pos]
	if RuneIsConsonant(gs.Runes[0]) {
		s.cc.FirstConsonant = *gs
		s.pos++
		if len(gs.Runes) == 1 {
			// Switch on the first consonant rune
			switch gs.Runes[0][0] {
			case THAI_CHARACTER_HO_HIP:
				return s.parseAfterFirstConsonantHoHip
			case THAI_CHARACTER_O_ANG:
				return s.parseAfterFirstConsonantOAng
			default:
				return s.parseAfterFirstConsonant
			}
		} else {
			// Don't check for HO_HIP or O_ANG here; if they are seen here
			// then they have a vowel or tone mark on them, so they are not
			// the class-changing versions of themselves; they're just
			// consonants.
			return s.parseAfterFirstConsonant
		}
	} else {
		cc.IsValidThai = false
		cc.InvalidReason = fmt.Sprintf("Glyph after front-vowel is not a consonant: %s", gs.Repr())
		s.pos++
		return s.parseNew
	}
}
*/

/*
	// Special cases
		if len(gs.Runes) == 2 {
			text := string(gs.Rune)
			if text == "ก็" || text == "อึ" {
				cc.Text = string(gs.Runes)
				cc.IsValidThai = true
				cc.FirstConsonant = *gs
				s.Chan <- cc
				cc = nil
				continue
			}
		}
*/
