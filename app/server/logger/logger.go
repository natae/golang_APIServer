package logger

import (
	"fmt"
)

type LogLevel int

const (
	Debug = iota
	Info
	Warning
	Error
)

func InitLogger(logLevel LogLevel) {
	_logger = &Logger{
		logLevel: logLevel,
	}
}

type Logger struct {
	logLevel LogLevel
}

func (l *Logger) log(logLevel LogLevel, printStackTrace bool, args ...interface{}) {
	if l.logLevel > logLevel {
		return
	}

	var prefix string
	switch logLevel {
	case Debug:
		prefix = "Debug"
	case Info:
		prefix = "Info"
	case Warning:
		prefix = "Warning"
	case Error:
		prefix = "Error"
	}

	fmt.Printf("[%s] ", prefix)
	fmt.Print(args...)
	fmt.Printf("\n")
}

func LogDebugF(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	_logger.log(Debug, false, message)
}

func LogInfoF(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	_logger.log(Info, false, message)
}

func LogWarningF(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	_logger.log(Warning, false, message)
}

func LogErrorF(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	_logger.log(Error, false, message)
}

func LogDebug(args ...interface{}) {
	_logger.log(Debug, false, args...)
}

func LogInfo(args ...interface{}) {
	_logger.log(Info, false, args...)
}

func LogWarning(args ...interface{}) {
	_logger.log(Warning, false, args...)
}

func LogError(args ...interface{}) {
	_logger.log(Error, false, args...)
}

var _logger *Logger
