package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var tcs = []struct {
	input string
	want  *PasswordPolicy
}{
	{
		"1-4 b: abc",
		&PasswordPolicy{
			BoundryL:  1,
			BoundryR:  4,
			Character: "b",
			Password:  "abc",
		},
	},
	{
		"2-4 c: asbasdccccccceafffff",
		&PasswordPolicy{
			BoundryL:  2,
			BoundryR:  4,
			Character: "c",
			Password:  "asbasdccccccceafffff",
		},
	},
}

func TestExtractPasswordPolicy(t *testing.T) {
	for _, tc := range tcs {
		p, _ := extractPasswordPolicy(tc.input)

		if diff := cmp.Diff(p, tc.want); diff != "" {
			t.Errorf("\ndifferences in PasswordPolicy found %s\n", diff)
		}
	}
}

func TestCheckValidByBoundaryOccurrence(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{
			"1-3 a: abcde",
			true,
		}, {
			"1-3 b: cdefg",
			false,
		}, {
			"2-9 c: ccccccccc",
			true,
		},
	} {
		p, _ := extractPasswordPolicy(tc.input)
		got := p.isValidByBoundaryOccurrence()
		if got != tc.want {
			t.Errorf("\nGot: %v\nWant: %v\nInput: %s", got, tc.want, tc.input)
		}
	}
}

func TestCheckValidByExactlyOneOccurrence(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{
			"1-3 a: abcde",
			true,
		}, {
			"1-3 b: bdbfg",
			false,
		}, {
			"2-9 c: ccccccccc",
			false,
		},
	} {
		p, _ := extractPasswordPolicy(tc.input)
		got := p.isValidByExactlyOnePosition()
		if got != tc.want {
			t.Errorf("\nGot: %v\nWant: %v\nInput: %s", got, tc.want, tc.input)
		}
	}
}
