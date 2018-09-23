package calg

import (
	"testing"
)

func TestColorable(t *testing.T) {
	input := "testString"
	want := "\x1b[31m" + input + "\x1b[0m"
	have := colorable(input)
	if want != have {
		t.Errorf("colorable() is invalid. have:%v want:%v ", have, want)
	}
}
