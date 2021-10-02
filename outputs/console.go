package outputs

import (
	"fmt"
)

type ConsoleOutput struct{}

func (*ConsoleOutput) Write(msg []byte) (int, error) {
	return fmt.Print(string(msg))
}

func (*ConsoleOutput) Close() error {
	return nil
}
