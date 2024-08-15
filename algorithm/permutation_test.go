package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	// empty 
	assert.Equal(t, []string{""}, Permutation("")) 
	// single 
	assert.Equal(t, []string{"a"}, Permutation("a")) 
	// multiple	
	assert.Equal(t, []string{"abc", "acb", "bac", "bca", "cab", "cba"}, Permutation("abc"))
	// duplicate 
	assert.Equal(t, []string{"aab", "aba", "baa"}, Permutation("aab")) 
}