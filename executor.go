package execs


import (
	"io"
)


type Executor interface {
	Run() error

	SetStdin(io.Reader) error
	SetStdout(io.WriteCloser) error
	SetStderr(io.WriteCloser) error
}
