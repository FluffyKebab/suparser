package rule_test

import (
	"suparser/rule"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input    string
		rules    rule.Rules
		from     int
		to       int
		expected rule.RuleResult
		fromLeft bool
		err      error
	}{
		{
			input: "++++foo bar--q",
			rules: rule.Rules{"main": rule.And(rule.Match("+++"), rule.Match("--"), rule.Match("q"))},
			from:  1,
			to:    14,
			expected: rule.RuleResult{
				RuleType: rule.AndType,
				From:     1,
				To:       14,
				SubRulesMatched: []rule.RuleResult{
					{
						RuleType: rule.MatchType,
						From:     1,
						To:       4,
					},
					{
						RuleType: rule.MatchType,
						From:     11,
						To:       13,
					},
					{
						RuleType: rule.MatchType,
						From:     13,
						To:       14,
					},
				},
			},
		},
		{
			input:    "++++foo+bar--q",
			rules:    rule.Rules{"main": rule.And(rule.Match("+"), rule.Match("bar"))},
			fromLeft: true,
			from:     1,
			to:       14,
			expected: rule.RuleResult{
				RuleType: rule.AndType,
				From:     1,
				To:       14,
				SubRulesMatched: []rule.RuleResult{
					{
						RuleType: rule.MatchType,
						From:     7,
						To:       8,
					},
					{
						RuleType: rule.MatchType,
						From:     8,
						To:       11,
					},
				},
			},
		},
		{
			input:    "++++foo+bar--q",
			rules:    rule.Rules{"main": rule.AndWithCenter(1, rule.Match("foo"), rule.Match("+"), rule.Match("bar"))},
			fromLeft: true,
			from:     1,
			to:       14,
			expected: rule.RuleResult{
				RuleType: rule.AndType,
				From:     1,
				To:       14,
				SubRulesMatched: []rule.RuleResult{
					{
						RuleType: rule.MatchType,
						From:     4,
						To:       7,
					},
					{
						RuleType: rule.MatchType,
						From:     7,
						To:       8,
					},
					{
						RuleType: rule.MatchType,
						From:     8,
						To:       11,
					},
				},
			},
		},
		{
			input:    "++++foo+bar--q",
			rules:    rule.Rules{"main": rule.AndWithCenter(1, rule.Match("foo"), rule.Match("+"), rule.Match("bar"))},
			fromLeft: false,
			from:     1,
			to:       14,
			expected: rule.RuleResult{},
		},
	}

	for _, tc := range testCases {
		output, err := tc.rules["main"].Validate(tc.input, tc.from, tc.to, tc.fromLeft, tc.rules)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.expected, output)
	}
}
