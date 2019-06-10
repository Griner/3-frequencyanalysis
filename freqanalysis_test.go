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
		{"aaa bbb ccc aaa aaa bbb", WordStats{{"aaa", 3}, {"bbb", 2}, {"ccc", 1}}},
		{"aaa bbb aaa aaa", WordStats{{"aaa", 3}, {"bbb", 1}}},
	}

	for _, testCase := range cases {

		out := Freq(testCase.in)
		if !reflect.DeepEqual(out, testCase.out) {
			t.Errorf("EXPECTED %v ; GOT %v\n", testCase.out, out)
		}
	}
}
