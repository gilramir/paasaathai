package gparser

import (
	. "gopkg.in/check.v1"
)

var vowels = []rune{'A', 'E', 'I', 'O', 'U'}
var consonants = []rune{'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K',
	'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z'}
var digits = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

var IOClassVowel = &IOClass[rune, *TASTNode]{
	"Vowel",
	func(r rune) (bool, *TASTNode) {
		for _, t := range vowels {
			if r == t {
				values := make([]rune, 1)
				values[0] = r
				return true, &TASTNode{
					values: values,
				}
			}
		}
		return false, nil
	},
}

var IOClassConsonant = &IOClass[rune, *TASTNode]{
	"Consonant",
	func(r rune) (bool, *TASTNode) {
		for _, t := range consonants {
			if r == t {
				values := make([]rune, 1)
				values[0] = r
				return true, &TASTNode{
					values: values,
				}
			}
		}
		return false, nil
	},
}

/*
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

var RuleTwoVowels = &ParserRule[*TASTNode]{
	Name: "TwoVowels",
	Patterns: []*ParserRulePattern[*TASTNode]{
		&ParserRulePattern[*TASTNode]{
			//Pattern: "Vowel{2}",
			Pattern: "Vowel{2} A B C",
			Matched: func(inputs []*TASTNode) *TASTNode {
				ParserAssertLenInputs(inputs, 2)
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
*/

type TASTNode struct {
	values   []rune
	consumed int
}

func (s *TASTNode) GetConsumed() int {
	return s.consumed
}

func (s *TASTNode) SetConsumed(v int) {
	s.consumed = v
}

// need to do the rule pattern tokenizatoin after everything is registered, i
// nthe finalzie step.

func (s *MySuite) TestIOClass01(c *C) {
	var parser Parser[rune, *TASTNode]
	parser.Initialize("Vowel")
	parser.RegisterIOClass(IOClassVowel)
	/*
		parser.RegisterRule(RuleTwoVowels)
		parser.RegisterIOMap(ioMap, ioMapper)
		parser.RegisterIOSeqMap(ioSeqMap, ioSeqMapper)
	*/
	parser.Finalize()

	input := []rune{'A'}
	results, err := parser.Parse(input)
	c.Assert(err, IsNil)
	c.Assert(results, HasLen, 1)
	c.Check(results[0].consumed, Equals, 1)
}

func (s *MySuite) TestIOClass02(c *C) {
	var parser Parser[rune, *TASTNode]
	parser.Initialize("Vowel", "Consonant")
	parser.RegisterIOClass(IOClassVowel)
	parser.RegisterIOClass(IOClassConsonant)
	/*
		parser.RegisterRule(RuleTwoVowels)
		parser.RegisterIOMap(ioMap, ioMapper)
		parser.RegisterIOSeqMap(ioSeqMap, ioSeqMapper)
	*/
	parser.Finalize()

	input := []rune{'A'}
	results, err := parser.Parse(input)
	c.Assert(err, IsNil)
	c.Assert(results, HasLen, 1)
	c.Check(results[0].consumed, Equals, 1)
}
