package arrayhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	assert := assert.New(t)

	pass_test := containsDuplicate([]int{1, 2, 3, 1})
	fail_test := containsDuplicate([]int{1, 2, 3})
	assert.Equal(pass_test, false)
	assert.Equal(fail_test, false)
}
