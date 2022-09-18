package errors

import (
	"errors"
	"fmt"
	"runtime"
)

func New(message string) error {
	stackTraceBuf := make([]byte, 1<<10)
	runtime.Stack(stackTraceBuf, true)

	return errors.New(fmt.Sprintf("%s\n, %s", message, stackTraceBuf))
}
