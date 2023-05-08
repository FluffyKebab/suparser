package rule

import "fmt"

type RuleType string

var (
	NameType  RuleType = "name"
	MatchType RuleType = "match"
	AndType   RuleType = "and"
)

type RuleResult struct {
	RuleName        string
	RuleType        RuleType
	From            int
	To              int
	SubRulesMatched []RuleResult
}

type Rule interface {
	Validate(input string, from, to int, fromRight bool, rules Rules) (RuleResult, error)
	GetError(rules Rules) error
}

type Rules map[string]Rule

func getSubstring(s string, from, to int) (string, error) {
	if from < 0 || to > len(s) {
		return "", fmt.Errorf(
			"%w: from: %v to: %v length s: %v",
			ErrInvalidFromTo,
			from,
			to,
			len(s),
		)
	}

	return string([]rune(s)[from:to]), nil
}
