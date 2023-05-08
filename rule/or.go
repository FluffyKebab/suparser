package rule

type or struct {
	subRules []Rule
}

var _ Rule = or{}

// Or creates a rule that validates if one of the sub rules validates. The
// or rule searches from left to right, so the leftmost rule that validates
// will be the result used.
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

func (r or) GetError(rules Rules) error {
	for _, rule := range r.subRules {
		if err := rule.GetError(rules); err != nil {
			return err
		}
	}

	return nil
}
