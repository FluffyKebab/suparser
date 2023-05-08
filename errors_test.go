package suparser_test

import (
	"suparser"
	"suparser/rule"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		rules       rule.Rules
		expectedErr error
	}{
		{
			rules: rule.Rules{
				"foo": rule.Match("foo"),
			},
			expectedErr: suparser.ErrMissingMainRule,
		},
		{
			rules: rule.Rules{
				"main": rule.Name("foo"),
			},
			expectedErr: rule.ErrRuleNotDefined,
		},
	}

	for _, tc := range testCases {
		_, err := suparser.New(tc.rules)
		assert.ErrorIs(t, err, tc.expectedErr)
	}
}
