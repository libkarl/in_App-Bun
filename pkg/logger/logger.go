package logger

import (
	"github.com/sirupsen/logrus"
)

// Logger is interface which satisfies Entry and Log
type Logger interface {
	logrus.FieldLogger
}

var _ Logger = (*logrus.Entry)(nil)
var _ Logger = (*Log)(nil)

// Log is wrapper around logrus
type Log struct {
	*logrus.Logger
}

// New will create new instance if logrus
func New() *Log {
	l := logrus.New()
	l.SetFormatter(logrus.New().Formatter)
	l.SetLevel(logrus.DebugLevel)
	return &Log{Logger: l}
}

// TextFormatter is struct implementing Format interface
// this is useful for fomating logs indifferent environments.
// This formater will format logs for terminal and testing.
type TextFormatter struct {
	// Force disabling colors. For a TTY colors are enabled by default.
	UseColors bool
	// Color scheme to use.
	scheme *compiledColorScheme
}


// NewTextFormatter will create new formtter for logrus
func NewTextFormatter(colors bool) *TextFormatter {
	f := &TextFormatter{}
	f.scheme = noColorsColorScheme
	if colors {
		f.scheme = defaultCompiledColorScheme
	}
	return f
}