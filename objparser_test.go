package paasaathai

import (
	. "gopkg.in/check.v1"
)

var vowels = []string{"A", "E", "I", "O", "U"}
var consonants = []string{"B", "C", "D", "F", "G", "H", "J", "K",
	"L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z"}
var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

var LeafClassVowel = &ObjParserLeafClass[string, *TASTNode]{
	"Vowel",
	func(s string) (bool, *TASTNode) {
		for _, t := range vowels {
			if s == t {
				values := make([]string, 1)
				values[0] = s
				return true, &TASTNode{
					values: values,
				}
			}
		}
		return false, nil
	},
}

var RuleTwoVowels = &ObjParserRule[*TASTNode]{
	Name: "TwoVowels",
	Patterns: []*ObjParserRulePattern[*TASTNode]{
		&ObjParserRulePattern[*TASTNode]{
			Pattern: "Vowel{2}",
			Matched: func(inputs []*TASTNode) *TASTNode {
				ObjParserAssertLenInputs(inputs, 2)
				values := make([]string, len(inputs[0].values)+len(inputs[1].values))
				copy(values, inputs[0].values)
				copy(values[len(inputs[0].values):], inputs[1].values)
				return &TASTNode{
					values: values,
				}
			},
		},
	},
}

type TASTNode struct {
	values   []string
	consumed int
}

func (s *TASTNode) SetConsumed(v int) {
	s.consumed = v
}

func (s *MySuite) TestOP01(c *C) {
	var parser ObjParser[string, *TASTNode]
	parser.Initialize("TwoVowels")
	parser.RegisterLeafClass(LeafClassVowel)
	parser.RegisterRule(RuleTwoVowels)
	parser.Finalize()

	input := []string{"A", "A"}
	results, err := parser.Parse(input)
	c.Assert(err, IsNil)
	c.Assert(results, HasLen, 1)
	c.Check(results[0].consumed, Equals, 2)
}
