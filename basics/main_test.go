package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_print(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go print("episilon", &wg)
	wg.Wait()

	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "episilon") {
		t.Error("Expected episilon, got ", output)
	}

}
