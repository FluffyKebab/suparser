package rule

import (
	"fmt"
	"regexp"
)

// MatchOption is a function type for modifying a match rule.
type MatchOption func(*match)

// MatchWithRegex returns a MatchOption that specifies that regex can be used.
func MatchWithRegex() MatchOption {
	return func(m *match) {
		m.useRegex = true
	}
}

func MatchWithWholeText() MatchOption {
	return func(m *match) {
		m.withWholeText = true
	}
}

type match struct {
	regex         string
	useRegex      bool
	withWholeText bool
}

var _ Rule = match{}

// Match creates an rule that validates if the toMatch string is found in the text.
func Match(toMatch string, opts ...MatchOption) match {
	r := match{
		regex: toMatch,
	}

	for _, opt := range opts {
		opt(&r)
	}

	if !r.useRegex {
		r.regex = regexp.QuoteMeta(r.regex)
	}

	if r.withWholeText {
		r.regex = "^" + r.regex + "$"
	}

	return r
}

func (r match) Validate(input string, from, to int, fromRight bool, rules Rules) (RuleResult, error) {
	inputToMatch, err := getSubstring(input, from, to)
	if err != nil {
		return RuleResult{}, err
	}

	re, err := regexp.Compile(r.regex)
	if err != nil {
		return RuleResult{}, fmt.Errorf("%w: %w", ErrInvalidRegex, err)
	}

	allMatchedIndexes := re.FindAllIndex([]byte(inputToMatch), -1)
	if len(allMatchedIndexes) == 0 {
		return RuleResult{}, nil
	}

	correctMatchedIndexes := allMatchedIndexes[0]
	if fromRight {
		correctMatchedIndexes = allMatchedIndexes[len(allMatchedIndexes)-1]
	}

	return RuleResult{
		RuleType: MatchType,
		From:     from + correctMatchedIndexes[0],
		To:       from + correctMatchedIndexes[1],
	}, nil
}
