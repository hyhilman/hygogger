package pq_logger

import (
	"pq-logger/outputs"
	"testing"
	"time"
)

var console = &outputs.ConsoleOutput{}
var file = &outputs.FileOutput{
	Path:     "/tmp/pq-logger/test.log",
	Rotation: time.Second * 100,
}

func TestLogger(t *testing.T) {
	NewLogger(console, DebugLevel)
	NewLogger(file, WarnLevel)

	msg := "my test"
	Debug(msg)
	Info(msg)
	Warn(msg)
	Error(msg)
}

func TestLogger_Fatal(t *testing.T) {
	exit = func(code interface{}) {
		if v, ok := code.(int); !ok {
			t.Errorf("the code is not an int")
		} else if v != 1 {
			t.Errorf("the value is %d instead 1", v)
		} else {
			t.Log("successfully calling os exit")
		}
	}

	msg := "my test"
	Fatal(msg)
}

func TestLogger_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("log should be thworing panic")
		} else {
			t.Log("log panic success")
		}
	}()

	msg := "my test"
	Panic(msg)
}
