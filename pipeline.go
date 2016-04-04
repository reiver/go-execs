package execs


import (
	"io"
	"sync"
)


// Pipeline represents a sequence of Executors where each Executor has its stdout connected
// to the stdin of the next Executor.
//
// Individuals familiar with shell interfaces such as "bash", "csh", "sh", "zsh", etc may
// know this concept from shell code like the following:
//
//	cat file.txt | grep something | wc -l
//
// Pipeline gives you a building block to do something similar.
//
// Example usage:
//
//	var pipeline execs.Pipeline
//	
//	pipeline.Executors = []execs.Executor{ex1, ex2, ex3, ex4}
//	
//	err := pipeline.Run()
//	if nil != err {
//		fmt.Printf("ERROR: %v\n", err)
//		os.Exit(1)
//	}
type Pipeline struct {

	Executors []Executor

	alreadyRun bool

	stderr io.WriteCloser
}









// Run calls the Run method on all the Executors it contains and waits for all of them to complete, after
// making sure the stdout of each Executor it contains is connected to the stdin of the next Executor.
//
// In the case where Run returns an error because one of more of the Executors it contains returned an error
// from its Run method, Run will return an error which can be type casted into an execs.Errors.
//
// Example usage:
//
//	var pipeline execs.Pipeline
//	
//	pipeline.Executors = []execs.Executor{ex1, ex2, ex3, ex4}
//	
//	err := pipeline.Run()
//	if nil != err {
//		fmt.Printf("ERROR: %v\n", err)
//		os.Exit(1)
//	}
func (pipeline *Pipeline) Run() error {
	if nil == pipeline {
		return errNilReceiver
	}



	if pipeline.alreadyRun {
		return errAlreadyRun
	}

	pipeline.alreadyRun = true



	executors := pipeline.Executors
	if nil == executors || len(executors) <= 0 {
		return errEmptyPipeline
	}



	var waitGroup sync.WaitGroup

	var errors []error
	var errorsMutex sync.Mutex


	lenExecutorsMinusOne := len(executors)-1
	for i, executor := range executors {
		executor.SetStderr(pipeline.stderr)


		if i != lenExecutorsMinusOne {

			var enhanced Enhancer
			if e, ok := executor.(Enhancer); ok {
				enhanced = e	
			} else {
				enhanced = &EnhancedExecutor{ Executor:executor }
			}

			r, err := enhanced.StdoutPipe()
			if nil != err {
				return err
			}

			nextExecutor := executors[i+1]

			if err := nextExecutor.SetStdin(r); nil != err {
//@TODO: Could we handle this error better. (In a way that we force that which was spawn to terminate?)
				panic(err)
			}
		}


		waitGroup.Add(1)
		go func(ex Executor) {
			defer waitGroup.Done()

			if err := ex.Run(); nil != err {
				errorsMutex.Lock()
				errors = append(errors, err)
				errorsMutex.Unlock()
			}
		}(executor)
	}


	waitGroup.Wait()


	if nil != errors && 0 < len(errors) {
		var errs Errors
		errs.errors = errors

		return errs
	}


	return nil
}









// MustRun is like Run, except it panic()s if there were any errors when trying to run.
//
// Example usage:
//
//	var pipeline execs.Pipeline
//	
//	pipeline.Executors = []execs.Executor{ex1, ex2, ex3, ex4}
//	
//	err := pipeline.MustRun() // This might panic()!
func (pipeline *Pipeline) MustRun() {
	if err := pipeline.Run(); nil != err {
		panic(err)
	}
}









func (pipeline *Pipeline) StdinPipe() (io.WriteCloser, error) {
	if nil == pipeline {
		return nil, errNilReceiver
	}

	if pipeline.alreadyRun {
		return nil, errAlreadyRun
	}

	executors := pipeline.Executors
	if nil == executors || len(executors) <= 0 {
		return nil, errEmptyPipeline
	}

	executor := executors[0]

	var enhanced Enhancer
	if e, ok := executor.(Enhancer); ok {
		enhanced = e	
	} else {
		enhanced = &EnhancedExecutor{ Executor:executor }
	}

	return enhanced.StdinPipe()
}


func (pipeline *Pipeline) StdoutPipe() (io.ReadCloser, error) {
	if nil == pipeline {
		return nil, errNilReceiver
	}

	if pipeline.alreadyRun {
		return nil, errAlreadyRun
	}

	executors := pipeline.Executors
	if nil == executors || len(executors) <= 0 {
		return nil, errEmptyPipeline
	}

	executor := executors[ len(executors)-1 ]

	var enhanced Enhancer
	if e, ok := executor.(Enhancer); ok {
		enhanced = e	
	} else {
		enhanced = &EnhancedExecutor{ Executor:executor }
	}

	return enhanced.StdoutPipe()
}


func (pipeline *Pipeline) StderrPipe() (stderrPipe io.ReadCloser, err error) {
	if nil == pipeline {
		return nil, errNilReceiver
	}

	if pipeline.alreadyRun {
		return nil, errAlreadyRun
	}

	executors := pipeline.Executors
	if nil == executors || len(executors) <= 0 {
		return nil, errEmptyPipeline
	}


	if nil != pipeline.stderr {
		return nil, errAlreadySet
	}
	stderrPipe, pipeline.stderr = io.Pipe()


        return stderrPipe, nil
}









func (pipeline *Pipeline) MustStdinPipe() io.WriteCloser {

	w, err := pipeline.StdinPipe()
	if nil != err {
		panic(err)
	}

	return w
}

func (pipeline *Pipeline) MustStdoutPipe() io.ReadCloser {

	r, err := pipeline.StdoutPipe()
	if nil != err {
		panic(err)
	}

	return r
}

func (pipeline *Pipeline) MustStderrPipe() io.ReadCloser {

	r, err := pipeline.StderrPipe()
	if nil != err {
		panic(err)
	}

	return r
}









// SetStdin sets the Pipeline's standard input.
//
// SetStdin returns an error if it cannot set the standard intput of the Pipeline.
//
// The standard input of a pipeline cannot be set twice, in that SetStdin will return an
// error if SetStdin is called after the standard input has already been set either via
// a call to SetStdin, MustSetStdin, StdinPipe, or MustStdinPipe.
//
// Example usage:
//
//	reader := strings.NewReader("apple banana cherry")
//	
//	err := pipeline.SetStdin(reader)
//	if nil != err {
//		fmt.Printf("ERROR: %v\n", err)
//		os.Exit(1)
//	}
func (pipeline *Pipeline) SetStdin(r io.Reader) error {
	if nil == pipeline {
		return errNilReceiver
	}

	if pipeline.alreadyRun {
		return errAlreadyRun
	}


	executors := pipeline.Executors
	if nil == executors || len(executors) <= 0 {
		return errEmptyPipeline
	}

	return executors[0].SetStdin(r)
}

// SetStdout sets the Pipeline's standard output.
//
// SetStdout returns an error if it cannot set the standard output of the Pipeline.
//
// The standard output of a pipeline cannot be set twice, in that SetStdout will return an
// error if SetStdout is called after the standard output has already been set either via
// a call to SetStdout, MustSetStdout, StdoutPipe, or MustStdoutPipe.
//
// Example usage:
//
//	var stdoutBuffer bytes.Buffer
//	
//	err := pipeline.SetStdout(&stdoutBuffer)
//	if nil != err {
//		fmt.Printf("ERROR: %v\n", err)
//		os.Exit(1)
//	}
func (pipeline *Pipeline) SetStdout(w io.WriteCloser) error {
	if nil == pipeline {
		return errNilReceiver
	}

	if pipeline.alreadyRun {
		return errAlreadyRun
	}


	executors := pipeline.Executors
	if nil == executors || len(executors) <= 0 {
		return errEmptyPipeline
	}

	return executors[ len(executors)-1 ].SetStdout(w)
}

// SetStderr sets the Pipeline's standard error.
//
// SetStderr returns an error if it cannot set the standard error of the Pipeline.
//
// The standard error of a pipeline cannot be set twice, in that SetStderr will return an
// error if SetStderr is called after the standard error has already been set either via
// a call to SetStderr, MustSetStderr, StderrPipe, or MustStderrPipe.
//
// Example usage:
//
//	var stderrBuffer bytes.Buffer
//	
//	err := pipeline.SetStderr(&stderrBuffer)
//	if nil != err {
//		fmt.Printf("ERROR: %v\n", err)
//		os.Exit(1)
//	}
func (pipeline *Pipeline) SetStderr(w io.WriteCloser) error {
	if nil == pipeline {
		return errNilReceiver
	}

	if pipeline.alreadyRun {
		return errAlreadyRun
	}


	executors := pipeline.Executors
	if nil == executors || len(executors) <= 0 {
		return errEmptyPipeline
	}


	if nil != pipeline.stderr {
		return errAlreadySet
	}

	pipeline.stderr = w

	return nil
}









// MustSetStdin is like SetStdin, except it panic()s if it cannot set the standard input of the Pipeline.
//
// Example usage:
//
//	reader := strings.NewReader("apple banana cherry")
//	
//	pipeline.MustSetStdin(reader) // This might panic()!
func (pipeline *Pipeline) MustSetStdin(r io.Reader) {
	if err := pipeline.SetStdin(r); nil != err {
		panic(err)
	}
}

// MustSetStdout is like SetStdin, except it panic()s if it cannot set the standard output of the Pipeline.
//
// Example usage:
//
//	var stdoutBuffer bytes.Buffer
//	
//	pipeline.MustSetStdout(&stdoutBuffer) // This might panic()!
func (pipeline *Pipeline) MustSetStdout(w io.WriteCloser) {
	if err := pipeline.SetStdout(w); nil != err {
		panic(err)
	}
}

// MustSetStderr is like SetStdin, except it panic()s if it cannot set the standard error of the Pipeline.
//
// Example usage:
//
//	var stderrBuffer bytes.Buffer
//	
//	pipeline.MustSetStderr(&stderrBuffer) // This might panic()!
func (pipeline *Pipeline) MustSetStderr(w io.WriteCloser) {
	if err := pipeline.SetStderr(w); nil != err {
		panic(err)
	}
}
