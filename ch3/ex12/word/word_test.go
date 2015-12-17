package word

import (
	"testing"
)

// -- Tests --

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		s    string
		t    string
		want bool
	}{
		{"", "", true},
		{"abcd", "cbda", true},
		{"a b.c4d", "cbda", true},
		{"a B.c4dda", "CDbada", true},
		{"ab", "abb", false},
	}
	for _, test := range tests {
		if got := IsAnagram(test.s, test.t); got != test.want {
			t.Errorf("IsAnagram(%q, %q) = %v", test.s, test.t, got)
		}
	}
}
