package suparser

import (
	"errors"
	"suparser/rule"
)

var (
	ErrMissingMainRule = errors.New("rules missing main rule")
	ErrInvalidFromTo   = errors.New("invalid from to in rule result")
)

func getErrors(rules rule.Rules) error {
	mainRule, ok := rules["main"]
	if !ok {
		return ErrMissingMainRule
	}

	return mainRule.GetError(rules)
}
