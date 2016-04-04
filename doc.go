/*
Package execs provides pipelining capabilities to both internal and external commands,
as well a convenient means of inlining internal commands.

Package execs is meant to work with the "os/exec" package, that is part of the Go standard library.

Example Usage

	
	ex1 := cmd.Command("echo", "hello")
	
	ex2 := &execs.Inline{ Func: func1 }
	
	ex3 := cmd.Command("grep", "city")
	
	ex4 := &execs.Inline{ Func: func2 }
	
	pipeline := execs.Pipeline{ Executors:[]Executor{ex1, ex2, ex3, ex4} }
	
	if err := pipeline.Run(); nil != err {
		fmt.Printf("Received an error: %v\n", err)
		os.Exit(1)
	}

*/
package execs
