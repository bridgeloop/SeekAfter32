package SeekAfter32

import (
	"bufio"
	"errors"
)

var BadLen = errors.New("len(bytes) > 32")

func SeekAfter32(br *bufio.Reader, bytes []byte) error {
	sz := len(bytes)
	if sz > 32 {
		return BadLen
	}
	if sz == 0 {
		return nil
	}

	cache := make([]uint32, 256)
	in_cache := [4]uint64 { 0, 0, 0, 0 }

	matchers := ^uint32(0)
	for {
		b, err := br.ReadByte()
		if err != nil {
			return err
		}

		matchers <<= 1

		v := &(cache[b])

		in_idx := (b & 0b11000000) >> 6
		in_bit := uint64(1) << (b & 0b00111111)

		if (in_cache[in_idx] & in_bit) == 0 {
			in_cache[in_idx] |= in_bit
			for it, by := range bytes {
				if by != b {
					*v |= 1 << it
				}
			}
		}

		// set each bit in matchers where bytes[Pos(bit)] != b
		matchers |= *v

		if matchers & (1 << (sz - 1)) == 0 {
			return nil
		}
	}
}
