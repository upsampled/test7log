package test7log

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

type TestFieldLogger struct {
	t *testing.T
	*logrus.Logger
}

func (t *TestFieldLogger) Write(in []byte) (int, error) {
	t.t.Log(string(in))
	return len(in), nil
}

func (t *TestFieldLogger) Fatalf(format string, any ...any) {
	t.t.Fatalf(format, any...)
}

func (t *TestFieldLogger) Fatal(args ...any) {
	t.t.Fatal(args...)
}

func (t *TestFieldLogger) SetColor(enabled bool) {
	t.Logger.Formatter.(*logrus.TextFormatter).ForceColors = enabled
}

func NewTestFieldLogger(t *testing.T) *TestFieldLogger {
	out := &TestFieldLogger{t: t}

	lg := logrus.StandardLogger()

	if v := os.Getenv("TEST7LOG_OUT"); v != "STDERR" {
		lg.Out = out
	}
	lg.Formatter.(*logrus.TextFormatter).ForceColors = true
	out.Logger = lg

	lvl, err := logrus.ParseLevel(os.Getenv("TEST7LOG_LVL"))
	if err != nil {
		lvl = logrus.DebugLevel
	}

	out.Logger.Level = lvl

	return out
}
