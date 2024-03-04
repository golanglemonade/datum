package ulids

import (
	"crypto/rand"
	"errors"
	"io"
	"time"

	"github.com/oklog/ulid/v2"
)

// The source of entropy used for all ULIDs generated by this package. This variable is
// initialized when the package is first imported. This package ensures that a
// cryptographic random number generator is used with a locked montonic source so that
// ulids are always increasing in a thread-safe manner
var entropy io.Reader

func init() {
	entropy = &ulid.LockedMonotonicReader{
		MonotonicReader: ulid.Monotonic(rand.Reader, 0),
	}
}

// Null ULID pre-allocated for easier checking.
var Null = ulid.ULID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

var (
	ErrUnknownType = errors.New("cannot parse input: unknown type")
)

// IsZero determines if the specified uid is the Null or zero-valued ULID. Useful for
// determining if a ULID has been passed into a method or is valid
func IsZero(uid ulid.ULID) bool {
	return uid.Compare(Null) == 0
}

// New creates a new montonically increasing ULID in a threadsafe manner using a
// cryptographically random source of entropy for security (e.g. to ensure that an
// attacker cannot guess the next ULID to be generated). If the ULID cannot be generated
// this method panics rather than returning an error
func New() ulid.ULID {
	ms := ulid.Timestamp(time.Now())
	uid, err := ulid.New(ms, entropy)

	if err != nil {
		panic(err)
	}

	return uid
}

// FromTime creates a ULID with the specified timestamp
func FromTime(ts time.Time) ulid.ULID {
	ms := ulid.Timestamp(ts)
	uid, err := ulid.New(ms, entropy)

	if err != nil {
		panic(err)
	}

	return uid
}

// Parse a ULID from a string or a []byte (or return a ulid.ULID). This method makes it
// easier to convert any user-specified type into a ULID
func Parse(uid any) (ulid.ULID, error) {
	switch t := uid.(type) {
	case ulid.ULID:
		return t, nil
	case string:
		if t == "" {
			return Null, nil
		}

		return ulid.Parse(t)

	case []byte:
		var id ulid.ULID
		if err := id.UnmarshalBinary(t); err != nil {
			return Null, err
		}

		return id, nil

	case [16]byte:
		return ulid.ULID(t), nil
	default:
		return Null, ErrUnknownType
	}
}

// MustParse parses the ULID but panics on errors
func MustParse(uid any) (id ulid.ULID) {
	var err error
	if id, err = Parse(uid); err != nil {
		panic(err)
	}

	return id
}

// Bytes parses a ULID and returns the []byte representation of the ULID
func Bytes(uid any) ([]byte, error) {
	sid, err := Parse(uid)
	if err != nil {
		return nil, err
	}

	return sid.Bytes(), nil
}

// MustBytes parses the ULID into bytes but panics on errors
func MustBytes(uid any) []byte {
	id, err := Bytes(uid)
	if err != nil {
		panic(err)
	}

	return id
}