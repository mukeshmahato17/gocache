package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestParseSetCommand(t *testing.T) {
	cmd := &CommandSet{
		Key:   []byte("Foo"),
		Value: []byte("Bar"),
		TTL:   2,
	}
	fmt.Println(cmd.Bytes())

	r := bytes.NewReader(cmd.Bytes())

	ParseCommand(r)
}
