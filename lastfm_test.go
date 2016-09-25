package main

import "testing"

func TestGetData(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	GetData()
}
