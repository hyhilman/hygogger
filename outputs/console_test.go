package outputs

import (
	"os"
	"testing"
)

func TestConsoleOutput(t *testing.T) {
	var testDir = "/tmp/pq-logger"

	var ff = []Output{
		&ConsoleOutput{},
		&ConsoleOutput{},
	}

	var msg = []byte("my test\n")

	defer os.RemoveAll(testDir)

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
