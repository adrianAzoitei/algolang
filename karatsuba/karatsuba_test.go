package karatsuba_test

import (
	"algorithms/karatsuba"
	"testing"
)

func TestKaratsuba(t *testing.T) {
	x := 1234
	y := 5678
	want := 7006652
	res, err := karatsuba.Karatsuba(x, y)
	if err != nil {
		t.Fatalf("Karatsuba returned an error: %s", err)
	}
	if res != want {
		t.Fatalf("Expected %d, got %d", want, res)
	}
}
