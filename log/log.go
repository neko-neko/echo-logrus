package log

import (
	"encoding/json"
	"io"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/gommon/log"
)

// Logger extend logrus.Logger
type Logger struct {
	*logrus.Logger
}

// Singleton logger
var singletonLogger = &Logger{
	Logger: logrus.New(),
}

// Logger return singleton logger
func Logger() *Logger {
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

// Return logger io.Writer
func (l *Logger) Output() io.Writer {
	return l.Out
}

// Set logger io.Writer
func (l *Logger) SetOutput(w io.Writer) {
	l.Out = w
}

// Return logger level
func (l *Logger) Level() log.Lvl {
	return toEchoLevel(l.Logger.Level)
}

// Set logger level
func (l *Logger) SetLevel(v log.Lvl) {
	l.Logger.Level = toLogrusLevel(v)
}

// Return logger fomatter
func (l *Logger) Formatter() logrus.Formatter {
	return l.Logger.Formatter
}

// Set logger formatter
// Only support logrus formatter
func (l *Logger) SetFormatter(formatter logrus.Formatter) {
	l.Logger.Formatter = formatter
}

// Return logger prefix
// This function do nothing
func (l *Logger) Prefix() string {
	return ""
}

// Set logger prefix
// This function do nothing
func (l *Logger) SetPrefix(p string) {
	// do nothing
}

// Output message of print level
func (l *Logger) Print(i ...interface{}) {
	l.Logger.Print(i...)
}

// Output format message of print level
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

// Output json of print level
func (l *Logger) Printj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Println(string(b))
}

// Output message of debug level
func (l *Logger) Debug(i ...interface{}) {
	l.Logger.Info(i...)
}

// Output format message of debug level
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

// Output message of debug level
func (l *Logger) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Debugln(string(b))
}

// Output message of info level
func (l *Logger) Info(i ...interface{}) {
	l.Logger.Info(i...)
}

// Output format message of info level
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

// Output json of info level
func (l *Logger) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Infoln(string(b))
}

// Output message of warn level
func (l *Logger) Warn(i ...interface{}) {
	l.Logger.Warn(i...)
}

// Output format message of warn level
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

// Output json of warn level
func (l *Logger) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Warnln(string(b))
}

// Output message of error level
func (l *Logger) Error(i ...interface{}) {
	l.Logger.Error(i...)
}

// Output format message of error level
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

// Output json of error level
func (l *Logger) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Errorln(string(b))
}

// Output message of fatal level
func (l *Logger) Fatal(i ...interface{}) {
	l.Logger.Fatal(i...)
}

// Output format message of fatal level
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

// Output json of fatal level
func (l *Logger) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Fatalln(string(b))
}

// Output message of panic level
func (l *Logger) Panic(i ...interface{}) {
	l.Logger.Panic(i...)
}

// Output format message of panic level
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.Logger.Panicf(format, args...)
}

// Output json of panic level
func (l *Logger) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Panicln(string(b))
}
