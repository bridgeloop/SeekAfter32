package SeekAfter31

import (
	"bufio"
	"errors"
)

var BadLen = errors.New("len(bytes) > 31")

func SeekAfter31(br *bufio.Reader, bytes []byte) error {
	sz := len(bytes)
	if sz > 31 {
		return BadLen
	}
	if sz == 0 {
		return nil
	}

	cache := make([]uint32, 256)

	matchers := ^uint32(0)
	for {
		b, err := br.ReadByte()
		if err != nil {
			return err
		}

		matchers <<= 1

		v := &cache[b]
		if *v&(1<<31) == 0 {
			*v = uint32(1 << 31)
			for it, by := range bytes {
				if by != b {
					*v |= 1 << it
				}
			}
		}

		// set each bit in matchers where bytes[Pos(bit)] != b
		matchers |= *v

		if matchers&(1<<(sz-1)) == 0 {
			return nil
		}
	}
}
