package wildcard_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/wildcard"
)

func TestMatch(t *testing.T) {
	testCases := []struct {
		pattern  string
		haystack string
		expect   bool
	}{
		// special cases
		{
			pattern:  "",
			haystack: "",
			expect:   true,
		},
		{
			pattern:  "",
			haystack: "alphabeta",
			expect:   false,
		},
		{
			pattern:  "a",
			haystack: "",
			expect:   false,
		},
		{
			pattern:  "*",
			haystack: "alphabeta",
			expect:   true,
		},

		// no wildcards
		{
			pattern:  "alpha",
			haystack: "alphabeta",
			expect:   false,
		},
		{
			pattern:  "alphabeta",
			haystack: "alphabeta",
			expect:   true,
		},

		// wildcard as suffix
		{
			pattern:  "alpha*",
			haystack: "alphabeta",
			expect:   true,
		},
		{
			pattern:  "alphabeta*",
			haystack: "alphabeta",
			expect:   true,
		},
		{
			pattern:  "beta*",
			haystack: "alphabeta",
			expect:   false,
		},

		// wildcard as prefix
		{
			pattern:  "*alpha",
			haystack: "alphabeta",
			expect:   false,
		},
		{
			pattern:  "*alphabeta",
			haystack: "alphabeta",
			expect:   true,
		},
		{
			pattern:  "*beta",
			haystack: "alphabeta",
			expect:   true,
		},

		// wildcard in middle
		{
			pattern:  "alpha*beta",
			haystack: "alphabeta",
			expect:   true,
		},
		{
			pattern:  "a*a",
			haystack: "alphabeta",
			expect:   true,
		},

		// multiple wildcards
		{
			pattern:  "*a*",
			haystack: "alphabeta",
			expect:   true,
		},
		{
			pattern:  "*c*",
			haystack: "alphabeta",
			expect:   false,
		},
		{
			pattern:  "a*b*",
			haystack: "alphabeta",
			expect:   true,
		},
		{
			pattern:  "a*b*",
			haystack: "alphabeta",
			expect:   true,
		},
		{
			pattern:  "*a*b",
			haystack: "alphabeta",
			expect:   false,
		},
		{
			pattern:  "*a*b*",
			haystack: "alphabeta",
			expect:   true,
		},

		// invalid casing
		{
			pattern:  "ALPHABETA",
			haystack: "alphabeta",
			expect:   false,
		},
		{
			pattern:  "alphabeta",
			haystack: "ALPHABETA",
			expect:   false,
		},
		{
			pattern:  "*A*B*",
			haystack: "alphabeta",
			expect:   false,
		},
		{
			pattern:  "*a*b*",
			haystack: "ALPHABETA",
			expect:   false,
		},
	}

	for _, tc := range testCases {
		result := wildcard.Match(tc.pattern, tc.haystack)
		argDescription := fmt.Sprintf(`pattern "%s", haystack "%s"`, tc.pattern, tc.haystack)
		assert.Equal(t, tc.expect, result, argDescription)
	}
}

func BenchmarkMatch(b *testing.B) {
	b.ReportAllocs()

	b.Run("no wildcards", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wildcard.Match("alpha", "alphabeta")
		}
	})

	b.Run("single wildcard", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wildcard.Match("alph*eta", "alphabeta")
		}
	})

	b.Run("multiple wildcards", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wildcard.Match("*a*b*", "alphabeta")
		}
	})
}

func TestMatchCI(t *testing.T) {
	testCases := []struct {
		pattern  string
		haystack string
		expect   bool
	}{
		{
			pattern:  "ALPHABETA",
			haystack: "alphabeta",
			expect:   true,
		},
		{
			pattern:  "alphabeta",
			haystack: "ALPHABETA",
			expect:   true,
		},
		{
			pattern:  "*A*B*",
			haystack: "alphabeta",
			expect:   true,
		},
		{
			pattern:  "*a*b*",
			haystack: "ALPHABETA",
			expect:   true,
		},
	}

	for _, tc := range testCases {
		result := wildcard.MatchCI(tc.pattern, tc.haystack)
		argDescription := fmt.Sprintf(`pattern "%s", haystack "%s"`, tc.pattern, tc.haystack)
		assert.Equal(t, tc.expect, result, argDescription)
	}
}

func BenchmarkMatchCI(b *testing.B) {
	b.ReportAllocs()

	b.Run("no wildcards", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wildcard.MatchCI("ALPHA", "alphabeta")
		}
	})

	b.Run("single wildcard", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wildcard.MatchCI("ALPH*ETA", "alphabeta")
		}
	})

	b.Run("multiple wildcards", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wildcard.MatchCI("*A*B*", "alphabeta")
		}
	})
}

func TestMatchSlice(t *testing.T) {
	haystack := []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta", "iota"}

	assert.False(t, wildcard.MatchSlice("foo", nil))
	assert.False(t, wildcard.MatchSlice("foo", []string{}))

	assert.True(t, wildcard.MatchSlice("gamma", haystack))
	assert.True(t, wildcard.MatchSlice("*a", haystack))
	assert.True(t, wildcard.MatchSlice("bet*", haystack))

	assert.False(t, wildcard.MatchSlice("foo", haystack))
	assert.False(t, wildcard.MatchSlice("*ETA", haystack))
	assert.False(t, wildcard.MatchSlice("GAMMA", haystack))
	assert.False(t, wildcard.MatchSlice("BET*", haystack))
}

func TestMatchSliceCI(t *testing.T) {
	haystack := []string{"alpha", "BETA", "gamma", "delta", "epsilon", "zeta", "eta", "theta", "iota"}
	assert.False(t, wildcard.MatchSliceCI("foo", haystack))

	assert.True(t, wildcard.MatchSliceCI("*ETA", haystack))
	assert.True(t, wildcard.MatchSliceCI("GAMMA", haystack))
	assert.True(t, wildcard.MatchSliceCI("BET*", haystack))
	assert.True(t, wildcard.MatchSliceCI("bet*", haystack))
}
