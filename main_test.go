package main

import (
	"os"
	"testing"
)

func TestFetchHost(t *testing.T) {
	os.Args = []string{"cmd", "https://gopl.io"}
	valor := os.Args[1:]

	got := FetchHost(&valor)
	if got != nil {
		t.Errorf("Abs(https://gopl.io) = %d; want nil", got)
	}
}