package day1

import "testing"

func TestDay1(t *testing.T) {
	type test struct {
		filename string
		reader   func(string) ([]string, error)
		want     int
	}

	tests := []test{
		{
			filename: "fake_file",
			reader: func(s string) ([]string, error) {
				return []string{
					"two1nine",
					"eightwothree",
				}, nil
			},
			want: 112,
		},
	}

	for _, uc := range tests {
		got := Day1(uc.filename, uc.reader)
		if got != uc.want {
			t.Fatalf("expected: %v, got: %v", uc.want, got)
		}
	}
}