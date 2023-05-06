package rule

type and struct {
	rules  []Rule
	center int
}

var _ Rule = and{}

// And returns a rule that validates if all sub rules validates in succession from
// left to right.
//
// For example And(Match("+"), Match("-")) will only validate the string "+-". All rules
// have to be validate for the and to be validated. And(Match("++"), Match("-")) will
// for example not validate the string "-+-++" because it will find the "++" first and
// there are no "-" to the left of it.
func And(subRules ...Rule) and {
	return and{
		rules:  subRules,
		center: 0,
	}
}

func AndWithCenter(center int, subRules ...Rule) and {
	return and{
		rules:  subRules,
		center: center,
	}
}

func (r and) Validate(input string, from, to int, fromLeft bool, rules Rules) (RuleResult, error) {
	subRulesMatched := make([]RuleResult, len(r.rules))

	result, err := r.rules[r.center].Validate(input, from, to, fromLeft, rules)
	if err != nil || result.RuleType == "" {
		return RuleResult{}, err
	}
	subRulesMatched[r.center] = result

	left := result.From
	for i := r.center - 1; i >= 0; i-- {
		if left <= from {
			return RuleResult{}, nil
		}

		curResult, err := r.rules[i].Validate(input, from, left, true /* find from left */, rules)
		if err != nil || curResult.RuleType == "" {
			return RuleResult{}, err
		}
		subRulesMatched[i] = curResult
		left = curResult.From
	}

	right := result.To
	for i := r.center + 1; i < len(r.rules); i++ {
		if left >= to {
			return RuleResult{}, nil
		}

		curResult, err := r.rules[i].Validate(input, right, to, false /* find from right */, rules)
		if err != nil || curResult.RuleType == "" {
			return RuleResult{}, err
		}

		subRulesMatched[i] = curResult
		left = curResult.To
	}

	return RuleResult{
		From:            from,
		To:              to,
		RuleType:        AndType,
		SubRulesMatched: subRulesMatched,
	}, nil
}

/*
Rules{
	"main": rule.Name("expression"),
	"expression": rule.Or(rule.Name("binaryOperation"), rule.Name("number"),
	"binaryOperation": rule.AndWithCenter(1, rule.Name("expression"), rule.Name("binaryOperator"), rule.Name("expression")),
	"binaryOperator": rule.Or(rule.FindFirst("*", "/"), rule.FindFirst("+", "-"")),
	"number": rule.Match("\d+"),
}
*/
