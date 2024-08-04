package logger

type ILog interface {
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Warn(format string, v ...any)
	CError(format string, v ...any)
	Error(format string, v ...any)
}
