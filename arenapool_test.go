package bufferpool

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPool_Get(t *testing.T) {
	p := NewPool[*bytes.Buffer](func() *bytes.Buffer {
		return new(bytes.Buffer)
	})
	buf := p.Get()
	assert.NotNil(t, buf)
}

func TestPool_GetWithLargeAlloc(t *testing.T) {
	p := NewPool[[]byte](func() []byte {
		return make([]byte, 1e7)
	})
	buf := p.Get()
	assert.NotNil(t, buf)
}
