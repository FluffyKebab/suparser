package rule

import "errors"

var (
	ErrRuleNotDefined = errors.New("rule not defined")
	ErrInvalidRegex   = errors.New("invalid regex")
)
