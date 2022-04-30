package policy_enforcer

import (
	"errors"
)

type EnforcerError error

var (
	// ErrUnsafeVarError
	ErrUnsafeVarError EnforcerError = errors.New("Unsafe Variable Error")
	// ErrParseError
	ErrParseError EnforcerError = errors.New("Parse Error")
)
