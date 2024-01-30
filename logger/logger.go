package logger

// Logger exposes functionality to write messages in stdout.
type Logger interface {
	// Error is used to send formatted as error message.
	Error(msg string, err error)
	// Debug is used to send formatted debug message.
	Debug(msg string)
	// Fatal is used to send formatted as error message and interact program.
	Fatal(msg string, err error)
}
