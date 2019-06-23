package freqanalysis

import (
	"reflect"
	"testing"
)

func TestFreq(t *testing.T) {

	cases := []struct {
		in  string
		out WordStats
	}{
		{"aaa1  \\  bbb  \n \n \t ccc\naaa1\naaa1 bbb", WordStats{{"aaa1", 3}, {"bbb", 2}, {"ccc", 1}}},
		{"aaa bbb ccc aaa aaa bbb", WordStats{{"aaa", 3}, {"bbb", 2}, {"ccc", 1}}},
		{"aaa bbb aaa aaa", WordStats{{"aaa", 3}, {"bbb", 1}}},
		{"aaa bbb aaa Aaa Ccc ccc", WordStats{{"aaa", 3}, {"ccc", 2}, {"bbb", 1}}},
		{"A a a a A a a A b b c c C c c C c C c c D d d d D e f F f F F F F g G G g G G G g G H h h H h h h h h H H i i I J j J j K K K k k k L l l l l l l l L L l L",
			WordStats{
				{"l", 12},
				{"h", 11},
				{"c", 10},
				{"g", 9},
				{"a", 8},
				{"f", 7},
				{"k", 6},
				{"d", 5},
				{"j", 4},
				{"i", 3},
				// {"b", 2},
				// {"e", 1},
			}},
	}

	for _, testCase := range cases {

		out := Freq(testCase.in)
		if !reflect.DeepEqual(out, testCase.out) {
			t.Errorf("EXPECTED %v ; GOT %v\n", testCase.out, out)
		}
	}
}
