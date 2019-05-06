package jwt

import "bytes"

// RawToken is a representation of a parsed JWT string.
type RawToken struct {
	token      []byte
	sep1, sep2 int
}

// Parse parses a byte slice representing a JWT and returns a raw JWT,
// which can be verified and decoded into a struct that implements Token.
func Parse(token []byte) (RawToken, error) {
	var t RawToken

	sep1 := bytes.IndexByte(token, '.')
	if sep1 < 0 {
		return t, ErrMalformed
	}
	t.sep1 = sep1

	cbytes := token[sep1+1:]
	sep2 := bytes.IndexByte(cbytes, '.')
	if sep2 < 0 {
		return t, ErrMalformed
	}
	t.sep2 = sep1 + 1 + sep2
	t.token = token
	return t, nil
}

// Decode decodes a raw JWT into a payload and returns its header.
func (r RawToken) Decode(payload interface{}) (Header, error) {
	var h Header
	if err := decode(r.header(), &h); err != nil {
		return h, err
	}
	return h, decode(r.payload(), payload)
}

// Verify verifies a JWT signature with a given Verifier.
func (r RawToken) Verify(vr Verifier) error {
	return vr.Verify(r.headerPayload(), r.sig())
}

func (r RawToken) header() []byte        { return r.token[:r.sep1] }
func (r RawToken) headerPayload() []byte { return r.token[:r.sep2] }
func (r RawToken) payload() []byte       { return r.token[r.sep1+1 : r.sep2] }
func (r RawToken) sig() []byte           { return r.token[r.sep2+1:] }
