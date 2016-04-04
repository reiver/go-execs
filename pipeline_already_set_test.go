package execs


import (
	"github.com/reiver/go-oi"

	"bytes"

	"testing"
)


func TestPipelineSetStderrAlreadySet(t *testing.T) {

	var pipeline *Pipeline = &Pipeline{ Executors:[]Executor{ &NopExecutor{} } }

	var buffer1 bytes.Buffer
	if err := pipeline.SetStderr( oi.WriteNopCloser(&buffer1) ); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	const limit = 10
	for i:=0; i<limit; i++ {
		var buffer2 bytes.Buffer
		if err := pipeline.SetStderr( oi.WriteNopCloser(&buffer2) ); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}
	for i:=0; i<limit; i++ {
		if _, err := pipeline.StderrPipe(); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}


	var buffer3 bytes.Buffer
	if err := pipeline.SetStdin(&buffer3); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	var buffer4 bytes.Buffer
	if err := pipeline.SetStdout( oi.WriteNopCloser(&buffer4) ); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}
}


func TestPipelineSetStdinAlreadySet(t *testing.T) {

	var pipeline *Pipeline = &Pipeline{ Executors:[]Executor{ &NopExecutor{} } }

	var buffer1 bytes.Buffer
	if err := pipeline.SetStdin(&buffer1); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	const limit = 10
	for i:=0; i<limit; i++ {
		var buffer2 bytes.Buffer
		if err := pipeline.SetStdin(&buffer2); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}
	for i:=0; i<limit; i++ {
		if _, err := pipeline.StdinPipe(); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}


	var buffer3 bytes.Buffer
	if err := pipeline.SetStdout( oi.WriteNopCloser(&buffer3) ); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	var buffer4 bytes.Buffer
	if err := pipeline.SetStderr( oi.WriteNopCloser(&buffer4) ); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}
}


func TestPipelineSetStdoutAlreadySet(t *testing.T) {

	var pipeline *Pipeline = &Pipeline{ Executors:[]Executor{ &NopExecutor{} } }

	var buffer1 bytes.Buffer
	if err := pipeline.SetStdout( oi.WriteNopCloser(&buffer1) ); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	const limit = 10
	for i:=0; i<limit; i++ {
		var buffer2 bytes.Buffer
		if err := pipeline.SetStdout( oi.WriteNopCloser(&buffer2) ); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}
	for i:=0; i<limit; i++ {
		if _, err := pipeline.StdoutPipe(); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}


	var buffer3 bytes.Buffer
	if err := pipeline.SetStdin(&buffer3); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	var buffer4 bytes.Buffer
	if err := pipeline.SetStderr( oi.WriteNopCloser(&buffer4) ); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}
}









func TestPipelineStderrPipeAlreadySet(t *testing.T) {

	var pipeline *Pipeline = &Pipeline{ Executors:[]Executor{ &NopExecutor{} } }

	if _, err := pipeline.StderrPipe(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	const limit = 10
	for i:=0; i<limit; i++ {
		if _, err := pipeline.StderrPipe(); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}
	for i:=0; i<limit; i++ {
		var buffer bytes.Buffer
		if err := pipeline.SetStderr( oi.WriteNopCloser(&buffer) ); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}


	if _, err := pipeline.StdinPipe(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if _, err := pipeline.StdoutPipe(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}
}


func TestPipelineStdinPipeAlreadySet(t *testing.T) {

	var pipeline *Pipeline = &Pipeline{ Executors:[]Executor{ &NopExecutor{} } }

	if _, err := pipeline.StdinPipe(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	const limit = 10
	for i:=0; i<limit; i++ {
		if _, err := pipeline.StdinPipe(); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}
	for i:=0; i<limit; i++ {
		var buffer bytes.Buffer
		if err := pipeline.SetStdin(&buffer); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}


	if _, err := pipeline.StdoutPipe(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if _, err := pipeline.StderrPipe(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}
}


func TestPipelineStdoutPipeAlreadySet(t *testing.T) {

	var pipeline *Pipeline = &Pipeline{ Executors:[]Executor{ &NopExecutor{} } }

	if _, err := pipeline.StdoutPipe(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	const limit = 10
	for i:=0; i<limit; i++ {
		if _, err := pipeline.StdoutPipe(); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}
	for i:=0; i<limit; i++ {
		var buffer bytes.Buffer
		if err := pipeline.SetStdout( oi.WriteNopCloser(&buffer) ); nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
	}


	if _, err := pipeline.StdinPipe(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if _, err := pipeline.StderrPipe(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}
}
