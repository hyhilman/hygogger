package pq_logger

type LevelString string
type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

const (
	DebugLevelString LevelString = "DEBUG"
	InfoLevelString  LevelString = "INFO"
	WarnLevelString  LevelString = "WARN"
	ErrorLevelString LevelString = "ERROR"
	FatalLevelString LevelString = "FATAL"
	PanicLevelString LevelString = "PANIC"
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return string(DebugLevelString)
	case InfoLevel:
		return string(InfoLevelString)
	case WarnLevel:
		return string(WarnLevelString)
	case ErrorLevel:
		return string(ErrorLevelString)
	case FatalLevel:
		return string(FatalLevelString)
	case PanicLevel:
		return string(PanicLevelString)
	default:
		panic("invalid level")
	}
}
