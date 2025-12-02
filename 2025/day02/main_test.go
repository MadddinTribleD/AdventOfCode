package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	invalid := []int64{
		55,
		6464,
		123123,
	}

	for _, i := range invalid {
		t.Run(fmt.Sprintf("invalid_%d", i), func(t *testing.T) {
			if isValid(i) {
				t.Fail()
			}
		})
	}

	valid := []int64{
		101,
		222220,
		5,
		1011,
	}

	for _, i := range valid {
		t.Run(fmt.Sprintf("valid_%d", i), func(t *testing.T) {
			if !isValid(i) {
				t.Fail()
			}
		})
	}
}
