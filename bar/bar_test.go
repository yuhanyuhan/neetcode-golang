package bar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaz(t *testing.T) {

	bar := Bar()
	assert := assert.New(t)

	assert.Contains(bar, "Bar")

}
