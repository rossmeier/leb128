package leb128

import (
	"io"
)

// WriteUint64 encodes n with LEB128 and writes it to the given writer
func WriteUint64(w io.Writer, n uint64) error {
	more := true
	b := []byte{0}
	for more {
		b[0] = byte(n & 0x7F)
		n >>= 7
		if n == 0 {
			more = false
		} else {
			b[0] = b[0] | 0x80
		}
		_, err := w.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteUint encodes n with LEB128 and writes it to the given writer
func WriteUint(w io.Writer, n uint) error {
	more := true
	b := []byte{0}
	for more {
		b[0] = byte(n & 0x7F)
		n >>= 7
		if n == 0 {
			more = false
		} else {
			b[0] = b[0] | 0x80
		}
		_, err := w.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadUint64 decodes LEB128-encoded stream into a uint64
func ReadUint64(r io.Reader) (uint64, error) {
	var result uint64
	var shift uint
	b := []byte{0}
	for {
		_, err := io.ReadFull(r, b)
		if err != nil {
			return 0, err
		}
		result |= (uint64(0x7F & b[0])) << shift
		if b[0]&0x80 == 0 {
			break
		}
		shift += 7
	}
	return result, nil
}

// ReadUint decodes LEB128-encoded stream into a uint
func ReadUint(r io.Reader) (uint, error) {
	var result uint
	var shift uint
	b := []byte{0}
	for {
		_, err := io.ReadFull(r, b)
		if err != nil {
			return 0, err
		}
		result |= (uint(0x7F & b[0])) << shift
		if b[0]&0x80 == 0 {
			break
		}
		shift += 7
	}
	return result, nil
}
