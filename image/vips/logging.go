package vips

// #include <glib.h>
import "C"
import "os"

type LogLevel int

// The logging verbosity levels classify and filter logging messages.
// From most to least verbose, they are debug, info, message, warning, critical and error.
const (
	LogLevelError    LogLevel = C.G_LOG_LEVEL_ERROR
	LogLevelCritical LogLevel = C.G_LOG_LEVEL_CRITICAL
	LogLevelWarning  LogLevel = C.G_LOG_LEVEL_WARNING
	LogLevelMessage  LogLevel = C.G_LOG_LEVEL_MESSAGE
	LogLevelInfo     LogLevel = C.G_LOG_LEVEL_INFO
	LogLevelDebug    LogLevel = C.G_LOG_LEVEL_DEBUG
)

//export fyntrixLoggingHandler
func fyntrixLoggingHandler(messageDomain *C.char, messageLevel C.int, message *C.char) {
	_vipsLogging(C.GoString(message), LogLevel(messageLevel), C.GoString(messageDomain))
}

func _vipsLogging(message string, level LogLevel, domain string) {
	switch level {
	case LogLevelError:
		_logging.Error(message, "domain", domain)
	case LogLevelCritical:
		_logging.Error(message, "domain", domain)
		os.Exit(1)
	case LogLevelWarning:
		_logging.Warn(message, "domain", domain)
	case LogLevelInfo, LogLevelMessage:
		_logging.Info(message, "domain", domain)
	case LogLevelDebug:
		_logging.Debug(message, "domain", domain)
	}
}
