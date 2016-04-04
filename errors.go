package execs


import (
	"bytes"
)


// Errors is an error that represents many errors.
//
// Since it is an error, it has an Error method.
//
// But to get access to the many errors it contains, use the Errs method.
//
// Errors might be the error returned from the Pipeline type's Run method.
//
// Since a single Pipeline could contain many type Executors, and when the Run method
// on a single Pipeline is called it in turn calls the Run methods on all its Executors,
// there could (potentially) be many errors returned (from the individual Run methods of
// the Executors the single Pipeline contains).
//
// An example usage might be:
//
//	var pipeline execs.Pipeline
//	
//	pipeline.Executors = []execs.Executor{ex1, ex2, ex3, ex4}
//	
//	err := pipeline.Run()
//	if nil != err {
//		switch e := err.(type) {
//		case execs.Error:
//			fmt.Fprint(os.Stderr, "ERROR: Received many errors!...\n")
//			for errorNumber, subErr := range e.Errs() {
//				fmt.Fprintf(os.Stderr, "\tERROR #%d: %v\n", errorNumber, subErr)
//			}
//			os.Exit(1)
//		default:
//			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
//			os.Exit(1)
//		}
//	}
type Errors struct {
	errors []error
}


func (errs Errors) Error() string {
	es := errs.errors


	if nil == es || 0 >= len(es) {
		return "No errors!"
	}


	var buffer bytes.Buffer

	buffer.WriteString("Many Errors:")

	for i, e := range es {
		if 0 != i {
			buffer.WriteByte(';')
		}
		buffer.WriteByte(' ')
		buffer.WriteString( e.Error() )
	}


	return buffer.String()
}


func (errs Errors) Errs() []error {
	return errs.errors
}
