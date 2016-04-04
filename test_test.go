package execs


import (
	"github.com/reiver/go-oi"

	"bytes"
	"io"
	"io/ioutil"
)


type NopExecutor struct {
	alreadyRun bool

	stdin  io.Reader
	stdout io.WriteCloser
	stderr io.WriteCloser
}


func (nop *NopExecutor) Run() error {
	if nop.alreadyRun {
		return errAlreadyRun
	}

	return nil
}



func (nop *NopExecutor) StdinPipe() (io.WriteCloser, error) {
	if nop.alreadyRun {
		return nil, errAlreadyRun
	}

	if nil != nop.stdin {
		return nil, errAlreadySet
	}


	var buffer bytes.Buffer
	nop.stdin = &buffer

	return oi.WriteNopCloser(ioutil.Discard), nil
}
func (nop *NopExecutor) StdoutPipe() (io.ReadCloser, error) {
	if nop.alreadyRun {
		return nil, errAlreadyRun
	}

	if nil != nop.stdout {
		return nil, errAlreadySet
	}


	nop.stdout = oi.WriteNopCloser(ioutil.Discard)

	var buffer bytes.Buffer
	return ioutil.NopCloser(&buffer), nil
}
func (nop *NopExecutor) StderrPipe() (io.ReadCloser, error) {
	if nop.alreadyRun {
		return nil, errAlreadyRun
	}

	if nil != nop.stderr {
		return nil, errAlreadySet
	}


	nop.stderr = oi.WriteNopCloser(ioutil.Discard)

	var buffer bytes.Buffer
	return ioutil.NopCloser(&buffer), nil
}



func (nop *NopExecutor) SetStdin(r io.Reader) error {
	if nop.alreadyRun {
		return errAlreadyRun
	}

	if nil != nop.stdin {
		return errAlreadySet
	}

	nop.stdin = r

	return nil
}

func (nop *NopExecutor) SetStdout(w io.WriteCloser) error {
	if nop.alreadyRun {
		return errAlreadyRun
	}

	if nil != nop.stdout {
		return errAlreadySet
	}

	nop.stdout = w

	return nil
}

func (nop *NopExecutor) SetStderr(w io.WriteCloser) error {
	if nop.alreadyRun {
		return errAlreadyRun
	}

	if nil != nop.stderr {
		return errAlreadySet
	}

	nop.stderr = w

	return nil
}
