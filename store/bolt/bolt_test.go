package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type A struct {
	X []byte
}

func TestPutAndGet(t *testing.T) {

	store := New(".testdummycanbedeleted", "dummydata.db")
	key := []byte("key")
	value := []byte("value")
	bucket := "bck"

	a := A{X: value}
	err := store.Create(key, a, bucket)
	assert.Nil(t, err)

	var b A
	err = store.Get(key, &b, bucket)
	assert.Nil(t, err)
	assert.Equal(t, string(a.X), string(b.X))

}
