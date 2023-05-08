package suparser

import (
	"errors"
	"suparser/rule"
)

// ErrNotParsable is returned if the main rule is not validated by the text
// given.
var ErrNotParsable = errors.New("text is not parsable with rules given")

// Node is the struct text is parsed into.
type Node struct {
	// Name stores the name of the rule the node represent.
	Name string
	// Text stores the raw text of what the rule.
	Text string
	// SubNodes stores all sub Nodes.
	SubNodes []Node
}

// Parser is the struct used to parse text using parser rules.
type Parser struct {
	rules rule.Rules
}

// New returns a new parser with the rules given.
func New(rules rule.Rules) (Parser, error) {
	return Parser{
		rules: rules,
	}, getErrors(rules)
}

func (p Parser) Parse(text string) (Node, error) {
	ruleResult, err := p.rules["main"].Validate(
		text,
		0,         // Parse start.
		len(text), // Parse end.
		false,     // Parse from left to right.
		p.rules,
	)
	if err != nil {
		return Node{}, err
	}
	if ruleResult.RuleType == "" {
		return Node{}, ErrNotParsable
	}

	mainRuleText, err := getSubstring(text, ruleResult.From, ruleResult.To)
	if err != nil {
		return Node{}, err
	}

	mainRuleSubNodes, err := parseRuleResult(text, ruleResult)

	return Node{
		Name:     "main",
		Text:     mainRuleText,
		SubNodes: mainRuleSubNodes,
	}, nil
}

// Parse parses the text using the rules given.
func Parse(text string, rules rule.Rules) (Node, error) {
	p, err := New(rules)
	if err != nil {
		return Node{}, err
	}

	return p.Parse(text)
}
