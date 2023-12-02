package day2

import (
	"fmt"
	"testing"
)

func TestDay1(t *testing.T) {
	type Test struct {
		filename string
		reader   func(string) ([]string, error)
		want     int
	}

	testCases := []Test{
		{
			filename: "example_part1",
			reader: func(string) ([]string, error) {
				return []string{
					"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
					"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
					"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
					"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
					"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
				}, nil
			},
			want: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testing '%v'", tc.filename), func(t *testing.T) {
			got := Day2(tc.filename, tc.reader)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}

}
