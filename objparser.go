package paasaathai

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	// graph "github.com/mcpar-land/generic-graph"
)

type OPType rune

const (
	OPRule         = 'R'
	OPLeafClass    = 'C'
	OPIOMap        = 'M'
	OPIOSeq        = 'S'
	OPPatternToken = 't'
	OPAnd          = '&'
	OPOr           = '|'
)

func OPTypeName(t OPType) string {
	switch t {
	case OPRule:
		return "rule"
	case OPLeafClass:
		return "leaf class"
	case OPIOMap:
		return "IO Map"
	case OPIOSeq:
		return "IO Seq"
	case OPAnd:
		return "AND"
	case OPOr:
		return "OR"
	case OPPatternToken:
		return "PatternToken"
	}
	return fmt.Sprintf("Error: '%v' is not a valid OPType", t)
}

// The output type O must conform to the ObjParserResult interface
type ObjParserResult interface {
	SetConsumed(v int)
}

// The holder of the definition of the parser
type ObjParser[I any, O ObjParserResult] struct {
	finalizedOk bool
	targetNames []string
	//orderedRules []*ObjParserRule[O]
	tree ObjParserTree[O]

	namespace    map[string]OPType
	ruleMap      map[string]*ObjParserRule[O]
	leafClassMap map[string]*ObjParserLeafClass[I, O]
	ioMap        map[string]I
	ioMapper     func(string, I) O
	ioSeqMap     map[string][]I
	ioSeqMapper  func(string, []I) O
}

func NewObjParser[I any, O ObjParserResult](targetNames ...string) *ObjParser[I, O] {
	s := &ObjParser[I, O]{}
	s.Initialize(targetNames...)
	return s
}

func (s *ObjParser[I, O]) Initialize(targetNames ...string) {
	s.assertNotFinalized()
	s.targetNames = targetNames
	//	s.orderedRules = make([]*ObjParserRule[O], 0)
	s.ruleMap = make(map[string]*ObjParserRule[O])
	s.leafClassMap = make(map[string]*ObjParserLeafClass[I, O])
	s.ioSeqMap = make(map[string][]I)
	s.ioSeqMap = make(map[string][]I)
	s.namespace = make(map[string]OPType)
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

func (s *ObjParser[I, O]) RegisterLeafClass(newClass *ObjParserLeafClass[I, O]) {
	s.assertNotFinalized()
	if ot, has := s.namespace[newClass.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a %s named \"%s\"", OPTypeName(ot), newClass.Name))
	}
	s.leafClassMap[newClass.Name] = newClass
	s.namespace[newClass.Name] = 'C'
}

func (s *ObjParser[I, O]) RegisterRule(newRule *ObjParserRule[O]) {
	s.assertNotFinalized()
	if ot, has := s.namespace[newRule.Name]; has {
		panic(fmt.Sprintf("ObjParser already has a %s named \"%s\"", OPTypeName(ot), newRule.Name))
	}

	err := newRule.tokenizePatterns()
	if err != nil {
		panic(fmt.Sprintf("Error parsing %s rule: %s", newRule.Name, err))
	}

	s.ruleMap[newRule.Name] = newRule
	//	s.orderedRules = append(s.orderedRules, newRule)
	s.namespace[newRule.Name] = 'R'
}

func (s *ObjParser[I, O]) RegisterIOMap(m map[string]I, f func(string, I) O) {
	s.assertNotFinalized()
	if len(s.ioMap) > 0 {
		panic("IOMap can only be registered once")
	}

	for name, _ := range m {
		if ot, has := s.namespace[name]; has {
			panic(fmt.Sprintf("ObjParser already has a %s named \"%s\"", OPTypeName(ot), name))
		}
		s.namespace[name] = 'M'
	}
	s.ioMap = m
	s.ioMapper = f
}

func (s *ObjParser[I, O]) RegisterIOSeqMap(m map[string][]I, f func(string, []I) O) {
	s.assertNotFinalized()
	if len(s.ioSeqMap) > 0 {
		panic("IOSeqMap can only be registered once")
	}

	for name, _ := range m {
		if ot, has := s.namespace[name]; has {
			panic(fmt.Sprintf("ObjParser already has a %s named \"%s\"", OPTypeName(ot), name))
		}
		s.namespace[name] = 'S'
	}
	s.ioSeqMap = m
	s.ioSeqMapper = f
}

func (s *ObjParser[I, O]) Finalize() {
	s.assertNotFinalized()

	// Sanity-check the targets
	for _, name := range s.targetNames {
		if _, has := s.namespace[name]; !has {
			panic(fmt.Sprintf("Did not find a registered item for target \"%s\"", name))
		}
	}

	// Sanity-check that the rules reference things that exists
	for rname, rule := range s.ruleMap {
		for _, pname := range rule.getAllPatternTokenNames() {
			if _, has := s.namespace[pname]; !has {
				panic(fmt.Sprintf("Rule \"%s\" references item \"%s\" but it is not defined", rname, pname))
			}
		}
	}

	// Creating the root tree node is a little special
	todo := make([]*ObjParserTree[O], 0)
	if len(s.targetNames) == 1 {
		newNode, newNodesToExamine := s.newTreeChild(s.targetNames[0])
		s.tree = *newNode
		todo = append(todo, newNodesToExamine...)
	} else {
		s.tree.Initialize(OPOr, "", len(s.targetNames))

		for _, name := range s.targetNames {
			chNode, newNodesToExamine := s.newTreeChild(name)
			s.tree.children = append(s.tree.children, chNode)
			fmt.Printf("Got target node: %s\n", chNode.Repr())
			todo = append(todo, newNodesToExamine...)
		}
	}

	// Now create the rest of the tree

	for len(todo) > 0 {
		nextNode := todo[0]
		todo = todo[1:]

		fmt.Printf("need to expand: %s\n", nextNode.Repr())
		newNodesToExamine := s.expandChildren(nextNode)
		todo = append(todo, newNodesToExamine...)
	}
	s.tree.Dump()

	s.finalizedOk = true
}

/*
type opTreeEdge struct {
	parentNode *ObjParserTree
	childName  string
}
*/
func (s *ObjParser[I, O]) expandChildren(node *ObjParserTree[O]) []*ObjParserTree[O] {
	newNodes := make([]*ObjParserTree[O], 0, len(node.children))
	/*
		for _, child := range node.children {
			switch node.opType {
			case OPRule:
				newNode, toExamine := s.newTreeChildRule(chName)
				newNodes = append(newNodes, newNode)
				newNodes = append(newNodes, toExamine...)

			case OPLeafClass:
				newNode := NewObjParserTree[O](OPLeafClass, chName, 0)
				newNodes = append(newNodes, newNode)

			case OPIOMap:
				newNode := NewObjParserTree[O](OPIOMap, chName, 0)
				newNodes = append(newNodes, newNode)

			case OPIOSeq:
				newNode := NewObjParserTree[O](OPIOSeq, chName, 0)
				newNodes = append(newNodes, newNode)

			case OPAnd:
				panic(fmt.Sprintf("Should not have seen OpIOSeqMap for node %s", chName))

			case OPOr:
				panic(fmt.Sprintf("Should not have seen OpIOSeqMap for node %s", chName))

			case OPPatternToken:
				panic(fmt.Sprintf("Should not have seen OpPatternToken for node %s", chName))
			}

		}
	*/
	return newNodes

}

func (s *ObjParser[I, O]) newTreeChild(chName string) (*ObjParserTree[O], []*ObjParserTree[O]) {
	var newNode *ObjParserTree[O]
	toExamine := make([]*ObjParserTree[O], 0)
	chOpType := s.namespace[chName]

	switch chOpType {
	case OPRule:
		newNode, toExamine = s.newTreeChildRule(chName)

	case OPLeafClass:
		newNode = NewObjParserTree[O](OPLeafClass, chName, 0)

	case OPIOMap:
		newNode = NewObjParserTree[O](OPIOMap, chName, 0)

	case OPIOSeq:
		newNode = NewObjParserTree[O](OPIOSeq, chName, 0)

	case OPAnd:
		panic(fmt.Sprintf("Should not have seen OpIOSeqMap for node %s", chName))

	case OPOr:
		panic(fmt.Sprintf("Should not have seen OpIOSeqMap for node %s", chName))

	case OPPatternToken:
		panic(fmt.Sprintf("Should not have seen OpPatternToken for node %s", chName))
	}

	return newNode, toExamine
}

func (s *ObjParser[I, O]) newTreeChildRule(chName string) (*ObjParserTree[O], []*ObjParserTree[O]) {

	r := s.ruleMap[chName]

	if len(r.Patterns) == 1 {
		// If there's only one pattern, we don't need a rule (OR) node
		newNode, toExamine := s.newTreeChildRulePattern(chName, 0)
		return newNode, toExamine
	}

	// Create the rule node
	rNode := NewObjParserTree[O](OPOr, chName, len(r.Patterns))
	toExamine := make([]*ObjParserTree[O], 0)

	for pi, _ := range r.Patterns {
		pNode, pToExamine := s.newTreeChildRulePattern(chName, pi)
		rNode.children = append(rNode.children, pNode)
		toExamine = append(toExamine, pToExamine...)
	}
	return rNode, toExamine
}

func (s *ObjParser[I, O]) newTreeChildRulePattern(chName string, patternN int) (*ObjParserTree[O], []*ObjParserTree[O]) {

	r := s.ruleMap[chName]
	p := r.Patterns[patternN]

	if len(p.patternTokens) == 1 {
		// If there's only one token, we don't need a pattern (AND) node
		newNode, toExamine := s.newTreeChildPatternToken(p.patternTokens[0])
		return newNode, toExamine
	}

	// Create the pattern node
	patternName := fmt.Sprintf("%s pattern #%d", chName, patternN+1)
	pNode := NewObjParserTree[O](OPAnd, patternName, len(p.patternTokens))
	toExamine := make([]*ObjParserTree[O], 0)

	for _, t := range p.patternTokens {
		tNode, tToExamine := s.newTreeChildPatternToken(t)
		pNode.children = append(pNode.children, tNode)
		toExamine = append(toExamine, tToExamine...)
	}
	return pNode, toExamine

}

func (s *ObjParser[I, O]) newTreeChildPatternToken(token *objParserPatternToken) (*ObjParserTree[O], []*ObjParserTree[O]) {
	newNode := NewObjParserTree[O](OPPatternToken, token.ShownAs, 0)
	newNode.patternToken = token

	chNode, toExamine := s.newTreeChild(token.TargetName)
	newNode.children = append(newNode.children, chNode)

	return newNode, toExamine
}

// -----------------------------------------------------------------------------------

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
	/*
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
	*/

	return s.results, nil
}

/*
type objParserTrialNode[I any, O ObjParserResult] struct {
	rule         *ObjParserRule[O]
	rulePattern  *ObjParserRulePattern[O]
	rulePatternN int
	class        *ObjParserClass[O]
	leafClass    *ObjParserLeafClass[I, O]

	children []*objParserTrialNode[I, O]
}
*/

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
	ShownAs    string
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
	ruleTokens := strings.Split(s.Pattern, " ")
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
					ShownAs:    ruleToken,
					TargetName: m[name_i],
					ExactCount: count,
					MinCount:   count,
					MaxCount:   count,
				}
			}
		}

		// ioMap

		// No match?
		if opToken == nil {
			return fmt.Errorf("Did not find a matching rule pattern #%d: \"%s\"", ti+1, ruleToken)
		} else {
			s.patternTokens = append(s.patternTokens, opToken)
		}
	}
	return nil
}
