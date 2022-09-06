package paasaathai

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	// graph "github.com/mcpar-land/generic-graph"
)

// The output type O must conform to the ObjParserResult interface
type ObjParserResult interface {
	SetConsumed(v int)
}

// The holder of the definition of the parser
type ObjParser[I any, O ObjParserResult] struct {
	targetNames  []string
	orderedRules []*ObjParserRule[O]
	ruleMap      map[string]*ObjParserRule[O]
	classMap     map[string]*ObjParserClass[O]
	leafClassMap map[string]*ObjParserLeafClass[I, O]
	allItems     Set[string]
	finalizedOk  bool
}

func NewObjParser[I any, O ObjParserResult](targetNames ...string) *ObjParser[I, O] {
	s := &ObjParser[I, O]{}
	s.Initialize(targetNames...)
	return s
}

func (s *ObjParser[I, O]) Initialize(targetNames ...string) {
	s.assertNotFinalized()
	s.targetNames = targetNames
	s.orderedRules = make([]*ObjParserRule[O], 0)
	s.ruleMap = make(map[string]*ObjParserRule[O])
	s.classMap = make(map[string]*ObjParserClass[O])
	s.leafClassMap = make(map[string]*ObjParserLeafClass[I, O])
}

func (s *ObjParser[I, O]) assertFinalized() {
	if !s.finalizedOk {
		panic("ObjParser.Finalized() has not been called yet")
	}
}
func (s *ObjParser[I, O]) assertNotFinalized() {
	if s.finalizedOk {
		panic("ObjParser.Finalized() has already been called.")
	}
}

func (s *ObjParser[I, O]) Finalize() {
	s.assertNotFinalized()

	s.allItems = NewSet[string]()

	// Create the set of all rule/class/leafClass names
	for name, _ := range s.leafClassMap {
		s.allItems.Add(name)
	}
	for name, _ := range s.classMap {
		s.allItems.Add(name)
	}
	for name, _ := range s.ruleMap {
		s.allItems.Add(name)
	}

	// Sanity-check the targets
	for _, name := range s.targetNames {
		if !s.allItems.Has(name) {
			panic(fmt.Sprintf("Did not find a leaf/class/rule for target \"%s\"", name))
		}
	}

	// Sanity-check that the rules reference things that exists
	for rname, rule := range s.ruleMap {
		for _, pname := range rule.getAllPatternTokenNames() {
			if !s.allItems.Has(pname) {
				panic(fmt.Sprintf("Rule \"%s\" references item \"%s\" but it is not defined", rname, pname))
			}
		}
	}

	s.finalizedOk = true
}

func (s *ObjParser[I, O]) RegisterLeafClass(newClass *ObjParserLeafClass[I, O]) {
	s.assertNotFinalized()
	if _, has := s.leafClassMap[newClass.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a leaf-class named \"%s\"", newClass.Name))
	}
	if _, has := s.ruleMap[newClass.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a rule named \"%s\"; can't name a leaf-class the same", newClass.Name))
	}
	if _, has := s.classMap[newClass.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a class named \"%s\"; can't name a leaf-class the same", newClass.Name))
	}
	s.leafClassMap[newClass.Name] = newClass
}

func (s *ObjParser[I, O]) RegisterClass(newClass *ObjParserClass[O]) {
	s.assertNotFinalized()
	if _, has := s.classMap[newClass.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a class named \"%s\"", newClass.Name))
	}
	if _, has := s.ruleMap[newClass.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a rule named \"%s\"; can't name a class the same", newClass.Name))
	}
	if _, has := s.leafClassMap[newClass.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a leaf-class named \"%s\"; can't name a class the same", newClass.Name))
	}
	s.classMap[newClass.Name] = newClass
}

func (s *ObjParser[I, O]) RegisterRule(newRule *ObjParserRule[O]) {
	s.assertNotFinalized()
	if _, has := s.ruleMap[newRule.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a rule named \"%s\"", newRule.Name))
	}
	if _, has := s.classMap[newRule.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a class named \"%s\"; can't name a rule the same", newRule.Name))
	}
	if _, has := s.leafClassMap[newRule.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a leaf-class named \"%s\"; can't name a rule the same", newRule.Name))
	}

	err := newRule.tokenizePatterns()
	if err != nil {
		panic(fmt.Sprintf("Error parsing %s rule: %s", newRule.Name, err))
	}

	s.ruleMap[newRule.Name] = newRule
	s.orderedRules = append(s.orderedRules, newRule)
}

// Parse the input and return all possible outputs, or an error
func (s *ObjParser[I, O]) Parse(input []I) ([]O, error) {
	var pstate objParserState[I, O]
	pstate.Initialize(s)
	return pstate.Parse(input)
}

// -----------------------------------------------------------------------------------

type objParserState[I any, O ObjParserResult] struct {
	objParser *ObjParser[I, O]
	input     []I
	results   []O
}

func (s *objParserState[I, O]) Initialize(parser *ObjParser[I, O]) {
	s.objParser = parser
}

func (s *objParserState[I, O]) Parse(input []I) ([]O, error) {
	s.input = input
	s.results = make([]O, 0, 1)

	// run through the tree of rule/classes/leafClasses
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
		if c, has := s.objParser.leafClassMap[targetName]; has {
			trialNode := &objParserTrialNode[I, O]{
				leafClass: c,
				children:  make([]*objParserTrialNode[I, O], 0),
			}
			trialNodes = append(trialNodes, trialNode)
			continue
		}
	}

	/*
		for _, trialNode := range trialNodes {
			trialNode.Try(s, 0)
		}
	*/

	return s.results, nil
}

type objParserTrialNode[I any, O ObjParserResult] struct {
	rule         *ObjParserRule[O]
	rulePattern  *ObjParserRulePattern[O]
	rulePatternN int
	class        *ObjParserClass[O]
	leafClass    *ObjParserLeafClass[I, O]

	children []*objParserTrialNode[I, O]
}

// -----------------------------------------------------------------------------------
func ObjParserAssertLenInputs[O ObjParserResult](inputs []O, size int) {
	if len(inputs) != size {
		panic(fmt.Sprintf("Expected %d inputs but got %d", size, len(inputs)))
	}
}

// Definition of a class of objects
type ObjParserLeafClass[I any, O ObjParserResult] struct {
	Name    string
	Matches func(I) (bool, O)
}

type ObjParserClass[O ObjParserResult] struct {
	Name    string
	Matches func(O) bool
}

// A pattern rule to match some classes
type ObjParserRule[O ObjParserResult] struct {
	Name     string
	Patterns []*ObjParserRulePattern[O]
}

func (s *ObjParserRule[O]) getAllPatternTokenNames() []string {
	names := NewSet[string]()
	for _, pattern := range s.Patterns {
		for _, opToken := range pattern.patternTokens {
			names.Add(opToken.TargetName)
		}
	}
	return names.Values()
}

func (s *ObjParserRule[O]) tokenizePatterns() error {
	var err error
	for _, pattern := range s.Patterns {
		err = pattern.tokenizePattern()
		if err != nil {
			return err
		}
	}
	return nil
}

type ObjParserRulePattern[O ObjParserResult] struct {
	Pattern       string
	Matched       func([]O) O
	patternTokens []*objParserPatternToken
}

type objParserPatternToken struct {
	TargetName string
	ExactCount int
	// -1 == no minimum
	MinCount int
	// -1 == no maximum
	MaxCount int
}

// Read the rule text and tokenize it
var re_exact_count = regexp.MustCompile(`^(?P<name>[A-Za-z\d_]+){(?P<count>\d+)}$`)

func (s *ObjParserRulePattern[O]) tokenizePattern() error {
	s.patternTokens = make([]*objParserPatternToken, 0, 1)
	ruleTokens := strings.Split(s.Pattern, " \t\n")
	for ti, ruleToken := range ruleTokens {
		var m []string
		var opToken *objParserPatternToken

		// re_exact_count
		if opToken == nil {
			m = re_exact_count.FindStringSubmatch(ruleToken)
			if len(m) > 0 {
				name_i := re_exact_count.SubexpIndex("name")
				count_i := re_exact_count.SubexpIndex("count")
				count, err := strconv.Atoi(m[count_i])
				if err != nil {
					return fmt.Errorf("Can't convert %s to integer: %w", m[count_i], err)
				}
				opToken = &objParserPatternToken{
					TargetName: m[name_i],
					ExactCount: count,
					MinCount:   count,
					MaxCount:   count,
				}
			}
		}

		// No match?
		if opToken == nil {
			return fmt.Errorf("Did not find a matching rule pattern #%d: \"%s\"", ti+1, ruleToken)
		} else {
			s.patternTokens = append(s.patternTokens, opToken)
		}
	}
	return nil
}
