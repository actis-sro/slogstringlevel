package slogstringlevel

import (
	"log/slog"
	"testing"
)

func TestLevel(t *testing.T) {
	testCases := []struct {
		desc     string
		given    string
		expected slog.Level
	}{
		{
			desc:     "From 'debug' to slog.LevelDebug",
			given:    "debug",
			expected: slog.LevelDebug,
		},
		{
			desc:     "From 'info' to slog.LevelInfo",
			given:    "info",
			expected: slog.LevelInfo,
		},
		{
			desc:     "From 'warn' to slog.LevelWarn",
			given:    "warn",
			expected: slog.LevelWarn,
		},
		{
			desc:     "From 'error' to slog.LevelError",
			given:    "error",
			expected: slog.LevelError,
		},
		{
			desc:     "From 'invalid' to slog.LevelInfo",
			given:    "invalid",
			expected: slog.LevelInfo,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			strLvl := SlogStringLevel(tC.given)
			actual := strLvl.Level()
			if actual != tC.expected {
				t.Errorf("expected %q, got %q", tC.expected, actual)
			}
		})
	}
}
