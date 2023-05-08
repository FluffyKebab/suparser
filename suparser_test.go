package suparser_test

import (
	"suparser"
	"suparser/rule"
	"testing"

	"github.com/stretchr/testify/assert"
)

var simpleCalculatorRules = rule.Rules{
	"main":       rule.And(rule.Name("expression")),
	"expression": rule.Or(rule.Name("binaryOperation"), rule.Name("number")),
	"binaryOperation": rule.AndWithCenter(
		1,
		rule.Name("expression"),
		rule.Name("binaryOperator"),
		rule.Name("expression"),
	),
	"number": rule.Match("[0-9]+", rule.MatchWithRegex(), rule.MatchWithWholeText()),
	"binaryOperator": rule.Or(
		rule.Match("*"),
		rule.Match("/"),
		rule.Match("+"),
		rule.Match("-"),
	),
}

func TestMatch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		text     string
		rules    rule.Rules
		expected suparser.Node
		err      error
	}{
		{
			rules: rule.Rules{
				"main": rule.And(rule.Name("box"), rule.Name("two")),
				"box":  rule.And(rule.Match("[0-9]+", rule.MatchWithRegex()), rule.Match("box")),
				"two":  rule.Match("2"),
			},
			text: "6380box2",
			expected: suparser.Node{
				Name: "main",
				Text: "6380box2",
				SubNodes: []suparser.Node{
					{
						Name:     "box",
						Text:     "6380box",
						SubNodes: []suparser.Node{},
					},
					{
						Name:     "two",
						Text:     "2",
						SubNodes: []suparser.Node{},
					},
				},
			},
		},
		{
			rules: simpleCalculatorRules,
			text:  "5+5",
			expected: suparser.Node{
				Name: "main",
				Text: "5+5",
				SubNodes: []suparser.Node{
					{
						Name: "expression",
						Text: "5+5",
						SubNodes: []suparser.Node{
							{
								Name: "binaryOperation",
								Text: "5+5",
								SubNodes: []suparser.Node{
									{
										Name: "expression",
										Text: "5",
										SubNodes: []suparser.Node{
											{
												Name:     "number",
												Text:     "5",
												SubNodes: []suparser.Node{},
											},
										},
									},
									{
										Name:     "binaryOperator",
										Text:     "+",
										SubNodes: []suparser.Node{},
									},
									{
										Name: "expression",
										Text: "5",
										SubNodes: []suparser.Node{
											{
												Name:     "number",
												Text:     "5",
												SubNodes: []suparser.Node{},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		node, err := suparser.Parse(tc.text, tc.rules)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.expected, node)
	}
}
