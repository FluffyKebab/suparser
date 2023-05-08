package rule

type or struct {
	subRules []Rule
}

var _ Rule = or{}

func Or(subRules ...Rule) or {
	return or{
		subRules: subRules,
	}
}

func (r or) Validate(input string, from, to int, fromRight bool, rules Rules) (RuleResult, error) {
	for _, rule := range r.subRules {
		result, err := rule.Validate(input, from, to, fromRight, rules)
		if err != nil || result.RuleType != "" {
			return result, err
		}
	}

	return RuleResult{}, nil
}
