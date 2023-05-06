package rule

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
	Validate(input string, from, to int, fromLeft bool, rules Rules) (RuleResult, error)
}

type Rules map[string]Rule

func getSubstring(s string, from, to int) (string, error) {
	// TODO: check if invalid.

	return string([]rune(s)[from:to]), nil
}

/*
parseRules = rule.Rules{
	"main" : rule.Name("expression")
	"expression" : rule.Or(rule.Name("number"), rule.Name("plus"))
	"number" : rule.Match("5")
	"plus" : rule.Match("+")
}


[ruleName: "or"]
5
*/
