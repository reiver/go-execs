package execs


import (
	"io"
)


// Inline can be used to create an inline garbage-free Executor.
//
// For example:
//
//	inline := execs.Inline{ Func: func(stdin io.Reader, stdout io.WriteCloser, stderr io.WriteCloser) error {
//		
//		if nil == stdin {
//			return errors.New("stdin is nil")
//		}
//		
//		if nil == stdout {
//			return errors.New("stdout is nil")
//		}
//		
//		//if nil == stderr {
//		//	return errors.New("stderr is nil")
//		//}
//              
//		                                  
//		buffer, err := ioutil.ReadAll(stdin)
//		if nil != err {
//			return err
//		}
//		
//		
//		buffer = bytes.ToUpper(buffer)
//		
//		oi.LongWrite(stdout, buffer)
//		
//		return nil
//	}
//	
//	var executor execs.Executor = inline
type Inline struct {
	Func func(io.Reader, io.WriteCloser, io.WriteCloser)error

	stdin  io.Reader
	stdout io.WriteCloser
	stderr io.WriteCloser
}



// Run calls does some "book-keeping" and then inline.Func.
func (inline *Inline) Run() (err error) {
	if nil == inline {
		return errNilReceiver
	}

	if nil == inline.Func {
		return errNilInlineFunc
	}


	stdin  := inline.stdin
	stdout := inline.stdout
	stderr := inline.stderr

	defer stdout.Close()
	defer stderr.Close()
	return inline.Func(stdin, stdout, stderr)
}



func (inline *Inline) SetStdin(r io.Reader) error {
	if nil == inline {
		return errNilReceiver
	}

	if nil != inline.stdin {
		return errAlreadySet
	}

	inline.stdin = r

	return nil
}

func (inline *Inline) SetStdout(w io.WriteCloser) error {
	if nil == inline {
		return errNilReceiver
	}

	if nil != inline.stdout {
		return errAlreadySet
	}

	inline.stdout = w

	return nil
}

func (inline *Inline) SetStderr(w io.WriteCloser) error {
	if nil == inline {
		return errNilReceiver
	}

	if nil != inline.stderr {
		return errAlreadySet
	}

	inline.stderr = w

	return nil
}









func (inline *Inline) StdinPipe() (io.WriteCloser, error) {
	if nil == inline {
		return nil, errNilReceiver
	}

	if nil != inline.stdin {
		return nil, errAlreadySet
	}


	var stdinPipe *io.PipeWriter
	inline.stdin, stdinPipe = io.Pipe()

	return stdinPipe, nil
}

func (inline *Inline) StdoutPipe() (io.ReadCloser, error) {
	if nil == inline {
		return nil, errNilReceiver
	}

	if nil != inline.stdout {
		return nil, errAlreadySet
	}


	var stdoutPipe *io.PipeReader
	stdoutPipe, inline.stdout = io.Pipe()

	return stdoutPipe, nil
}

func (inline *Inline) StderrPipe() (io.ReadCloser, error) {
	if nil == inline {
		return nil, errNilReceiver
	}

	if nil != inline.stderr {
		return nil, errAlreadySet
	}


	var stderrPipe *io.PipeReader
	stderrPipe, inline.stderr = io.Pipe()

	return stderrPipe, nil
}
