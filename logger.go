package hygogger

import (
	"fmt"
	"github.com/hyhilman/hygogger/outputs"
	"log"
	"os"
)

var pool []logger

var MessagesHandlerFactory func(output outputs.Output, level Level) MessageHandler

type logger struct {
	debug MessageHandler
	info  MessageHandler
	warn  MessageHandler
	error MessageHandler
	fatal MessageHandler
	panic MessageHandler
}

type MessageHandler func(string) error

var exit = func(code interface{}) {
	os.Exit(code.(int))
}

func NewLogger(output outputs.Output, level Level) {
	if MessagesHandlerFactory == nil {
		MessagesHandlerFactory = func(output outputs.Output, level Level) MessageHandler {
			l := log.New(output, fmt.Sprintf("[%s]", level), log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
			return func(msg string) (err error) {
				return l.Output(3, msg)
			}
		}
	}

	l := logger{}

	switch level {
	case DebugLevel:
		l.debug = MessagesHandlerFactory(output, DebugLevel)
		fallthrough
	case InfoLevel:
		l.info = MessagesHandlerFactory(output, InfoLevel)
		fallthrough
	case WarnLevel:
		l.warn = MessagesHandlerFactory(output, WarnLevel)
		fallthrough
	case ErrorLevel:
		l.error = MessagesHandlerFactory(output, ErrorLevel)
		fallthrough
	default:
		l.fatal = MessagesHandlerFactory(output, FatalLevel)
		l.panic = MessagesHandlerFactory(output, PanicLevel)
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
