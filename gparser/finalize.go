package gparser

import "fmt"

func (s *Parser[I, O]) Finalize() {
	s.assertNotFinalized()

	// Sanity-check the targets
	for _, name := range s.targetNames {
		if _, has := s.namespace[name]; !has {
			panic(fmt.Sprintf("Did not find a registered item for target \"%s\"", name))
		}
	}

	// Sanity-check that the rules reference things that exist
	for rname, rule := range s.ruleMap {
		for _, pname := range rule.getAllPatternTokenNames() {
			if _, has := s.namespace[pname]; !has {
				panic(fmt.Sprintf("Rule \"%s\" references item \"%s\" but it is not defined", rname, pname))
			}
		}
	}

	// The list of ParserTree nodes that have names but have not
	// been finalized.
	todo := make([]*ParserTree[O], 0)

	// Creating the root tree node is a little special
	//	fmt.Printf("Target names: %v\n", s.targetNames)
	if len(s.targetNames) == 1 {
		s.tree.name = s.targetNames[0]
		todo = append(todo, &s.tree)
	} else {
		//s.tree.InitializeSetType(OPOr, "", len(s.targetNames))
		s.tree.InitializeSetItem(NewItemOr(), "", len(s.targetNames))

		for _, name := range s.targetNames {
			chNode := s.tree.AddChild(name)
			todo = append(todo, chNode)
		}
		s.tree.SetFinalized()
	}

	// Now create the rest of the tree
	for len(todo) > 0 {
		thisNode := todo[0]
		todo = todo[1:]

		//		fmt.Printf("need to expand: %s\n", thisNode.Repr())

		newNodesToExamine := s.expandNamedNode(thisNode)
		todo = append(todo, newNodesToExamine...)
	}
	s.tree.Dump()

	// Now that the parser definition is parsed into a tree, convert
	// into a non-deterministic finite automata (NFA).
	s.tree2nfa()

	s.finalizedOk = true
}

/*
type opTreeEdge struct {
	parentNode *ParserTree
	childName  string
}
*/

// The node has a name, but no item yet
// Returns any new nodes with names that we need to iterate on
func (s *Parser[I, O]) expandNamedNode(node *ParserTree[O]) []*ParserTree[O] {

	defer node.SetFinalized()

	nodeType, has := s.namespace[node.name]
	if !has {
		panic(fmt.Sprintf("Did not find node named '%s'\n", node.name))
	}

	//	fmt.Printf("expanding named node %s\n", node.name)
	newNodes := make([]*ParserTree[O], 0, len(node.childrenNames))
	switch nodeType {
	case OPRule:
		rule := s.ruleMap[node.name]
		ruleNewNodes := s.expandRuleInPlace(rule, node)
		newNodes = append(newNodes, ruleNewNodes...)

	case OPIOClass:
		ioClass := s.ioClassMap[node.name]
		node.item = ioClass
	}

	if len(node.childrenNames) == 0 {
		return newNodes
	}

	for _, chName := range node.childrenNames {
		//chNode := s.tree.AddChild(chName)
		chNode := node.AddChild(chName)
		newNodes = append(newNodes, chNode)
	}

	return newNodes

}

func (s *Parser[I, O]) expandRuleInPlace(r *ParserRule[O], node *ParserTree[O]) []*ParserTree[O] {

	if len(r.Patterns) == 1 {
		// If there's only one pattern, we don't need a rule (OR) node
		// ......
		pNode, newNodes := s.expandRulePattern(r, 0)
		pNode.CopyInto(node)
		return newNodes
	}

	// Create the rule node
	orNode := NewTreeNodeWithItem[O](NewItemOr(), r.Name, len(r.Patterns))
	newNodes := make([]*ParserTree[O], 0)

	for pi, _ := range r.Patterns {
		pNode, pNewNodes := s.expandRulePattern(r, pi)
		newNodes = append(newNodes, pNewNodes...)
		orNode.children = append(orNode.children, pNode)
	}
	orNode.CopyInto(node)
	return newNodes
}

func (s *Parser[I, O]) expandRulePattern(r *ParserRule[O], patternN int) (*ParserTree[O], []*ParserTree[O]) {

	p := r.Patterns[patternN]

	if len(p.Tokens) == 1 {
		// If there's only one token, we don't need a pattern (AND) node
		tNode, newNodes := s.expandRulePatternToken(r, p, 0)
		return tNode, newNodes
	}

	// Create the pattern node
	patternName := fmt.Sprintf("%s pattern %d", r.Name, patternN+1)
	andNode := NewTreeNodeWithItem[O](NewItemAnd(), patternName, len(p.Tokens))
	newNodes := make([]*ParserTree[O], 0)

	for ti, _ := range p.Tokens {
		tNode, _ := s.expandRulePatternToken(r, p, ti)
		//newNodes = append(newNodes, tNewNodes...)
		newNodes = append(newNodes, tNode)
		andNode.children = append(andNode.children, tNode)
	}
	//fmt.Printf("andNode children=%v\n", andNode.children)
	return andNode, newNodes
}

// TODO - remove 2nd returned argument
func (s *Parser[I, O]) expandRulePatternToken(r *ParserRule[O], p *ParserRulePattern[O], tokenN int) (*ParserTree[O], []*ParserTree[O]) {
	t := p.Tokens[tokenN]

	newNode := NewTreeNodeWithItem[O](t, t.ShownAs, 0)

	//	newNode.childrenNames = append(newNode.childrenNames, t.TargetName)

	newNodes := make([]*ParserTree[O], 0)

	return newNode, newNodes
}
