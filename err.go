package execs


import (
	"errors"
)


var (
	errAlreadyRun    = errors.New("Already Run")
	errAlreadySet    = errors.New("Already Set")
	errEmptyPipeline = errors.New("Empty Pipeline")
	errNilInlineFunc = errors.New("Nil Inline Func")
	errNilReceiver   = errors.New("Nil Receiver")
)
