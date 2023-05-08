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
// there are no "-" to the right of it.
func And(subRules ...Rule) and {
	return and{
		rules:  subRules,
		center: 0,
	}
}

// AndWithCenter creates an and rule with a specified center. The center must be an index
// in the rules given, and the rule with the center index will be found first. Rules to the
// left of the center rule will be found from the left, witch means that the leftmost match
// is used. The opposite is true for the rules to the right of the center.
//
// AndWithCenter can for example to find a binary operator in an expression. Lets say that
// that we have the rule "AndWithCenter(1, Name("expression"), Name("binaryOperator"),
// Name("expression"))" and that the rule with the name "expression" validates numbers
// or the AndWithCenter rule, and that the rule with the name "binaryOperator" validates
// standard mathematical operators. Then the rule would validate strings like "34+6-1".
// Firstly, since the center is specified as the second rule, the plus symbol is found by
// the "binaryOperator" rule. Then the 34 to the left it is validated because it is a number.
// Because the AndWithCenter rule is a part of the rule with the name "expression", the minus
// will be fond first by and afterwards the thu numbers will be found.
func AndWithCenter(center int, subRules ...Rule) and {
	return and{
		rules:  subRules,
		center: center,
	}
}

// Validate validates if the input strings follows the and rule from "from" to "to". If fromLeft
// is true, the center will be the rightmost match of the center rule. Otherwise it will be the
// leftmost.
func (r and) Validate(input string, from, to int, fromLeft bool, rules Rules) (RuleResult, error) {
	subRulesMatched := make([]RuleResult, len(r.rules))

	result, err := r.rules[r.center].Validate(input, from, to, fromLeft, rules)
	if err != nil || result.RuleType == "" {
		return RuleResult{}, err
	}
	subRulesMatched[r.center] = result

	right := result.From
	for i := r.center - 1; i >= 0; i-- {
		if right < from {
			return RuleResult{}, nil
		}

		curResult, err := r.rules[i].Validate(input, from, right, true /* find from right */, rules)
		if err != nil || curResult.RuleType == "" {
			return RuleResult{}, err
		}
		subRulesMatched[i] = curResult
		right = curResult.From
	}

	left := result.To
	for i := r.center + 1; i < len(r.rules); i++ {
		if left >= to {
			return RuleResult{}, nil
		}

		curResult, err := r.rules[i].Validate(input, left, to, false /* find from left */, rules)
		if err != nil || curResult.RuleType == "" {
			return RuleResult{}, err
		}

		subRulesMatched[i] = curResult
		left = curResult.To
	}

	return RuleResult{
		From:            right,
		To:              left,
		RuleType:        AndType,
		SubRulesMatched: subRulesMatched,
	}, nil
}
