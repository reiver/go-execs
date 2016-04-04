package execs


import (
	"io"
)


type Enhancer interface {
	Executor

	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
}


// EnhancedExecutor can be used to turn an Executor into an Enhancer.
type EnhancedExecutor struct {
	Executor Executor
}









func (enhanced *EnhancedExecutor) Run() error {
	return enhanced.Executor.Run()
}









func (enhanced *EnhancedExecutor) SetStdin(r io.Reader) error {
	return enhanced.Executor.SetStdin(r)
}

func (enhanced *EnhancedExecutor) SetStdout(w io.WriteCloser) error {
	return enhanced.Executor.SetStdout(w)
}

func (enhanced *EnhancedExecutor) SetStderr(w io.WriteCloser) error {
	return enhanced.Executor.SetStderr(w)
}



func (enhanced *EnhancedExecutor) StdinPipe() (io.WriteCloser, error) {
	if nil == enhanced {
		return nil, errNilReceiver
	}


	var stdinPipe *io.PipeWriter
	stdin, stdinPipe := io.Pipe()

	if err := enhanced.SetStdin(stdin); nil != err {
		return nil, err
	}

	return stdinPipe, nil
}

func (enhanced *EnhancedExecutor) StdoutPipe() (io.ReadCloser, error) {
	if nil == enhanced {
		return nil, errNilReceiver
	}


	var stdoutPipe *io.PipeReader
	stdoutPipe, stdout := io.Pipe()

	if err := enhanced.SetStdout(stdout); nil != err {
		return nil, err
	}

        return stdoutPipe, nil
}

func (enhanced *EnhancedExecutor) StderrPipe() (io.ReadCloser, error) {
	if nil == enhanced {
		return nil, errNilReceiver
	}


	var stderrPipe *io.PipeReader
	stderrPipe, stderr := io.Pipe()

	if err := enhanced.SetStderr(stderr); nil != err {
		return nil, err
	}

        return stderrPipe, nil
}









func (enhanced *EnhancedExecutor) MustStdinPipe() io.WriteCloser {

	w, err := enhanced.StdinPipe()
	if nil != err {
		panic(err)
	}

	return w
}

func (enhanced *EnhancedExecutor) MustStdoutPipe() io.ReadCloser {

	r, err := enhanced.StdoutPipe()
	if nil != err {
		panic(err)
	}

	return r
}

func (enhanced *EnhancedExecutor) MustStderrPipe() io.ReadCloser {

	r, err := enhanced.StderrPipe()
	if nil != err {
		panic(err)
	}

	return r
}
