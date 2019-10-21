package main

import (
	"testing"
)

func TestHahuTesting(t *testing.T) {
	if Add(1, 1) != 2 {
		t.Errorf("not good")
	}
}
