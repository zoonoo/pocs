package popcount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	ast := assert.New(t)
	for i, v := range pc {
		ast.Equal(int(v), Count(int64(i)), "parameter : %v", i)
	}
}

func TestLoop(t *testing.T) {
	ast := assert.New(t)
	for i, v := range pc {
		ast.Equal(int(v), CountLoop(int64(i)), "parameter : %v", i)
	}
}

func TestBruteForce(t *testing.T) {
	ast := assert.New(t)
	for i, v := range pc {
		ast.Equal(int(v), CountBruteForce(int64(i)), "parameter : %v", i)
	}
}
