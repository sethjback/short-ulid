package shortulid

import (
	"crypto/rand"
	"errors"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

/*
SULID is a 64 bit unique lexographically sortable identifier.

The first 48 bits are a unix timestamp stored with the MSB first.

The last 16 bit are randomly generated data.
*/
type SULID [8]byte

var (
	ErrDecodeSize = errors.New("decoded ID not correct size")
)

// New returns a SULID for the current timestamp and random entropy
// or error if there is an issue generating the random bits
func New() (*SULID, error) {
	id := &SULID{}

	id.SetTime(time.Now().UTC().UnixMilli())

	b := make([]byte, 2)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	id.SetEntropy([2]byte{b[0], b[1]})

	return id, nil
}

// Must returns a SULID for the current timestamp and random entropy. Will
// panic if there is a problem generating the random bits
func Must() *SULID {
	id := &SULID{}

	id.SetTime(time.Now().UTC().UnixMilli())
	b := make([]byte, 2)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	id.SetEntropy([2]byte{b[0], b[1]})

	return nil
}

// Set the the unix timestamp for the SULID
func (id *SULID) SetTime(ts int64) {
	id[0] = byte(ts >> 40)
	id[1] = byte(ts >> 32)
	id[2] = byte(ts >> 24)
	id[3] = byte(ts >> 16)
	id[4] = byte(ts >> 8)
	id[5] = byte(ts)
}

// Set the entropy for the SULID
func (id *SULID) SetEntropy(b [2]byte) {
	id[6] = b[0]
	id[7] = b[1]
}

// Converts the SULID bytes to a 10 character base58 encoded string
func (id *SULID) String() string {
	return base58.Encode(id[:])
}

// Returns the unixtimestamp for the SULID
func (id *SULID) Timestamp() int64 {
	return int64(id[5]) | int64(id[4])<<8 |
		int64(id[3])<<16 | int64(id[2])<<24 |
		int64(id[1])<<32 | int64(id[0])<<40
}

// Returns a time.Time for the SULID
func (id *SULID) Time() time.Time {
	return time.UnixMilli(int64(id.Timestamp()))
}

// Decode a SULID string and return a SULID
func Decode(sulid string) (*SULID, error) {
	b := base58.Decode(sulid)
	if len(b) != 8 {
		return nil, ErrDecodeSize
	}

	return (*SULID)(b), nil
}
