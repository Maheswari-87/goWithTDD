package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Mahi")
	got := buffer.String()
	want := "Hello, Mahi"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
