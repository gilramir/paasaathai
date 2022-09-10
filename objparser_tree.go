package paasaathai

import "fmt"

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
	s.children = make([]*ObjParserTree[O], 0, capacity)
}

func (s *ObjParserTree[O]) AddNewChild(opType OPType, chName string, capacity int) *ObjParserTree[O] {
	fmt.Printf("AddnewChild: %s\n", chName)
	s.assertNotAttempted()
	ch := NewObjParserTree[O](opType, chName, capacity)
	s.children = append(s.children, ch)
	return ch
}
