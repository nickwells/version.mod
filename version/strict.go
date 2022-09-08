//go:build !version_no_check

package version

import (
	"fmt"
	"os"
)

// This init function will check that at least one of the version strings
// have been set and if none of them have been it will report an error and
// exit. In order to skip these checks and allow a build without any version
// strings being set you should build the program with the build tag
// 'version_no_check' set.
func init() {
	if tag == "" && commit == "" &&
		author == "" && date == "" &&
		buildDate == "" && buildUser == "" {
		fmt.Fprintln(os.Stderr, "None of the version values have been set")
		fmt.Fprint(os.Stderr, "The build-time version values can be set"+
			" as follows:"+
			"\n\n"+
			"go build -ldflags=\"-X '<module>/version.<xxx>=<value>'\"\n")

		os.Exit(1)
	}
}
