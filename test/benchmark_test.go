package test

import (
	"crypto/rand"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//BenchmarkRandInt ...
func BenchmarkRandInt(b *testing.B) {
	r := strings.NewReader("Hello, Reader!")
	num := big.NewInt(100)
	for i := 0; i < b.N; i++ {
		_, _ = rand.Int(r, num)
	}
}

//TestTimeConsuming ...
func TestTimeConsuming(t *testing.T) {
	assert := assert.New(t)
	assert.True(true)
	// if testing.Short() {
	// 	t.Skip("skipping test in short mode.")
	// }
}
