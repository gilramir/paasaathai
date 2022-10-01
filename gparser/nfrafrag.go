package gparser

import (
	"fmt"
	"strings"
)

// https://medium.com/@phanindramoganti/regex-under-the-hood-implementing-a-simple-regex-compiler-in-go-ef2af5c6079
// https://github.com/phanix5/Simple-Regex-Complier-in-Go/blob/master/regex.go
// https://www.oilshell.org/archive/Thompson-1968.pdf

/*
 * Represents an NFA state plus zero or one or two arrows exiting.
 * if c == 0, it's an IOClass, and only out is set
 * if c == Match, no arrows out; matching state.
 * If c == Split, unlabeled arrows to out and out1 (if != NULL).
 */
const (
	Match = iota + 1
	Split
)

type State[I any, O ParserResult] struct {
	c         int
	io        *IOClass[I, O]
	out, out1 *State[I, O]
	lastlist  int
}

func (s *State[I, O]) Repr0() string {
	var label string
	if s.c == Match {
		label = "MATCH"
	} else if s.c == Split {
		label = "SPLIT"
	} else {
		label = s.io.Name
	}
	return fmt.Sprintf("<State %s lastlist=%d>", label, s.lastlist)
}

func (s *State[I, O]) Repr() string {
	return s.ReprN(0)
}

func (s *State[I, O]) ReprN(n int) string {
	indent := strings.Repeat("  ", n)
	txt := fmt.Sprintf("%s%s", indent, s.Repr0())
	if s.out != nil {
		txt += "\n" + s.out.ReprN(n+1)
	}
	if s.out1 != nil {
		txt += "\n" + s.out1.ReprN(n+1)
	}
	return txt
}

func (s *State[I, O]) RecursiveClearState() {
	s.lastlist = 0
	if s.out != nil {
		s.out.RecursiveClearState()
	}
	if s.out1 != nil {
		s.out1.RecursiveClearState()
	}
}

type Frag[I any, O ParserResult] struct {
	start *State[I, O]
	out   []**State[I, O]
}

func (s *Frag[I, O]) Repr() string {
	out_repr := make([]string, len(s.out))
	for i, o := range s.out {
		out_repr[i] = (*o).Repr()
	}
	return fmt.Sprintf("start: %s out: %v", s.start.Repr(), out_repr)
}

func (s *Parser[I, O]) tree2nfa() {
	// stp is where a new item will be placed in the stack.
	// nfastack must always have allocated space for an item at index 'stp'
	s.stp = 0
	s.stack = make([]Frag[I, O], 1)
	s.node2nfa(&s.tree)

	// After pushing and popping the stack, it should be empty
	s.stp--
	e := s.stack[s.stp]
	if s.stp != 0 {
		panic(fmt.Sprintf("tree2nfa failed: stp=%d e=%s", s.stp,
			e.Repr()))
	}
	s.patch(e.out, &s.matchstate)
	s.nfa = e.start

	// Dump it.
	fmt.Printf("nfa:\n%s\n", s.nfa.Repr())
}

/* Patch the list of states at out to point to s. */
func (s *Parser[I, O]) patch(out []**State[I, O], ns *State[I, O]) {
	for _, p := range out {
		*p = ns
	}
}

func (s *Parser[I, O]) ensure_stack_space() {
	if len(s.stack) <= s.stp {
		extra := s.stp - len(s.stack) + 1
		s.stack = append(s.stack, make([]Frag[I, O], extra)...)
	}
}

func (s *Parser[I, O]) node2nfa(node *ParserTree[O]) {

	switch node.item.NodeType() {
	default:
		e := fmt.Sprintf("node2nfa: %s not yet handled\n", NodeTypeName(node.item.NodeType()))
		panic(e)

	case OPIOClass:
		ns := State[I, O]{io: node.item.(*IOClass[I, O]), out: nil, out1: nil}
		s.stack[s.stp] = Frag[I, O]{&ns, []**State[I, O]{&ns.out}}
		s.stp++
		s.ensure_stack_space()

	case OPOr: /*alternate*/
		// Create entries in the stack for each child, and along the
		// way, pair them with SPLIT nodes. Each SPLIT node can only
		// have 2 outputs
		if len(node.children) < 2 {
			panic(fmt.Sprintf("%s has only %d children", node.Repr(), len(node.children)))
		}
		for i, ch := range node.children {
			s.node2nfa(ch)
			if i >= 1 {

				s.stp--
				e2 := s.stack[s.stp]
				s.stp--
				e1 := s.stack[s.stp]
				ns := State[I, O]{c: Split, out: e1.start, out1: e2.start}
				s.stack[s.stp] = Frag[I, O]{&ns, append(e1.out, e2.out...)}
				s.stp++
				// No need to call ensure_stack_space here; we popped 2
				// and added 1
			}
		}
	}
}

func (s *Parser[I, O]) Parse(input []I) ([]O, error) {
	fmt.Printf("Parsing %v\n", input)
	o := make([]O, 0)
	// Reset the state after any previous parse
	s.nfa.RecursiveClearState()
	s.listid = 0

	m := s.match(s.nfa, input)
	fmt.Printf("Matches: %v\n", m)
	return o, nil
}

func (s *Parser[I, O]) stateListRepr(stateList []*State[I, O]) string {
	labels := make([]string, len(stateList))
	for i, ns := range stateList {
		labels[i] = ns.Repr0()
	}
	return fmt.Sprintf("[%s]", strings.Join(labels, ", "))
}

func (s *Parser[I, O]) match(start *State[I, O], input []I) bool {

	var clist, nlist []*State[I, O]
	s.listid++
	clist = s.addstate(clist, start)

	for i, ch := range input {
		fmt.Printf("Input #%d: %v: clist=%s nlist=%s\n",
			i, ch, s.stateListRepr(clist),
			s.stateListRepr(nlist))
		nlist = s.step(clist, ch, nlist)
		clist, nlist = nlist, clist
	}
	return s.ismatch(clist)

}

/* Add s to l, following unlabeled arrows. */
func (s *Parser[I, O]) addstate(l []*State[I, O], ns *State[I, O]) []*State[I, O] {
	if ns == nil || ns.lastlist == s.listid {
		return l
	}
	ns.lastlist = s.listid
	if ns.c == Split {
		l = s.addstate(l, ns.out)
		l = s.addstate(l, ns.out1)
	}
	l = append(l, ns)
	return l
}

/*
 * Step the NFA from the states in clist
 * past the character ch,
 * to create next NFA state set nlist.
 */
func (s *Parser[I, O]) step(clist []*State[I, O], ch I, nlist []*State[I, O]) []*State[I, O] {
	s.listid++
	nlist = nlist[:0]
	fmt.Printf("step: nlist=%s\n", s.stateListRepr(nlist))
	for i, ns := range clist {
		if ns.io == nil {
			continue
		}
		m, _ := ns.io.Matches(ch)
		fmt.Printf("step: clist %d %s => %v\n", i, ns.Repr0(), m)
		// TODO - how to record the output?
		if m {
			nlist = s.addstate(nlist, ns.out)
		}
	}
	return nlist
}

/* Check whether state list contains a match. */
func (s *Parser[I, O]) ismatch(l []*State[I, O]) bool {
	for _, ns := range l {
		if ns == &s.matchstate {
			return true
		}
	}
	return false
}
