package util

import (
	"regexp"
)

var PasswordRegex *regexp.Regexp
var UserNameRegex *regexp.Regexp
var EmailRegex *regexp.Regexp

func init() {
	UserNameRegex, _ = regexp.Compile("^([a-zA-Z0-9_-.]{3,50})$")
	PasswordRegex, _ = regexp.Compile("^([a-zA-Z0-9@*#_-.]{8,15})$")
	EmailRegex, _ = regexp.Compile(`^[A-Za-z0-9](([_\.\-]?[a-zA-Z0-9]+)*)@([A-Za-z0-9]+)(([\.\-]?[a-zA-Z0-9]+)*)\.([A-Za-z]{2,})$`)
}
