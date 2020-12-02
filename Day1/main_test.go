package main

import "testing"

func TestCheckForSumPairAndMultiply(t *testing.T) {
	for _, tc := range []struct{
		list []int
		sum, want int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			10,
			24,
		},
		{
			[]int{5, 2, 12, 33, 10, 6},
			45,
			396,
		},
	}{
		got, err := checkForSumPairAndMultiply(tc.list, tc.sum)
		if err != nil {
			t.Fatal(err)
		}

		if got != tc.want {
			t.Errorf("\nGot: %d\nWant: %d", got, tc.want)
		}
	}
}

func TestCheckForTripletSumMultiply(t *testing.T) {
	for _, tc := range []struct{
		list []int
		sum, want int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			6,
			6,
		},
		{
			[]int{5, 2, 12, 33, 10, 6},
			19,
			120,
		},
	}{
		got, err := checkForTripletsSumAndMultiply(tc.list, tc.sum)
		if err != nil {
			t.Fatal(err)
		}

		if got != tc.want {
			t.Errorf("\nGot: %d\nWant: %d", got, tc.want)
		}
	}
}
