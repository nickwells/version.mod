/*
version holds various strings describing the current version of a
program. The intention is that the values are set through the -ldflags
argument to the go build and go install tools as follows:

  PKG=github.com/nickwells/version.mod/version
  go build -ldflags="-X '$PKG.tag=v1.2.3' ..."

The values that can be set are:

  tag
  commit
  author
  date
  buildUser
  buildDate

Each of these values is a string.

A suggested way of setting each of these values when using git on a Linux
platform is as follows:

first set up a shell variable to hold the date format to avoid retyping

  DATEFMT="%Y/%m/%d %H:%M:%S"

  tag=$(git --no-pager tag --contains=HEAD)
  commit=$(git rev-parse HEAD)
  author=$(git --no-pager log -1 --format='%an')
  date=$(git --no-pager log -1 --date="format:$DATEFMT" --format='%cd')
  buildUser=$(id --user --real --name)
  buildDate=$(date "+$DATEFMT")

Add these values after the package name as shown above with the -ldflags
parameter to go build or go install.

There is a shell script available (see _sh/goBuildLdflags) which will
construct these arguments for you.

If you wish to build the program without the checks that the build values
have been set then build the program with the build tag version_no_check set
as follows:

  go build -tags version_no_check

Without this build tag a program built without any version information having
been set will panic at startup.

Deprecated: this package superseded by the information in runtime/debug. It
is error-prone, depending as it does on the need for complex build-time
parameters to be set up. It is no longer needed as some of the features it
offered are now provided automatically by the runtime through the
runtime/debug.BuildInfo type. The go version command also provides access to
the same information for binaries compiled with Go versions since go1.18
*/
package version
