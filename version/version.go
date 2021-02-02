package version

var (
	tag    string
	commit string
	author string
	date   string

	buildDate string
	buildUser string
)

// Tag returns the version tag. If semantic versioning is being followed this
// should be of the form: vx.y.z where x, y and z are numbers.
func Tag() string {
	return tag
}

// Commit returns the latest commit ID.
func Commit() string {
	return commit
}

// Author returns the author of the latest commit.
func Author() string {
	return author
}

// Date returns the date of the latest commit.
func Date() string {
	return date
}

// BuildDate returns the date that the program was built.
func BuildDate() string {
	return buildDate
}

// BuildUser returns the user that built the program.
func BuildUser() string {
	return buildUser
}

// All returns a string giving all the version details. It is a formatted
// string with the values shown following names.
func All() string {
	return "" +
		"      Version: " + tag + "\n" +
		"    Commit ID: " + commit + "\n" +
		"Commit Author: " + author + "\n" +
		"  Commit Date: " + date + "\n" +
		"   Build Date: " + buildDate + "\n" +
		"     Built By: " + buildUser
}
