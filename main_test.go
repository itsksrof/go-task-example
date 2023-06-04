package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func Test_out(t *testing.T) {
	file, err := os.CreateTemp("./", "")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	str := "foo"
	buf := bytes.NewBufferString(str)
	if _, err := file.Write(buf.Bytes()); err != nil {
		t.Error(err)
	}

	cmd := commands()["-o"]
	out := cmd(file.Name())

	if !strings.Contains(out, str) {
		t.Error("strings are not equal")
	}

	if err := os.Remove(file.Name()); err != nil {
		t.Error(err)
	}
}
