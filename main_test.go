package main

import (
	"testing"
)

func TestIsInRange(t *testing.T) {
	testCases := []struct {
		code    int
		inRange bool
	}{
		{100, true},
		{99, false},
		{599, true},
		{300, true},
		{1, false},
		{-400, false},
		{1000, false},
		{201, true},
		{588, true},
	}

	for _, tc := range testCases {
		if tc.inRange != isInRange(tc.code) {
			t.Errorf("isInRange returns %v for %d", !tc.inRange, tc.code)
		}
	}
}
