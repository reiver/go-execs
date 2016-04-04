package execs


import (
	"github.com/reiver/go-oi"

	"bytes"

	"testing"
)


func TestPipelineAlreadyRun(t *testing.T) {

	var alreadyRunPipeline *Pipeline = &Pipeline{alreadyRun:true}
	var buffer bytes.Buffer



	if err := alreadyRunPipeline.Run(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errAlreadyRun != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}



	if err := alreadyRunPipeline.SetStderr( oi.WriteNopCloser(&buffer) ); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errAlreadyRun != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if err := alreadyRunPipeline.SetStdin(&buffer); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errAlreadyRun != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if err := alreadyRunPipeline.SetStdout( oi.WriteNopCloser(&buffer) ); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errAlreadyRun != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}



	if _, err := alreadyRunPipeline.StderrPipe(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errAlreadyRun != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if _, err := alreadyRunPipeline.StdinPipe(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errAlreadyRun != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if _, err := alreadyRunPipeline.StdoutPipe(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errAlreadyRun != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}



	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errAlreadyRun != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		alreadyRunPipeline.MustRun()

		t.Errorf("Should never get here!")
	}()



	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errAlreadyRun != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		alreadyRunPipeline.MustSetStderr( oi.WriteNopCloser(&buffer ))

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errAlreadyRun != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		alreadyRunPipeline.MustSetStdin(&buffer)

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errAlreadyRun != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		alreadyRunPipeline.MustSetStdout( oi.WriteNopCloser(&buffer) )

		t.Errorf("Should never get here!")
	}()



	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errAlreadyRun != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		_ = alreadyRunPipeline.MustStderrPipe()

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errAlreadyRun != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		_ = alreadyRunPipeline.MustStdinPipe()

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errAlreadyRun != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		_ = alreadyRunPipeline.MustStdoutPipe()

		t.Errorf("Should never get here!")
	}()
}
