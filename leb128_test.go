package leb128

import (
	"bytes"
	"testing"
)

func TestUint64(t *testing.T) {
	data := []uint64{0, 1, 34, 5436345, 42034823, 34}
	buf := &bytes.Buffer{}
	for _, i := range data {
		err := WriteUint64(buf, i)
		if err != nil {
			t.Error(err)
		}
	}
	for _, i := range data {
		j, err := ReadUint64(buf)
		if err != nil {
			t.Error(err)
		}
		if i != j {
			t.Errorf("%d != %d", i, j)
		}
	}
}

func TestUint(t *testing.T) {
	data := []uint{0, 1, 34, 5436345, 42034823, 34}
	buf := &bytes.Buffer{}
	for _, i := range data {
		err := WriteUint(buf, i)
		if err != nil {
			t.Error(err)
		}
	}
	for _, i := range data {
		j, err := ReadUint(buf)
		if err != nil {
			t.Error(err)
		}
		if i != j {
			t.Errorf("%d != %d", i, j)
		}
	}
}
