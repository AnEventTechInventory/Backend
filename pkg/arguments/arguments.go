package arguments

import (
	"flag"
)

var Args args

type args struct {
	Verbose bool
}

func ParseArgs() {
	// Define a boolean flag named "verbose" with a default value of false,
	// a brief description, and a shorthand "-v".
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output (shorthand)")

	// Parse the command-line arguments.
	flag.Parse()

	Args = args{
		Verbose: verbose,
	}
}
