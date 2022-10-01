package gparser

import (
	"fmt"
	// graph "github.com/mcpar-land/generic-graph"
)

// The output type O must conform to the ParserResult interface
type ParserResult interface {
	GetConsumed() int
	SetConsumed(v int)
}

// The holder of the definition of the parser
type Parser[I any, O ParserResult] struct {
	finalizedOk bool
	targetNames []string
	tree        ParserTree[O]

	namespace  map[string]NodeType
	ruleMap    map[string]*ParserRule[O]
	ioClassMap map[string]*IOClass[I, O]
	/*
		ioMap       map[string]I
		ioMapper    func(string, I) O
		ioSeqMap    map[string][]I
		ioSeqMapper func(string, []I) O
	*/
}

func NewParser[I any, O ParserResult](targetNames ...string) *Parser[I, O] {
	s := &Parser[I, O]{}
	s.Initialize(targetNames...)
	return s
}

func (s *Parser[I, O]) Initialize(targetNames ...string) {
	s.assertNotFinalized()
	s.targetNames = targetNames
	s.namespace = make(map[string]NodeType)
	s.ruleMap = make(map[string]*ParserRule[O])
	s.ioClassMap = make(map[string]*IOClass[I, O])
	/*
		s.ioMap = make(map[string]I)
		s.ioSeqMap = make(map[string][]I)
	*/
}

func (s *Parser[I, O]) assertFinalized() {
	if !s.finalizedOk {
		panic("Parser.Finalized() has not been called yet")
	}
}

func (s *Parser[I, O]) assertNotFinalized() {
	if s.finalizedOk {
		panic("Parser.Finalized() has already been called.")
	}
}

func (s *Parser[I, O]) RegisterIOClass(newClass *IOClass[I, O]) {
	s.assertNotFinalized()
	if ot, has := s.namespace[newClass.Name]; has {
		panic(fmt.Sprintf("Parser already has a %s named \"%s\"", NodeTypeName(ot), newClass.Name))
	}
	s.ioClassMap[newClass.Name] = newClass
	s.namespace[newClass.Name] = 'C'
}

func (s *Parser[I, O]) RegisterRule(newRule *ParserRule[O]) {
	s.assertNotFinalized()
	if ot, has := s.namespace[newRule.Name]; has {
		panic(fmt.Sprintf("Parser already has a %s named \"%s\"", NodeTypeName(ot), newRule.Name))
	}

	err := newRule.tokenizePatterns()
	if err != nil {
		panic(fmt.Sprintf("Error parsing %s rule: %s", newRule.Name, err))
	}

	s.ruleMap[newRule.Name] = newRule
	//	s.orderedRules = append(s.orderedRules, newRule)
	s.namespace[newRule.Name] = 'R'
}

/*
func (s *Parser[I, O]) RegisterIOMap(m map[string]I, f func(string, I) O) {
	s.assertNotFinalized()
	if len(s.ioMap) > 0 {
		panic("IOMap can only be registered once")
	}

	for name, _ := range m {
		if ot, has := s.namespace[name]; has {
			panic(fmt.Sprintf("Parser already has a %s named \"%s\"", NodeTypeName(ot), name))
		}
		s.namespace[name] = 'M'
	}
	s.ioMap = m
	s.ioMapper = f
}

func (s *Parser[I, O]) RegisterIOSeqMap(m map[string][]I, f func(string, []I) O) {
	s.assertNotFinalized()
	if len(s.ioSeqMap) > 0 {
		panic("IOSeqMap can only be registered once")
	}

	for name, _ := range m {
		if ot, has := s.namespace[name]; has {
			panic(fmt.Sprintf("Parser already has a %s named \"%s\"", NodeTypeName(ot), name))
		}
		s.namespace[name] = 'S'
	}
	s.ioSeqMap = m
	s.ioSeqMapper = f
}
*/

// This function can be called from with a Rule
// This is only useful during development of gparser itself
func ParserAssertLenEq[T any](inputs []T, size int) {
	if len(inputs) != size {
		panic(fmt.Sprintf("Expected len=%d but got %d", size, len(inputs)))
	}
}
