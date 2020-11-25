// Package log wrapped logrus functions
package log

import (
	"encoding/json"
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// MyLogger extend logrus.Logger
type MyLogger struct {
	*logrus.Logger
}

// Singleton logger
var singletonLogger = &MyLogger{
	Logger: logrus.New(),
}

// Logger return singleton logger
func Logger() *MyLogger {
	return singletonLogger
}

// Print output message of print level
func Print(i ...interface{}) {
	singletonLogger.Print(i...)
}

// Printf output format message of print level
func Printf(format string, i ...interface{}) {
	singletonLogger.Printf(format, i...)
}

// Printj output json of print level
func Printj(j log.JSON) {
	singletonLogger.Printj(j)
}

// Debug output message of debug level
func Debug(i ...interface{}) {
	singletonLogger.Debug(i...)
}

// Debugf output format message of debug level
func Debugf(format string, args ...interface{}) {
	singletonLogger.Debugf(format, args...)
}

// Debugj output json of debug level
func Debugj(j log.JSON) {
	singletonLogger.Debugj(j)
}

// Info output message of info level
func Info(i ...interface{}) {
	singletonLogger.Info(i...)
}

// Infof output format message of info level
func Infof(format string, args ...interface{}) {
	singletonLogger.Infof(format, args...)
}

// Infoj output json of info level
func Infoj(j log.JSON) {
	singletonLogger.Infoj(j)
}

// Warn output message of warn level
func Warn(i ...interface{}) {
	singletonLogger.Warn(i...)
}

// Warnf output format message of warn level
func Warnf(format string, args ...interface{}) {
	singletonLogger.Warnf(format, args...)
}

// Warnj output json of warn level
func Warnj(j log.JSON) {
	singletonLogger.Warnj(j)
}

// Error output message of error level
func Error(i ...interface{}) {
	singletonLogger.Error(i...)
}

// Errorf output format message of error level
func Errorf(format string, args ...interface{}) {
	singletonLogger.Errorf(format, args...)
}

// Errorj output json of error level
func Errorj(j log.JSON) {
	singletonLogger.Errorj(j)
}

// Fatal output message of fatal level
func Fatal(i ...interface{}) {
	singletonLogger.Fatal(i...)
}

// Fatalf output format message of fatal level
func Fatalf(format string, args ...interface{}) {
	singletonLogger.Fatalf(format, args...)
}

// Fatalj output json of fatal level
func Fatalj(j log.JSON) {
	singletonLogger.Fatalj(j)
}

// Panic output message of panic level
func Panic(i ...interface{}) {
	singletonLogger.Panic(i...)
}

// Panicf output format message of panic level
func Panicf(format string, args ...interface{}) {
	singletonLogger.Panicf(format, args...)
}

// Panicj output json of panic level
func Panicj(j log.JSON) {
	singletonLogger.Panicj(j)
}

// To logrus.Level
func toLogrusLevel(level log.Lvl) logrus.Level {
	switch level {
	case log.DEBUG:
		return logrus.DebugLevel
	case log.INFO:
		return logrus.InfoLevel
	case log.WARN:
		return logrus.WarnLevel
	case log.ERROR:
		return logrus.ErrorLevel
	}

	return logrus.InfoLevel
}

// To Echo.log.lvl
func toEchoLevel(level logrus.Level) log.Lvl {
	switch level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	}

	return log.OFF
}

// Output return logger io.Writer
func (l *MyLogger) Output() io.Writer {
	return l.Out
}

// SetOutput logger io.Writer
func (l *MyLogger) SetOutput(w io.Writer) {
	l.Out = w
}

// Level return logger level
func (l *MyLogger) Level() log.Lvl {
	return toEchoLevel(l.Logger.Level)
}

// SetLevel logger level
func (l *MyLogger) SetLevel(v log.Lvl) {
	l.Logger.Level = toLogrusLevel(v)
}

// SetHeader logger header
// Managed by Logrus itself
// This function do nothing
func (l *MyLogger) SetHeader(h string) {
	// do nothing
}

// Formatter return logger formatter
func (l *MyLogger) Formatter() logrus.Formatter {
	return l.Logger.Formatter
}

// SetFormatter logger formatter
// Only support logrus formatter
func (l *MyLogger) SetFormatter(formatter logrus.Formatter) {
	l.Logger.Formatter = formatter
}

// Prefix return logger prefix
// This function do nothing
func (l *MyLogger) Prefix() string {
	return ""
}

// SetPrefix logger prefix
// This function do nothing
func (l *MyLogger) SetPrefix(p string) {
	// do nothing
}

// Print output message of print level
func (l *MyLogger) Print(i ...interface{}) {
	l.Logger.Print(i...)
}

// Printf output format message of print level
func (l *MyLogger) Printf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

// Printj output json of print level
func (l *MyLogger) Printj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Println(string(b))
}

// Debug output message of debug level
func (l *MyLogger) Debug(i ...interface{}) {
	l.Logger.Debug(i...)
}

// Debugf output format message of debug level
func (l *MyLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

// Debugj output message of debug level
func (l *MyLogger) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Debugln(string(b))
}

// Info output message of info level
func (l *MyLogger) Info(i ...interface{}) {
	l.Logger.Info(i...)
}

// Infof output format message of info level
func (l *MyLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

// Infoj output json of info level
func (l *MyLogger) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Infoln(string(b))
}

// Warn output message of warn level
func (l *MyLogger) Warn(i ...interface{}) {
	l.Logger.Warn(i...)
}

// Warnf output format message of warn level
func (l *MyLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

// Warnj output json of warn level
func (l *MyLogger) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Warnln(string(b))
}

// Error output message of error level
func (l *MyLogger) Error(i ...interface{}) {
	l.Logger.Error(i...)
}

// Errorf output format message of error level
func (l *MyLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

// Errorj output json of error level
func (l *MyLogger) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Errorln(string(b))
}

// Fatal output message of fatal level
func (l *MyLogger) Fatal(i ...interface{}) {
	l.Logger.Fatal(i...)
}

// Fatalf output format message of fatal level
func (l *MyLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

// Fatalj output json of fatal level
func (l *MyLogger) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Fatalln(string(b))
}

// Panic output message of panic level
func (l *MyLogger) Panic(i ...interface{}) {
	l.Logger.Panic(i...)
}

// Panicf output format message of panic level
func (l *MyLogger) Panicf(format string, args ...interface{}) {
	l.Logger.Panicf(format, args...)
}

// Panicj output json of panic level
func (l *MyLogger) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Panicln(string(b))
}
