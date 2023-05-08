package rule

import "errors"

var (
	ErrAndCenterInvalid = errors.New("invalid center in and")
	ErrRuleNotDefined   = errors.New("rule not defined")
	ErrInvalidRegex     = errors.New("invalid regex")
	ErrInvalidFromTo    = errors.New("invalid from to")
)
