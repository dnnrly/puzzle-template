package puzzle_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dnnrly/puzzle-template"
)

func TestIsPrime(t *testing.T) {
	assert.False(t, puzzle.IsPrime(1))
}
