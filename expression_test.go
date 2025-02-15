package timexpr

import (
	"testing"
	"time"

	"github.com/Zenithar/timexpr/internal/parser"
)

type staticClock struct {
	now time.Time
}

func (c *staticClock) Now() time.Time {
	return c.now
}

func TestParse(t *testing.T) {
	// Set a static clock for testing
	now := time.Date(2021, 10, 10, 10, 0, 0, 0, time.UTC)
	parser.SetClock(&staticClock{now: time.Date(2021, 10, 10, 10, 0, 0, 0, time.UTC)})

	tests := []struct {
		input    string
		expected time.Time
		hasError bool
	}{
		{
			input:    "2023-10-10T10:00:00Z",
			expected: time.Date(2023, 10, 10, 10, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "now",
			expected: time.Date(2021, 10, 10, 10, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "today",
			expected: time.Date(2021, 10, 10, 0, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "yesterday",
			expected: time.Date(2021, 10, 9, 0, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "tomorrow",
			expected: time.Date(2021, 10, 11, 0, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "6s ago",
			expected: time.Date(2021, 10, 10, 9, 59, 54, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "6 seconds ago",
			expected: time.Date(2021, 10, 10, 9, 59, 54, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "6m ago",
			expected: time.Date(2021, 10, 10, 9, 54, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "24h ago",
			expected: time.Date(2021, 10, 9, 10, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "6d ago",
			expected: time.Date(2021, 10, 4, 10, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "6M ago",
			expected: time.Date(2021, 4, 10, 10, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "6y ago",
			expected: time.Date(2015, 10, 10, 10, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "6y later",
			expected: time.Date(2027, 10, 10, 10, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "next hour",
			expected: time.Date(2021, 10, 10, 11, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "next 2h",
			expected: time.Date(2021, 10, 10, 12, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "last 2h",
			expected: time.Date(2021, 10, 10, 8, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "last week",
			expected: time.Date(2021, 10, 3, 10, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			input:    "",
			expected: time.Time{},
			hasError: true,
		},
		{
			input:    "6y 6M 6d 6h 6m 6s ago",
			expected: time.Time{},
			hasError: true,
		},
		{
			input:    "invalid time expression",
			expected: time.Time{},
			hasError: true,
		},
	}

	for _, test := range tests {
		result, err := parse(test.input, now)
		if test.hasError {
			if err == nil {
				t.Errorf("expected error for input %q, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error for input %q: %v", test.input, err)
			}
			if !result.Equal(test.expected) {
				t.Errorf("for input %q, expected %v, but got %v", test.input, test.expected, result)
			}
		}
	}
}

func FuzzParse(f *testing.F) {
	now := time.Date(2021, 10, 10, 10, 0, 0, 0, time.UTC)
	corpus := []string{"2023-10-10T10:00:00Z", "6y ago", "6M ago", "6d ago", "6m ago", "6s ago", "24h ago", "invalid time expression"}
	for _, input := range corpus {
		f.Add(input)
	}
	f.Fuzz(func(t *testing.T, input string) {
		result, err := parse(input, now)
		if result != nil && err != nil {
			t.Errorf("unexpected error for input %q: %v", input, err)
		}
	})
}

func BenchmarkParse(b *testing.B) {
	now := time.Date(2021, 10, 10, 10, 0, 0, 0, time.UTC)
	input := "6y ago"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parse(input, now)
	}
}
