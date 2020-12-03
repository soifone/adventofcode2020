// Day 2
// --- Part One ---
// Your flight departs in a few days from the coastal airport; the easiest way down to the coast from here is via toboggan.
//
// The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day. "Something's wrong with our computers;
// we can't log in!" You ask if you can take a look.
//
// Their password database seems to be a little corrupted: some of the passwords wouldn't have been allowed by the
// Official Toboggan Corporate Policy that was in effect when they were chosen.
//
// To try to debug the problem, they have created a list (your puzzle input) of passwords (according to the corrupted
// database) and the corporate policy when that password was set.
//
// For example, suppose you have the following list:
//
// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc
// Each line gives the password policy and then the password. The password policy indicates the lowest and highest
// number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password
// must contain a at least 1 time and at most 3 times.
//
// In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b,
// but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the
// limits of their respective policies.
//
// How many passwords are valid according to their policies?
//
// link: https://adventofcode.com/2020/day/2

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	BoundryL, BoundryR  int
	Character, Password string
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	validOcc := 0
	validExact := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		p, err := extractPasswordPolicy(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if p.isValidByBoundaryOccurrence() {
			validOcc++
		}

		if p.isValidByExactlyOnePosition() {
			validExact++
		}
	}

	fmt.Printf("%d passwords occurance are valid\n", validOcc)
	fmt.Printf("%d passwords exact are valid\n", validExact)

}

// extractPasswordPolicy takes a string i and will return
// the given string as a password policy type.
//
// Given i
//
//		1-2 z: aaccddzz
//
//	Will return
//
//	{
//		Min: 1,
//		Max: 2,
//		Character: z,
//		Password: aaccddzz
//	}, nil
func extractPasswordPolicy(i string) (*PasswordPolicy, error) {
	l := strings.Replace(i, ":", "", 1)
	p := strings.Split(l, " ")
	boundaries := strings.Split(p[0], "-")

	min, err := strconv.Atoi(boundaries[0])
	if err != nil {
		return nil, fmt.Errorf("min cannot be cast to int.\n%v\nline:%s", err, i)
	}

	max, err := strconv.Atoi(boundaries[1])
	if err != nil {
		return nil, fmt.Errorf("max cannot be cast to int.\n%v\nline:%s", err, i)
	}

	return &PasswordPolicy{
		BoundryL:  min,
		BoundryR:  max,
		Character: p[1],
		Password:  p[2],
	}, nil
}

// isValid will check if the password in isValidByBoundaryOccurrence is valid
// given by the boundaries left, right and the character.
func (p *PasswordPolicy) isValidByBoundaryOccurrence() bool {
	s := strings.Split(p.Password, "")
	sort.Strings(s)

	occurrence := 0
	for _, r := range s {
		if r > p.Character {
			break
		}

		if r == p.Character {
			occurrence++
		}
	}

	return occurrence >= p.BoundryL && occurrence <= p.BoundryR
}

func (p *PasswordPolicy) isValidByExactlyOnePosition() bool {
	s := strings.Split(p.Password, "")
	return (s[p.BoundryL-1] == p.Character) != (s[p.BoundryR-1] == p.Character)
}
