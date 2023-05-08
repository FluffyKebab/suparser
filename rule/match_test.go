package rule_test

import (
	"suparser/rule"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input     string
		rules     rule.Rules
		from      int
		to        int
		expected  rule.RuleResult
		fromRight bool
		err       error
	}{
		{
			input: "+-",
			rules: rule.Rules{"main": rule.Match(`\-`, rule.MatchWithRegex())},
			from:  0,
			to:    2,
			expected: rule.RuleResult{
				RuleType: rule.MatchType,
				From:     1,
				To:       2,
			},
		},
		{
			input: "a-",
			rules: rule.Rules{"main": rule.Match(`a+\-`, rule.MatchWithRegex())},
			from:  0,
			to:    2,
			expected: rule.RuleResult{
				RuleType: rule.MatchType,
				From:     0,
				To:       2,
			},
		},
		{
			input: "+-",
			rules: rule.Rules{"main": rule.Match(`+-`)},
			from:  0,
			to:    2,
			expected: rule.RuleResult{
				RuleType: rule.MatchType,
				From:     0,
				To:       2,
			},
		},
		{
			input: "foo ba bar fo + za",
			rules: rule.Rules{"main": rule.Match(`+`)},
			from:  4,
			to:    18,
			expected: rule.RuleResult{
				RuleType: rule.MatchType,
				From:     14,
				To:       15,
			},
		},
		{
			input: "fo+ ba bar fo + za",
			rules: rule.Rules{"main": rule.Match(`+`)},
			from:  0,
			to:    18,
			expected: rule.RuleResult{
				RuleType: rule.MatchType,
				From:     14,
				To:       15,
			},
			fromRight: true,
		},
		{
			input: "fo5 ba b3r fo 12 za",
			rules: rule.Rules{"main": rule.Match(`[0-9]+`, rule.MatchWithRegex())},
			from:  0,
			to:    18,
			expected: rule.RuleResult{
				RuleType: rule.MatchType,
				From:     14,
				To:       16,
			},
			fromRight: true,
		},
	}

	for _, tc := range testCases {
		output, err := tc.rules["main"].Validate(tc.input, tc.from, tc.to, tc.fromRight, tc.rules)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.expected, output)
	}
}
