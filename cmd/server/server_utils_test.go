package server

import (
	"math"
	"testing"
)

func TestGenerateSecureRandInt64(t *testing.T) {

	got, _ := generateSecureRandInt64()

	if got > math.MaxInt64 {
		t.Errorf("number out of range returned")
	}

	if got < math.MinInt64 {
		t.Errorf("number out of range returned")
	}

}
