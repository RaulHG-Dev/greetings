package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Carlos"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, error := Hello("Carlos")

	if !want.MatchString(msg) || error != nil {
		t.Fatalf(`Hello("Carlos") = %q, %v, quiere coincidencia para %#q, nil`, msg, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, error := Hello("")
	if msg != "" || error == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, error)
	}
}
