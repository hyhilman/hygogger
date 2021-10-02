package pq_logger

import (
	"fmt"
	"log"
	"os"
	"pq-logger/outputs"
)

var pool []logger

type msgHandler func(string) error

var exit = func(code interface{}) {
	os.Exit(code.(int))
}

var msgHandlerFactory = func(output outputs.Output, level Level) msgHandler {
	l := log.New(output, fmt.Sprintf("[%s]", level), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	return func(msg string) (err error) {
		return l.Output(3, msg)
	}
}

type logger struct {
	debug msgHandler
	info  msgHandler
	warn  msgHandler
	error msgHandler
	fatal msgHandler
	panic msgHandler
}

func NewLogger(output outputs.Output, level Level) {
	l := logger{}

	switch level {
	case DebugLevel:
		l.debug = msgHandlerFactory(output, DebugLevel)
		fallthrough
	case InfoLevel:
		l.info = msgHandlerFactory(output, InfoLevel)
		fallthrough
	case WarnLevel:
		l.warn = msgHandlerFactory(output, WarnLevel)
		fallthrough
	case ErrorLevel:
		l.error = msgHandlerFactory(output, ErrorLevel)
		fallthrough
	default:
		l.fatal = msgHandlerFactory(output, FatalLevel)
		l.panic = msgHandlerFactory(output, PanicLevel)
	}

	pool = append(pool, l)
}

func Debug(msg string) {
	for _, logger := range pool {
		if logger.debug != nil {
			if err := logger.debug(msg); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func Warn(msg string) {
	for _, logger := range pool {
		if logger.warn != nil {
			if err := logger.warn(msg); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func Info(msg string) {
	for _, logger := range pool {
		if logger.info != nil {
			if err := logger.info(msg); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func Error(msg string) {
	for _, logger := range pool {
		if logger.error != nil {
			if err := logger.error(msg); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func Fatal(msg string) {
	for _, logger := range pool {
		if logger.fatal != nil {
			if err := logger.fatal(msg); err != nil {
				fmt.Println(err)
			}
		}
	}

	exit(1)
}

func Panic(msg string) {
	for _, logger := range pool {
		if logger.panic != nil {
			if err := logger.panic(msg); err != nil {
				fmt.Println(err)
			}
		}
	}

	panic(msg)
}
