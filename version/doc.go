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
*/
package version
