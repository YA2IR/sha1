package sha1

import (
	"fmt"
	"strings"
	"testing"
)

/*
 *  test vectors taken from https://www.di-mgt.com.au/sha_testvectors.html
 **/

func TestSHA1(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"",
			"da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		{
			"abc",
			"a9993e364706816aba3e25717850c26c9cd0d89d",
		},
		{
			"abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
			"84983e441c3bd26ebaae4aa1f95129e5e54670f1",
		},
		{
			"abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
			"a49b2446a02c645bf419f995b67091253a04a259",
		},
		{
			strings.Repeat("a", 1000*1000),
			"34aa973cd4c4daa4f61eeb2bdbad27316534016f",
		},
		// there's one more test that'll require ~1GB of RAM. Not included.
	}

	sha1 := NewSHA1()
	for i, test := range tests {
		result := sha1.Hash([]byte(test.input))
		if fmt.Sprintf("%x", result) != test.expected {
			t.Errorf("Test %d failed:\nInput: %s\nExpected: %s\nGot: %x",
				i, test.input, test.expected, result)
		}
	}
}
