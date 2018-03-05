package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestRun_Version(t *testing.T) {
	mockWriter := new(bytes.Buffer)
	os.Args = []string{"command", "--version"}
	fmt.Println(os.Args[1:])

	want := 0
	got := realMain(mockWriter)
	if want != got {
		t.Fatalf("bad return value \nwant %d \ngot  %d", want, got)
	}

	outWant := "0.1.0\n"
	outGot := mockWriter.String()
	if outWant != outGot {
		t.Fatalf("bad stdout \nwant %s \ngot  %s", outWant, outGot)
	}
}
