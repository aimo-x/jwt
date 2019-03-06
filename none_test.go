package jwt_test

import (
	"testing"

	. "github.com/gbrlsnchs/jwt/v3"
)

func TestNone(t *testing.T) {
	testCases := []testCase{
		{new(None), new(None), nil, nil, nil, nil},
		{new(None), NewHMAC(SHA256, []byte("secret")), nil, nil, nil, ErrHMACVerification},
		{NewHMAC(SHA256, []byte("secret")), new(None), nil, nil, nil, nil},
	}
	testJWT(t, testCases)
}
