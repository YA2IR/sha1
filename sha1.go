package sha1

import (
	"encoding/binary"
)

var k = []uint32{
	0x5A827999, // t : 0-19
	0x6ED9EBA1, // t : 20-39
	0x8F1BBCDC, // t : 40-59
	0xCA62C1D6, // t : 60-79
}

type SHA1 struct {
	state [5]uint32
}

func NewSHA1() *SHA1 {
	sha := &SHA1{}
	// state is (re)-initilaized with each Hash() call
	return sha
}

/*
 *   takes n leftmost bits and puts them in rightmost side of the word.
 **/
func rotateLeft(X uint32, n int) uint32 {
	if n < 0 || n > 32 {
		panic("shiftleft n is out of bounds")
	}
	return (X << n) | (X >> (32 - n))
}

/*
 * mandatory padding, for more info check FIPS 180-1
 **/
func pad(message []byte) []byte {
	blockSizeBits := 512
	metadataBits := 64
	msgLenBits := len(message) * 8

	zeroBits := ((blockSizeBits - metadataBits) - (msgLenBits + 1)) % blockSizeBits
	if zeroBits < 0 {
		zeroBits += blockSizeBits
	}

	totalBits := msgLenBits + 1 + zeroBits + metadataBits
	finalLen := totalBits / 8
	buff := make([]byte, finalLen)
	copy(buff, message)
	buff[len(message)] = 0x80 // note that it also appends 7 zeros

	binary.BigEndian.PutUint64(buff[finalLen-8:], uint64(msgLenBits))
	return buff
}

func f(t int, b, c, d uint32) uint32 {
	switch {
	case 0 <= t && t <= 19:
		return (b & c) | ((^b) & d)
	case 40 <= t && t <= 59:
		return (b & c) | (b & d) | (c & d)
	case (20 <= t && t <= 39) || (60 <= t && t <= 79):
		return b ^ c ^ d
	default:
		panic("t out of bound")
	}
}

func (s *SHA1) Hash(message []byte) [20]byte {
	s.state[0] = 0x67452301
	s.state[1] = 0xEFCDAB89
	s.state[2] = 0x98BADCFE
	s.state[3] = 0x10325476
	s.state[4] = 0xC3D2E1F0

	blocks := pad(message)
	var w [80]uint32
	var a, b, c, d, e uint32
	numBlocks := len(blocks) / 64

	for block := 0; block < numBlocks; block++ {

		blockStart := block * 64

		for i := 0; i < 16; i++ {
			w[i] = binary.BigEndian.Uint32(blocks[blockStart+i*4 : blockStart+i*4+4])
		}

		for t := 16; t < 80; t++ {
			w[t] = rotateLeft((w[t-3] ^ w[t-8] ^ w[t-14] ^ w[t-16]), 1)
		}

		a = s.state[0]
		b = s.state[1]
		c = s.state[2]
		d = s.state[3]
		e = s.state[4]

		for t := 0; t < 80; t++ {
			temp := rotateLeft(a, 5) + f(t, b, c, d) + e + w[t] + k[t/20]
			e = d
			d = c
			c = rotateLeft(b, 30)
			b = a
			a = temp
		}
		s.state[0] += a
		s.state[1] += b
		s.state[2] += c
		s.state[3] += d
		s.state[4] += e

	}
	var digest [20]byte
	binary.BigEndian.PutUint32(digest[0:], s.state[0])
	binary.BigEndian.PutUint32(digest[4:], s.state[1])
	binary.BigEndian.PutUint32(digest[8:], s.state[2])
	binary.BigEndian.PutUint32(digest[12:], s.state[3])
	binary.BigEndian.PutUint32(digest[16:], s.state[4])
	return digest

}
