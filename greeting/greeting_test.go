package greeting

import (
    "regexp"
    "testing"
)

// https://pkg.go.dev/testing#T
func TestGreetName(t *testing.T) {
    name := "Adam"

    want := regexp.MustCompile(`\b`+name+`\b`)

    msg, err := Greet("Adam")

    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Adam") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

func TestGreetEmpty(t *testing.T) {
    msg, err := Greet("")

    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
