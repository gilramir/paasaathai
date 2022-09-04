package paasaathai

import (
	"sync"
)

// Based on ideas from, but not the implemention defined in,
// "Character Cluster Based Thai Information Retrieval"
// By Thanaruk Theeramunkong, Virach Sornlertlamvanich,
// Thanasan Tanhermhong, and Wirat Chinnan

type CharacterCluster struct {
	Text string

	IsThai      bool
	IsValidThai bool

	FrontVowel     GlyphStack
	FirstConsonant GlyphStack
	// The Tail might contain consontants and vowels,
	// only consonants, or only vowels.
	Tail []GlyphStack
}

func (s *CharacterCluster) HasFrontVowel() bool {
	return len(s.FrontVowel.Runes) > 0
}

type CharacterClusterParser struct {
	Chan chan *CharacterCluster
	Wg   sync.WaitGroup
}

func ParseCharacterClusters(input []*GlyphStack) []*CharacterCluster {
	var parser CharacterClusterParser
	parser.GoParse(input)

	// 4 is just a guess right now; it has not been checked
	estimatedAllocation := len(input) / 4
	if estimatedAllocation < 1 {
		estimatedAllocation = 1
	}

	clusters := make([]*CharacterCluster, 0, estimatedAllocation)
	for c := range parser.Chan {
		clusters = append(clusters, c)
	}

	parser.Wg.Wait()
	return clusters
}

func (s *CharacterClusterParser) GoParse(input []*GlyphStack) {
	s.Chan = make(chan *CharacterCluster)

	s.Wg.Add(1)
	go s.parse(input)
}

func (s *CharacterClusterParser) parse(input []*GlyphStack) {
	defer close(s.Chan)
	defer s.Wg.Done()

	var cc *CharacterCluster

	for i := 0; i < len(input); {
		gs := input[i]

		if cc == nil {
			// New cluster
			cc = new(CharacterCluster)
			if len(gs.Runes) == 0 {
				panic("GlyphStack has 0 glyphs")
			}
			if !RuneIsThai(gs.Runes[0]) {
				cc.IsThai = false
				cc.Text = string(gs.Runes)
				i++
				continue
			}

			// This is a Thai GlyphStack
			cc.IsThai = true

			// Special cases
			if len(gs.Runes) == 2 {
				// ก็
				if gs.Runes[0] == THAI_CHARACTER_KO_KAI && gs.Runes[1] == THAI_CHARACTER_MAITAIKHU {
					cc.Text = string(gs.Runes)
					cc.IsValidThai = true
					cc.FirstConsonant = *gs
					s.Chan <- cc
					cc = nil
				}
			}
		} else {
			// In-progress cluster
		}
		i++
	}
}
