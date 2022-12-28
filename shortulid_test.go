package shortulid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ts  = int64(1672245117110)
	ent = [2]byte{0xFF, 0xFF}
	enc = "FkmXjJEAq4"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	id, err := New()
	assert.Nil(err)

	did, err := Decode(id.String())
	assert.Nil(err)
	assert.Equal(did, id)

	assert.Equal(did.String(), id.String())
}

func TestDecode(t *testing.T) {
	assert := assert.New(t)

	id := &SULID{}
	id.SetTime(ts)
	id.SetEntropy(ent)

	assert.Equal(enc, id.String())

	did, err := Decode(enc)
	if !assert.Nil(err) {
		return
	}

	assert.Equal(ts, did.Timestamp())
	assert.Equal(did, id)
}
