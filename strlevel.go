// Package slogstringlevel provides a string that implements slog.Leveler.
// Useful for unmarshalling from a string value in a config (e.g. via spf13/viper).
package slogstringlevel

import "log/slog"

// SlogStringLevel is a string that implements slog.Leveler.
type SlogStringLevel string

const (
	LevelDebug SlogStringLevel = "debug"
	LevelInfo  SlogStringLevel = "info"
	LevelWarn  SlogStringLevel = "warn"
	LevelError SlogStringLevel = "error"
)

// Level returns the slog.Level corresponding to the SlogStringLevel.
func (l SlogStringLevel) Level() slog.Level {
	switch l {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

var _ slog.Leveler = (*SlogStringLevel)(nil)
