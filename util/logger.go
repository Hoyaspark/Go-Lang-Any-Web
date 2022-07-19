package util

type Logger interface {
	Log(message string, err error)
}

type LoggerFunc func(message string, err error)

func (lf LoggerFunc) Log(message string, err error) {
	lf(message, err)
}
