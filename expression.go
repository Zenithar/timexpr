// Package timexpr is a PEG grammar for parsing time expressions.
package timexpr

import (
	"fmt"
	"time"

	"github.com/Zenithar/timexpr/internal/parser"
)

// Parse parses a time expression and returns the time it represents.
func Parse(input string) (time.Time, error) {
	return parse(input, time.Now())
}

// ParseWithReference parses a time expression and returns the time it represents.
// The reference time is used to evaluate relative time expressions.
func ParseWithReference(input string, referenceTime time.Time) (time.Time, error) {
	if referenceTime.IsZero() {
		// If the reference time is zero, use the current time as the reference time.
		// This is done to ensure that relative time expressions are evaluated correctly.
		return parse(input, time.Now())
	}

	return parse(input, referenceTime)
}

// parse parses a time expression and returns the time it represents.
func parse(input string, referenceTime time.Time) (time.Time, error) {
	// Parse the input
	expr, err := parser.Parse("timexpr", []byte(input))
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time expression: %w", err)
	}

	// Evaluate the expression
	switch t := expr.(type) {
	case time.Time:
		return t, nil
	case parser.TimeOffset:
		ts := t.Apply(referenceTime)
		return ts, nil
	default:
	}

	return time.Time{}, fmt.Errorf("unexpected expression type: %T", expr)
}
