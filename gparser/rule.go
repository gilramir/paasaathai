package gparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// A pattern rule to match some classes
type ParserRule[O ParserResult] struct {
	Name     string
	Patterns []*ParserRulePattern[O]
}

func (s *ParserRule[O]) NodeType() NodeType {
	return OPRule
}

func (s *ParserRule[O]) getAllPatternTokenNames() []string {
	names := NewSet[string]()
	for _, pattern := range s.Patterns {
		for _, opToken := range pattern.Tokens {
			names.Add(opToken.TargetName)
		}
	}
	return names.Values()
}

func (s *ParserRule[O]) tokenizePatterns() error {
	var err error
	for _, pattern := range s.Patterns {
		err = pattern.tokenizePattern()
		if err != nil {
			return err
		}
	}
	return nil
}

type ParserRulePattern[O ParserResult] struct {
	Pattern string
	Matched func([]O) O
	Tokens  []*ParserPatternToken
}

type ParserPatternToken struct {
	ShownAs    string
	TargetName string
	ExactCount int
	// -1 == no minimum
	MinCount int
	// -1 == no maximum
	MaxCount int
}

func (s *ParserPatternToken) NodeType() NodeType {
	return OPPatternToken
}

// Read the rule text and tokenize it
var re_exact_count = regexp.MustCompile(`^(?P<name>[A-Za-z\d_]+){(?P<count>\d+)}$`)

// Just the name of another item
var re_simple_name = regexp.MustCompile(`^(?P<name>[A-Za-z\d_]+)$`)

func (s *ParserRulePattern[O]) tokenizePattern() error {
	s.Tokens = make([]*ParserPatternToken, 0, 1)
	ruleTokens := strings.Split(s.Pattern, " ")
	for _, ruleToken := range ruleTokens {
		var m []string
		var opToken *ParserPatternToken

		// re_simple_name
		if opToken == nil {
			m = re_simple_name.FindStringSubmatch(ruleToken)
			if len(m) > 0 {
				name_i := re_exact_count.SubexpIndex("name")
				opToken = &ParserPatternToken{
					ShownAs:    ruleToken,
					TargetName: m[name_i],
					ExactCount: 1,
					MinCount:   1,
					MaxCount:   1,
				}
			}
		}

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
				opToken = &ParserPatternToken{
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
			return fmt.Errorf("Did not find a matching item name: \"%s\"", ruleToken)
		} else {
			s.Tokens = append(s.Tokens, opToken)
		}
	}
	return nil
}
