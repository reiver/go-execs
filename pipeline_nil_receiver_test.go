package execs


import (
	"github.com/reiver/go-oi"

	"bytes"

	"testing"
)


func TestPipelineNilReceiver(t *testing.T) {

	var nilPipeline *Pipeline = nil
	var buffer bytes.Buffer



	if err := nilPipeline.Run(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errNilReceiver != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}



	if err := nilPipeline.SetStderr( oi.WriteNopCloser(&buffer) ); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errNilReceiver != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if err := nilPipeline.SetStdin(&buffer); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errNilReceiver != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if err := nilPipeline.SetStdout( oi.WriteNopCloser(&buffer) ); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errNilReceiver != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}



	if _, err := nilPipeline.StderrPipe(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errNilReceiver != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if _, err := nilPipeline.StdinPipe(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errNilReceiver != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}

	if _, err := nilPipeline.StdoutPipe(); nil == err {
		t.Error("Expected an error, but did not actually get one.")
	} else if errNilReceiver != err {
		t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
	}



	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errNilReceiver != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		nilPipeline.MustRun()

		t.Errorf("Should never get here!")
	}()



	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errNilReceiver != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		nilPipeline.MustSetStderr( oi.WriteNopCloser(&buffer) )

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errNilReceiver != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		nilPipeline.MustSetStdin(&buffer)

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errNilReceiver != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		nilPipeline.MustSetStdout( oi.WriteNopCloser(&buffer) )

		t.Errorf("Should never get here!")
	}()



	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errNilReceiver != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		_ = nilPipeline.MustStderrPipe()

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errNilReceiver != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		_ = nilPipeline.MustStdinPipe()

		t.Errorf("Should never get here!")
	}()

	func(){
		defer func(){
			if r := recover(); nil != r {
				if err, ok := r.(error); !ok {
					t.Error("Expected an error, but did not actually get one.")
				} else if errNilReceiver != r {
					t.Errorf("Expected a different error; actually got: (%T) %v", err, err)
				}
			}
		}()

		_ = nilPipeline.MustStdoutPipe()

		t.Errorf("Should never get here!")
	}()
}
