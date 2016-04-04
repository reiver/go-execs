package execs


import (
	"github.com/reiver/go-oi"

	"bytes"

	"testing"
)


func TestPipelineEmptyPipeline(t *testing.T) {

	// Run needs to use a separate `emptyPipeline` variable, so the that the rest of the tests
	// are not affected by the pipeline changing into an "already run" state.
	func() {
		var emptyPipeline *Pipeline = &Pipeline{}

		if err := emptyPipeline.Run(); nil == err {
			t.Error("Expected an error, but did not actually get one.")
		} else if errEmptyPipeline != err {
			t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
		}
	}()

	// MustRun needs to use a separate `emptyPipeline` variable, so the that the rest of the tests
	// are not affected by the pipeline changing into an "already run" state.
	func(){
		var emptyPipeline *Pipeline = &Pipeline{}

		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errEmptyPipeline != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		emptyPipeline.MustRun()

		t.Errorf("Should never get here!")
	}()



	var emptyPipeline *Pipeline = &Pipeline{}
	var buffer bytes.Buffer



	if err := emptyPipeline.SetStderr( oi.WriteNopCloser(&buffer) ); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errEmptyPipeline != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if err := emptyPipeline.SetStdin(&buffer); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errEmptyPipeline != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if err := emptyPipeline.SetStdout( oi.WriteNopCloser(&buffer) ); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errEmptyPipeline != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}



	if _, err := emptyPipeline.StderrPipe(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errEmptyPipeline != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if _, err := emptyPipeline.StdinPipe(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errEmptyPipeline != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if _, err := emptyPipeline.StdoutPipe(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errEmptyPipeline != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}



	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errEmptyPipeline != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		emptyPipeline.MustSetStderr( oi.WriteNopCloser(&buffer) )

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errEmptyPipeline != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		emptyPipeline.MustSetStdin(&buffer)

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errEmptyPipeline != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		emptyPipeline.MustSetStdout( oi.WriteNopCloser(&buffer) )

		t.Errorf("Should never get here!")
	}()



	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errEmptyPipeline != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		_ = emptyPipeline.MustStderrPipe()

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errEmptyPipeline != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		_ = emptyPipeline.MustStdinPipe()

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errEmptyPipeline != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		_ = emptyPipeline.MustStdoutPipe()

		t.Errorf("Should never get here!")
	}()
}
