package namesgenerator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameFormat(t *testing.T) {
	name := GetRandomName(0)

	assert.True(t, strings.Contains(name, "_"), "Generated name does not contain an underscore")
	assert.False(t, strings.ContainsAny(name, "0123456789"), "Generated name contains numbers")

}

func TestNameRetries(t *testing.T) {
	name := GetRandomName(1)

	assert.True(t, strings.Contains(name, "_"), "Generated name does not contain an underscore")
	assert.True(t, strings.ContainsAny(name, "0123456789"), "Generated name doesn't contain a number")
	
}

func BenchmarkGetRandomName(b *testing.B) {
	b.ReportAllocs()
	var out string
	for n := 0; n < b.N; n++ {
		out = GetRandomName(5)
	}
	b.Log("Last result:", out)
}
