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

func (s *State[I, O]) Repr() string {
	return s.ReprN(0)
}

func (s *State[I, O]) ReprN(n int) string {
	indent := strings.Repeat("  ", n)
	if s.c == Match {
		return fmt.Sprintf("%s<State MATCH lastlist=%d>", indent, s.lastlist)
	} else if s.c == Split {
		txt := fmt.Sprintf("%s<State SPLIT lastlist=%d>", indent, s.lastlist)
		if s.out != nil {
			txt += "\n" + s.out.ReprN(n+1)
		}
		if s.out1 != nil {
			txt += "\n" + s.out1.ReprN(n+1)
		}
		return txt
	} else {
		txt := fmt.Sprintf("%s<State %s lastlist=%d>", indent, s.io.Name, s.lastlist)
		if s.out != nil {
			txt += "\n" + s.out.ReprN(n+1)
		}
		if s.out1 != nil {
			txt += "\n" + s.out1.ReprN(n+1)
		}
		return txt
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
		fmt.Printf("node2nfa: %s not yet handled\n", NodeTypeName(node.item.NodeType()))

	case OPIOClass:
		ns := State[I, O]{io: node.item.(*IOClass[I, O]), out: nil, out1: nil}
		s.stack[s.stp] = Frag[I, O]{&ns, []**State[I, O]{&ns.out}}
		s.stp++
		s.ensure_stack_space()
	}

}

func (s *Parser[I, O]) Parse(input []I) ([]O, error) {
	fmt.Printf("Parsing %v\n", input)
	o := make([]O, 0)
	m := s.match(s.nfa, input)
	fmt.Printf("Matches: %v\n", m)
	return o, nil
}

func (s *Parser[I, O]) match(start *State[I, O], input []I) bool {

	var clist, nlist []*State[I, O]
	s.listid++
	clist = s.addstate(clist, start)

	for _, ch := range input {
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
	for _, ns := range clist {
		m, _ := ns.io.Matches(ch)
		// TODO - how to record this?
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
