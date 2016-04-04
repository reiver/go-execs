package execs


import (
	"github.com/reiver/go-oi"

	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"testing"
)


func TestPipelineRun(t *testing.T) {

	tests := []struct{
		Executors []Executor
		Stdin string
		ExpectedStdout string
		ExpectedStderr string
	}{
		{
			Executors: []Executor{
				&Inline{
					Func: func(stdin io.Reader, stdout io.WriteCloser, stderr io.WriteCloser) error {
						if nil == stdin {
							return errors.New("stdin is nil")
						}

						if nil == stdout {
							return errors.New("stdout is nil")
						}

						buffer, err := ioutil.ReadAll(stdin)
						if nil != err {
							return err
						}

						buffer = bytes.ToUpper(buffer)

						oi.LongWrite(stdout, buffer)

						return nil
					},
				},
			},
			Stdin:          "Apple\nBanana\nCherry\n",
			ExpectedStdout: "APPLE\nBANANA\nCHERRY\n",
			ExpectedStderr: "",
		},
		{
			Executors: []Executor{
				&Inline{
					Func: func(stdin io.Reader, stdout io.WriteCloser, stderr io.WriteCloser) error {
						if nil == stdin {
							return errors.New("stdin is nil")
						}

						if nil == stdout {
							return errors.New("stdout is nil")
						}

						buffer, err := ioutil.ReadAll(stdin)
						if nil != err {
							return err
						}

						buffer = bytes.ToLower(buffer)

						oi.LongWrite(stdout, buffer)

						return nil
					},
				},
			},
			Stdin:          "Apple\nBanana\nCherry\n",
			ExpectedStdout: "apple\nbanana\ncherry\n",
			ExpectedStderr: "",
		},



		{
			Executors: []Executor{
				&Inline{
					Func: func(stdin io.Reader, stdout io.WriteCloser, stderr io.WriteCloser) error {
						if nil == stdin {
							return errors.New("stdin is nil")
						}

						if nil == stdout {
							return errors.New("stdout is nil")
						}

						if nil == stderr {
							return errors.New("stderr is nil")
						}

						buffer, err := ioutil.ReadAll(stdin)
						if nil != err {
							return err
						}

						buffer[2] = '2'

						oi.LongWrite(stdout, buffer)

						return nil
					},
				},
				&Inline{
					Func: func(stdin io.Reader, stdout io.WriteCloser, stderr io.WriteCloser) error {
						if nil == stdin {
							return errors.New("stdin is nil")
						}

						if nil == stdout {
							return errors.New("stdout is nil")
						}

						if nil == stderr {
							return errors.New("stderr is nil")
						}

						buffer, err := ioutil.ReadAll(stdin)
						if nil != err {
							return err
						}


						buffer[7] = '7'

						oi.LongWrite(stdout, buffer)

						return nil
					},
				},
				&Inline{
					Func: func(stdin io.Reader, stdout io.WriteCloser, stderr io.WriteCloser) error {
						if nil == stdin {
							return errors.New("stdin is nil")
						}

						if nil == stdout {
							return errors.New("stdout is nil")
						}

						if nil == stderr {
							return errors.New("stderr is nil")
						}

						buffer, err := ioutil.ReadAll(stdin)
						if nil != err {
							return err
						}


						buffer[7] = '7'

						oi.LongWrite(stdout, buffer)

						return nil
					},
				},
			},
			Stdin:          "Apple\nBanana\nCherry\n",
			ExpectedStdout: "Ap2le\nB7nana\nCherry\n",
			ExpectedStderr: "",
		},
	}


	for testNumber, test := range tests {

		pipeline := Pipeline{ Executors: test.Executors }

		pipeline.SetStdin( strings.NewReader(test.Stdin) )

		var stdoutBuffer bytes.Buffer
		pipeline.SetStdout( oi.WriteNopCloser(&stdoutBuffer) )

		var stderrBuffer bytes.Buffer
		pipeline.SetStderr( oi.WriteNopCloser(&stderrBuffer) )

		if err := pipeline.Run(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v", testNumber, err)
			continue
		}

		if expected, actual := test.ExpectedStdout, stdoutBuffer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}

		if expected, actual := test.ExpectedStderr, stderrBuffer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
