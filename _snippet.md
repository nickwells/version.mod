## Version information

This uses the `version` package which provides version information for the
program. It is important that the program is built with either the arguments

`-tags version_no_check`

which prevents the startup check that the version information has been set,
or else with the ldflags set; the script

`github.com/nickwells/version.mod/_sh/goBuildLdflags`

will print a string setting the ldflags appropriately. It can be passed as
the value of the `-ldflags` parameter to `go build` or `go install`.
