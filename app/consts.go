package golog

// Level ...
type Level int

const (
	// DefaultLevel Level
	DefaultLevel = InfoLevel

	// PanicLevel, when there is no recover
	PanicLevel Level = iota
	// FatalLevel, when the error is fatal to the application
	FatalLevel
	// ErrorLevel, when there is a controlled error
	ErrorLevel
	// WarnLevel, when there is a warning
	WarnLevel
	// InfoLevel, when it is a informational message
	InfoLevel
	// DebugLevel, when it is a debugging message
	DebugLevel

	// Special Prefixes
	// Add the level value to the prefix
	LEVEL = "{{LEVEL}}"
	// Add the timestamp value to the prefix
	TIMESTAMP = "{{TIMESTAMP}}"
	// Add the date value to the prefix
	DATE = "{{DATE}}"
	// Add the time value to the prefix
	TIME = "{{TIME}}"
	// Add the client ip address
	IP = "{{IP}}"
	// Add the error trace
	TRACE = "{{TRACE}}"
	// Add the package name
	PACKAGE = "{{PACKAGE}}"
	// Add the file
	FILE = "{{FILE}}"
	// Add the function name
	FUNCTION = "{{FUNCTION}}"
	// Add the debug stack
	STACK = "{{STACK}}"
)
