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

func (r name) Validate(input string, from, to int, fromLeft bool, rules Rules) (RuleResult, error) {
	ruleToMatch, ok := rules[r.ruleName]
	if !ok {
		return RuleResult{}, ErrRuleNotDefined
	}

	result, err := ruleToMatch.Validate(input, from, to, fromLeft, rules)
	if err != nil {
		return RuleResult{}, err
	}

	if result.RuleName == "" {
		return RuleResult{}, nil
	}

	return RuleResult{
		RuleName:        r.ruleName,
		RuleType:        NameType,
		From:            from,
		To:              to,
		SubRulesMatched: []RuleResult{result},
	}, nil
}
