package paasaathai

import (
	"fmt"
	"strings"
)

type ObjParserTree[O ObjParserResult] struct {
	name      string
	attempted bool
	success   bool
	output    O

	opType   OPType
	children []*ObjParserTree[O]

	// if opType == OPPatternToken, since no name maps to it
	patternToken *objParserPatternToken
}

func (s *ObjParserTree[O]) Dump() {
	s.dumpLevel(0)
}

func (s *ObjParserTree[O]) Repr() string {
	return fmt.Sprintf("<OPTree %s %s>", OPTypeName(s.opType), s.name)
}

func (s *ObjParserTree[O]) dumpLevel(level int) {
	spaces := strings.Repeat("  ", level)
	fmt.Printf("%s%s\n", spaces, s.Repr())
	for _, child := range s.children {
		child.dumpLevel(level + 1)
	}
}

func (s *ObjParserTree[O]) assertAttempted() {
	if !s.attempted {
		panic("ObjParserTree.Attempt() has not been called yet")
	}
}
func (s *ObjParserTree[O]) assertNotAttempted() {
	if s.attempted {
		panic("ObjParser.Attempt() has already been called.")
	}
}

func NewObjParserTree[O ObjParserResult](opType OPType, name string, capacity int) *ObjParserTree[O] {
	n := new(ObjParserTree[O])
	n.Initialize(opType, name, capacity)
	return n
}

func (s *ObjParserTree[O]) Initialize(opType OPType, name string, capacity int) {
	s.assertNotAttempted()
	s.opType = opType
	s.name = name
	s.children = make([]*ObjParserTree[O], 0, capacity)
}

func (s *ObjParserTree[O]) AddNewChild(opType OPType, chName string, capacity int) *ObjParserTree[O] {
	fmt.Printf("AddnewChild: %s\n", chName)
	s.assertNotAttempted()
	ch := NewObjParserTree[O](opType, chName, capacity)
	s.children = append(s.children, ch)
	return ch
}
