// +build version_strict

package version

import (
	"fmt"
	"os"
)

// advice prints advice on how to set the version values and then exits
func advice() {
	fmt.Fprint(os.Stderr,
		"The build-time version values can be set as follows:"+
			"\n\n"+
			"go build -ldflags=\"-X '<module>/version.<xxx>=<value>'\"\n")
	os.Exit(1)
}

func init() {
	if tag == "" &&
		commit == "" &&
		author == "" &&
		date == "" &&
		buildDate == "" &&
		buildUser == "" {
		fmt.Fprintln(os.Stderr, "None of the version values have been set")
		advice()
	}
}
