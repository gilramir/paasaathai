package gparser

import (
	"fmt"
	"strings"
)

type TrialStack[I any, O ParserResult] struct {
	inputPosition int
	nodeStack     []*ParserTree[O]
}

func (s *TrialStack[I, O]) Dump() {
	nodeStrings := make([]string, len(s.nodeStack))
	for i, ns := range s.nodeStack {
		nodeStrings[i] = ns.Repr()
	}
	fmt.Printf("<TrialStack iPos=%d nodes=%s>\n", s.inputPosition, strings.Join(nodeStrings, ", "))
}

// -----------------------------------------------------------------------------------

// Parse the input and return all possible outputs, or an error
func (s *Parser[I, O]) Parse(input []I) ([]O, error) {
	//	var pstate objParserState[I, O]
	//	pstate.Initialize(s)

	trialStacks := make([]*TrialStack[I, O], 1)
	firstTs := &TrialStack[I, O]{
		nodeStack: make([]*ParserTree[O], 1),
	}
	firstTs.nodeStack[0] = &s.tree
	trialStacks[0] = firstTs

	results := make([]O, 0)

	for len(trialStacks) > 0 {
		ts := trialStacks[0]
		trialStacks = trialStacks[1:]

		ts.Parse(input, &trialStacks, &results)
	}

	return results, nil
}

func (s *TrialStack[I, O]) Parse(input []I, trialStacks *[]*TrialStack[I, O], results *[]O) {

}

// -----------------------------------------------------------------------------------

type objParserState[I any, O ParserResult] struct {
	objParser *Parser[I, O]
	input     []I
	results   []O
}

func (s *objParserState[I, O]) Initialize(parser *Parser[I, O]) {
	s.objParser = parser
}

func (s *objParserState[I, O]) Parse(input []I) ([]O, error) {
	s.input = input
	s.results = make([]O, 0, 1)
	/*
		// run through the tree of rule/classes/ioClasses
		trialNodes := make([]*objParserTrialNode[I, O], 0, len(s.objParser.targetNames))

		for _, targetName := range s.objParser.targetNames {
			if r, has := s.objParser.ruleMap[targetName]; has {
				for pi, p := range r.Patterns {
					trialNode := &objParserTrialNode[I, O]{
						rule:         r,
						rulePattern:  p,
						rulePatternN: pi,
						children:     make([]*objParserTrialNode[I, O], 0),
					}
					trialNodes = append(trialNodes, trialNode)
				}
				continue
			}
			if c, has := s.objParser.ioClassMap[targetName]; has {
				trialNode := &objParserTrialNode[I, O]{
					ioClass: c,
					children:  make([]*objParserTrialNode[I, O], 0),
				}
				trialNodes = append(trialNodes, trialNode)
				continue
			}
		}
	*/

	return s.results, nil
}

/*
type objParserTrialNode[I any, O ParserResult] struct {
	rule         *ParserRule[O]
	rulePattern  *ParserRulePattern[O]
	rulePatternN int
	class        *ParserClass[O]
	ioClass    *ParserIOClass[I, O]

	children []*objParserTrialNode[I, O]
}
*/
