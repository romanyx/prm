package prm

import (
	"testing"
)

func TestSetStack(t *testing.T) {
	stack := []SkipFunc{
		SkipIfFirst,
		SkipIfPrev,
		SkipIfApos,
	}
	SetStack(stack)

	if Parameterize("C.I.A. and K.G.B.", '-') != "c-i-a-and-k-g-b" {
		t.Errorf("should replace stack'")
	}
}
