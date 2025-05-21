package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Milos"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Milos")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Milos") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
