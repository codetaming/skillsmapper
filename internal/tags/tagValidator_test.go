package tags

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tags := []string{"java"}
	assert.Equal(t, tags, Validate(tags))
}
