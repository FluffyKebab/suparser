package rule

import "fmt"

type RuleType string

var (
	NameType  RuleType = "name"
	MatchType RuleType = "match"
	AndType   RuleType = "and"
)

// RuleResult is the result returned by a rule.
type RuleResult struct {
	// RuleName is only used by the Name rule and contains the name of the
	// rule validated.
	RuleName string
	// RuleType is the type of rule that the rule is.
	RuleType        RuleType
	From            int
	To              int
	SubRulesMatched []RuleResult
}

// Rule is the standard interface for rules.
type Rule interface {
	// Validates returns a "RuleResult" with values, if the input text follows
	// the rule. If "FromRight" is true the rule result will be the rightmost
	// part of the text that follows the rule, otherwise it is the leftmost.
	Validate(input string, from, to int, fromRight bool, rules Rules) (RuleResult, error)
	// GetError returns an error if the rule is invalid.
	GetError(rules Rules) error
}

// Rules is the parser rules given to create a new parser. The key is the name
// of the rule.
//
// The rules must contain a rule with the name "main". This rule will be used as
// the rule the text must validate.
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
