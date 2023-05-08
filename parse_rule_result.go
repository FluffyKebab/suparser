package suparser

import "suparser/rule"

func parseRuleResult(text string, ruleResult rule.RuleResult) ([]Node, error) {
	subNodes := make([]Node, 0, len(ruleResult.SubRulesMatched))
	for _, r := range ruleResult.SubRulesMatched {
		nodes, err := parseRuleResult(text, r)
		if err != nil {
			return nil, err
		}

		subNodes = append(subNodes, nodes...)
	}

	if ruleResult.RuleType == rule.NameType {
		nodeText, err := getSubstring(text, ruleResult.From, ruleResult.To)
		if err != nil {
			return nil, err
		}

		return []Node{
			{
				Name:     ruleResult.RuleName,
				Text:     nodeText,
				SubNodes: subNodes,
			},
		}, nil
	}

	return subNodes, nil
}

func getSubstring(s string, from, to int) (string, error) {
	// TODO: check if invalid.

	return string([]rune(s)[from:to]), nil
}
