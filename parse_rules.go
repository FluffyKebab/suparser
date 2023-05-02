package suparser

import (
	"fmt"
	"strings"
)

type InvalidRulesError struct {
	Line   int
	Reason string
}

func (e InvalidRulesError) Error() string {
	return fmt.Sprintf("invalid rule on line %v: %s", e.Line, e.Reason)
}

type rule interface{}

//func parseRules(rules string) (map[string]rule, error)

// rulesToRuleMap transforms the list of rules into a map of the rule
// name and rule content.
func rulesToRuleMap(rules string) (map[string]string, error) {
	ruleMap := make(map[string]string)

	lines := strings.Split(rules, "\n")
	for i, line := range lines {
		// Skip if line has no content.
		if strings.TrimSpace(line) == "" {
			continue
		}

		colonSplits := strings.Split(line, ":")
		if len(colonSplits) < 2 {
			return nil, InvalidRulesError{Line: i + 1, Reason: `no ":"`}
		}

		if len(colonSplits) > 2 {
			return nil, InvalidRulesError{Line: i + 1, Reason: `multiple ":"`}
		}

		ruleName := strings.TrimSpace(colonSplits[0])
		ruleContent := strings.TrimSpace(colonSplits[1])

		if _, ok := ruleMap[ruleName]; ok {
			return nil, InvalidRulesError{Line: i + 1, Reason: fmt.Sprintf(`rule "%s" has multiple declarations`, ruleName)}
		}

		if ruleName == "" {
			return nil, InvalidRulesError{Line: i + 1, Reason: "missing rule name"}
		}

		if ruleContent == "" {
			return nil, InvalidRulesError{Line: i + 1, Reason: "missing rule content"}
		}

		ruleMap[ruleName] = ruleContent
	}

	return ruleMap, nil
}
