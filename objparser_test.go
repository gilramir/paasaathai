package paasaathai

import (
	"strings"

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

var ioMap = map[string]string{
	"A": "A",
	"B": "B",
	"C": "C",
}

func ioMapper(name string, inputTarget string) *TASTNode {
	values := make([]string, 1)
	values[0] = string(inputTarget)
	return &TASTNode{
		values:   values,
		consumed: 1,
	}
}

var ioSeqMap = map[string][]string{
	"ABC": []string{"A", "B", "C"},
	"abc": []string{"a", "b", "c"},
}

func ioSeqMapper(name string, inputTarget []string) *TASTNode {
	// We combine them into one;
	values := make([]string, 0)
	values[0] = strings.Join(inputTarget, "")
	return &TASTNode{
		values:   values,
		consumed: len(inputTarget),
	}
}

var RuleTwoVowels = &ObjParserRule[*TASTNode]{
	Name: "TwoVowels",
	Patterns: []*ObjParserRulePattern[*TASTNode]{
		&ObjParserRulePattern[*TASTNode]{
			//Pattern: "Vowel{2}",
			Pattern: "Vowel{2} A B C",
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

// need to do the rule pattern tokenizatoin after everything is registered, i
// nthe finalzie step.

func (s *MySuite) TestOP01(c *C) {
	var parser ObjParser[string, *TASTNode]
	parser.Initialize("TwoVowels", "Vowel")
	parser.RegisterLeafClass(LeafClassVowel)
	parser.RegisterRule(RuleTwoVowels)
	parser.RegisterIOMap(ioMap, ioMapper)
	parser.RegisterIOSeqMap(ioSeqMap, ioSeqMapper)
	parser.Finalize()

	input := []string{"A", "A"}
	results, err := parser.Parse(input)
	c.Assert(err, IsNil)
	c.Assert(results, HasLen, 1)
	c.Check(results[0].consumed, Equals, 2)
}
