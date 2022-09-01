package notify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Notify_SetDebug(t *testing.T) {
	SetDebug(true)
	assert.True(t, isDebug)
	SetDebug(false)
	assert.False(t, isDebug)
}
