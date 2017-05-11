package prm

import (
	"testing"
)

func TestParameterize(t *testing.T) {
	cases := []struct {
		i, e string
	}{
		{"-a-a", "a-a"},
		{"--a-a", "a-a"},
		{"a--a", "a-a"},
		{"a-a--", "a-a"},
		{"a-a-", "a-a"},
		{"Hey Anthony", "hey-anthony"},
		{"It's more fun to compute", "its-more-fun-to-compute"},
		{"It 's more fun to compute", "it-s-more-fun-to-compute"},
		{"--K.G.B.--", "kgb"},
		{"C.I.A. and K.G.B.", "cia-and-kgb"},
	}

	for _, c := range cases {
		g := Parameterize(c.i, '-')
		if g != c.e {
			t.Errorf("'%s' should be parameterizeed to '%s' but got '%s'", c.i, c.e, g)
		}
	}
}

func BenchmarkParameterize(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Parameterize("Beam myself into the future", '-')
	}
}
