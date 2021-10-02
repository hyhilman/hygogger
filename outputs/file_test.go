package outputs

import (
	"testing"
	"time"
)

func TestFileOutput(t *testing.T) {
	var testDir = "/tmp/pq-logger"

	var ff = []Output{
		&FileOutput{
			Path:     testDir + "/test-no-rotation.log",
			Rotation: 0,
		},
		&FileOutput{
			Path:     testDir + "/test-rotation.log",
			Rotation: time.Second * 10,
		},
	}

	var msg = []byte("my test\n")

	for _, output := range ff {
		bl, err := output.Write(msg)
		if err != nil {
			t.Error(err)
		}

		if bl != len(msg) {
			t.Errorf("write %d bytes instead %d bytes", bl, len(msg))
		} else {
			t.Log("message written successfully")
		}

		if err := output.Close(); err != nil {
			t.Error(err)
		} else {
			t.Log("file closed successfully")
		}
	}
}

func TestFileOutput_Invalid(t *testing.T) {
	var testDir = "/root/pq-logger"

	var ff = []Output{
		&FileOutput{
			Path:     testDir + "/test-no-rotation.log",
			Rotation: 0,
		},
		&FileOutput{
			Path:     testDir + "/test-rotation.log",
			Rotation: time.Second * 10,
		},
	}

	var msg = []byte("my test\n")

	for _, output := range ff {
		bl, err := output.Write(msg)
		if err == nil {
			t.Error("should be returning error")
		} else {
			t.Log("success returning error")
		}

		if bl == len(msg) {
			t.Errorf("write should be 0 byte instead %d bytes", bl)
		} else {
			t.Log("message not written successfully")
		}

		if err := output.Close(); err == nil {
			t.Error("the writer should not be able to close because never initialized")
		} else {
			t.Log("writer not closed successfully")
		}
	}
}
