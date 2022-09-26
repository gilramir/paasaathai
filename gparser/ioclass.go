package gparser

// Definition of a class of objects
type IOClass[I any, O ParserResult] struct {
	Name    string
	Matches func(I) (bool, O)
}

func (s *IOClass[I, O]) NodeType() NodeType {
	return OPIOClass
}

type LogicalItemOr struct{}

func (s *LogicalItemOr) NodeType() NodeType {
	return OPOr
}

func NewItemOr() *LogicalItemOr {
	return &LogicalItemOr{}
}

type LogicalItemAnd struct{}

func (s *LogicalItemAnd) NodeType() NodeType {
	return OPAnd
}

func NewItemAnd() *LogicalItemAnd {
	return &LogicalItemAnd{}
}
