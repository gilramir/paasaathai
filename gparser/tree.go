package gparser

import (
	"fmt"
	"strings"
)

type ParserTree[O ParserResult] struct {
	name string
	// used before finalization
	childrenNames []string

	finalized bool
	/*
		success   bool
		output    O
	*/

	item TreeNode[O]
	//opType NodeType

	children []*ParserTree[O]

	// if opType == OPPatternToken, since no name maps to it
	//patternToken *ParserPatternToken
}

func (s *ParserTree[O]) CopyInto(t *ParserTree[O]) {
	t.name = s.name
	t.childrenNames = s.childrenNames
	t.finalized = s.finalized
	t.item = s.item
	t.children = s.children
}

func (s *ParserTree[O]) Dump() {
	s.dumpLevel(0)
}

func (s *ParserTree[O]) Repr() string {
	//return fmt.Sprintf("<OPTree %s %s>", NodeTypeName(s.opType), s.name)
	if s.finalized {
		if s.name == "" {
			return fmt.Sprintf("<OPTree %s>", NodeTypeName(s.item.NodeType()))
		} else {
			return fmt.Sprintf("<OPTree %s %s>", NodeTypeName(s.item.NodeType()), s.name)
		}
	} else {
		return fmt.Sprintf("<OPTree ? %s>", s.name)
	}
}

func (s *ParserTree[O]) dumpLevel(level int) {
	spaces := strings.Repeat("  ", level)
	fmt.Printf("%s%s\n", spaces, s.Repr())
	for _, child := range s.children {
		child.dumpLevel(level + 1)
	}
}

func (s *ParserTree[O]) assertFinalized() {
	if !s.finalized {
		panic("ParserTree.Attempt() has not been called yet")
	}
}
func (s *ParserTree[O]) assertNotFinalized() {
	if s.finalized {
		panic("Parser.Attempt() has already been called.")
	}
}

// WILL bE REMOVED
/*
func NewParserTree[O ParserResult](opType NodeType, name string, capacity int) *ParserTree[O] {
	n := new(ParserTree[O])
	n.Initialize(opType, name, capacity)
	return n
}
*/

func NewTreeNode[O ParserResult](name string) *ParserTree[O] {
	n := new(ParserTree[O])
	n.InitializeSetName(name)
	return n
}

func (s *ParserTree[O]) InitializeSetName(name string) {
	s.assertNotFinalized()
	s.name = name
	s.children = make([]*ParserTree[O], 0)
}

func (s *ParserTree[O]) AddChild(chName string) *ParserTree[O] {
	s.assertNotFinalized()
	ch := NewTreeNode[O](chName)
	s.children = append(s.children, ch)
	return ch
}

func NewTreeNodeWithItem[O ParserResult](item TreeNode[O], name string, capacity int) *ParserTree[O] {
	n := new(ParserTree[O])
	n.InitializeSetItem(item, name, capacity)
	return n
}

func (s *ParserTree[O]) InitializeSetItem(item TreeNode[O], name string, capacity int) {
	s.assertNotFinalized()
	s.item = item
	s.name = name
	s.children = make([]*ParserTree[O], 0, capacity)
}

/*
func (s *ParserTree[O]) InitializeSetType(opType NodeType, name string, capacity int) {
	s.assertNotFinalized()
	s.opType = opType
	s.name = name
	s.children = make([]*ParserTree[O], 0, capacity)
}
*/

func (s *ParserTree[O]) SetFinalized() {
	s.assertNotFinalized()
	s.finalized = true
}

// WILL BE REMOVED
/*
func (s *ParserTree[O]) Initialize(opType NodeType, name string, capacity int) {
	s.opType = opType
	s.name = name
	s.children = make([]*ParserTree[O], 0, capacity)
}
*/

// WILL BE REMOVED
/*
func (s *ParserTree[O]) AddNewChild(opType NodeType, chName string, capacity int) *ParserTree[O] {
	fmt.Printf("AddnewChild: %s\n", chName)
	ch := NewParserTree[O](opType, chName, capacity)
	s.children = append(s.children, ch)
	return ch
}
*/

type TreeNode[O ParserResult] interface {
	NodeType() NodeType
	//	ExpandChildren(*ParserTree[O]) []*ParserTree[O]
}

type NodeType rune

const (
	OPRule         = 'R'
	OPIOClass      = 'C'
	OPIOMap        = 'M'
	OPIOSeq        = 'S'
	OPPatternToken = 't'
	OPAnd          = '&'
	OPOr           = '|'
)

func NodeTypeName(t NodeType) string {
	switch t {
	case OPRule:
		return "Rule"
	case OPIOClass:
		return "IOClass"
	case OPIOMap:
		return "IOMap"
	case OPIOSeq:
		return "IOSeq"
	case OPAnd:
		return "AND"
	case OPOr:
		return "OR"
	case OPPatternToken:
		return "PatternToken"
	}
	return fmt.Sprintf("Error: '%v' is not a valid NodeType", t)
}
