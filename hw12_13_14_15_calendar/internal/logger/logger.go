package logger

import (
	"fmt"
	"strings"
)

type Logger struct {
	level Level
}

type Level uint32

const (
	ErrorLevel = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

func New(level string) *Logger {
	parsedLevel, err := ParseLevel(level)
	if err != nil {
		fmt.Println("failed to create logger:", err)
	}
	return &Logger{
		level: parsedLevel,
	}
}

func (l Logger) Info(msg string) {
	l.log(InfoLevel, msg)
}

func (l Logger) Error(msg string) {
	l.log(ErrorLevel, msg)
}

func (l Logger) Warn(msg string) {
	l.log(WarnLevel, msg)
}

func (l Logger) Debug(msg string) {
	l.log(DebugLevel, msg)
}

func (l Logger) log(level Level, msg string) {
	if l.level >= level {
		fmt.Println(msg)
	}
}

func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid log Level: %q", lvl)
}

func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case DebugLevel:
		return []byte("debug"), nil
	case InfoLevel:
		return []byte("info"), nil
	case WarnLevel:
		return []byte("warning"), nil
	case ErrorLevel:
		return []byte("error"), nil
	}

	return nil, fmt.Errorf("not a valid log level %d", level)
}
