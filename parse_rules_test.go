package suparser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRulesToRuleMap(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input       string
		want        map[string]string
		expectedErr error
	}{
		{
			input:       "rule1: bla bla",
			want:        map[string]string{"rule1": "bla bla"},
			expectedErr: nil,
		},
		{
			input:       "rule1: bla bla\n \t \n avi: foo ba",
			want:        map[string]string{"rule1": "bla bla", "avi": "foo ba"},
			expectedErr: nil,
		},
		{
			input:       "rule1 bla bla",
			want:        nil,
			expectedErr: InvalidRulesError{Line: 1, Reason: `no ":"`},
		},
		{
			input:       "rule1:bla:bla",
			want:        nil,
			expectedErr: InvalidRulesError{Line: 1, Reason: `multiple ":"`},
		},
		{
			input:       "rule1:yo\n rule1:no",
			want:        nil,
			expectedErr: InvalidRulesError{Line: 2, Reason: `rule "rule1" has multiple declarations`},
		},
		{
			input:       ":yo",
			want:        nil,
			expectedErr: InvalidRulesError{Line: 1, Reason: `missing rule name`},
		},
		{
			input:       " sd : ",
			want:        nil,
			expectedErr: InvalidRulesError{Line: 1, Reason: `missing rule content`},
		},
	}

	for _, tc := range cases {
		result, err := rulesToRuleMap(tc.input)
		require.Equal(t, tc.expectedErr, err)
		require.Equal(t, tc.want, result)
	}
}
