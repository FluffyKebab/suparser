package rule

type name struct {
	ruleName string
}

var _ Rule = name{}

func Name(ruleName string) name {
	return name{
		ruleName: ruleName,
	}
}

func (r name) Validate(input string, from, to int, fromRight bool, rules Rules) (RuleResult, error) {
	ruleToMatch := rules[r.ruleName]
	result, err := ruleToMatch.Validate(input, from, to, fromRight, rules)
	if err != nil {
		return RuleResult{}, err
	}

	if result.RuleType == "" {
		return RuleResult{}, nil
	}

	return RuleResult{
		RuleName:        r.ruleName,
		RuleType:        NameType,
		From:            result.From,
		To:              result.To,
		SubRulesMatched: []RuleResult{result},
	}, nil
}

func (r name) GetError(rules Rules) error {
	if _, ok := rules[r.ruleName]; !ok {
		return ErrRuleNotDefined
	}

	return nil
}
