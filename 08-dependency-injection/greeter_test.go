package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Guilherme")

	got := buffer.String()
	want := "Hello, Guilherme"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
