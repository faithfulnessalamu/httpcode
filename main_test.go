package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		desc        string
		args        []string
		shouldErr   bool
		errIfShould error
	}{
		{
			"httpcode -v 200",
			[]string{"-v", "200"},
			false,
			nil,
		},
		{
			"httpcode 200",
			[]string{"200"},
			false,
			nil,
		},
		{
			"httpcode -v",
			[]string{"-v"},
			true,
			errInvalidCode,
		},
		{
			"httpcode",
			[]string{},
			false,
			nil,
		},
		{
			"httpcode -v 1000",
			[]string{"-v", "1000"},
			true,
			errOutOfRange,
		},
		{
			"httpcode 1000 -v",
			[]string{"1000", "-v"},
			true,
			errInvalidCode,
		},
		{
			"httpcode 599",
			[]string{"599"},
			true,
			errNonStandard,
		},
		{
			"httpcode -v 599",
			[]string{"-v", "599"},
			true,
			errNonStandard,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := run(tc.args)
			if err != nil && !tc.shouldErr {
				t.Errorf("got error '%v' for args '%+v'", err, tc.args)
			} else if err == nil && tc.shouldErr {
				t.Errorf("error '%v' expected for args '%+v'", tc.errIfShould, tc.args)
			} else if err != nil && tc.shouldErr {
				if err != tc.errIfShould {
					t.Errorf("expected err '%v', got '%v'", tc.errIfShould, err)
				}
			}
		})
	}
}

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
