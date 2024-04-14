package SeekAfter32

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"testing"
)

func compareReader(br *bufio.Reader, bytes []byte) bool {
	for _, expected := range bytes {
		b, err := br.ReadByte()
		if err != nil || b != expected {
			return false
		}
	}
	_, err := br.ReadByte()
	return err != nil
}
func reader(str string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(str))
}

func TestEmpty(t *testing.T) {
	s := "hiii :3"
	r := reader(s)
	if SeekAfter32(r, []byte("")) != nil {
		t.Fatal("unexpected error")
	}
	if !compareReader(r, []byte(s)) {
		t.Fatal("unexpected reader output")
	}
}
func TestLong(t *testing.T) {
	str := strings.Repeat("a", 32)
	r := reader(str + "b")
	err := SeekAfter32(r, []byte(str))
	if err != nil {
		t.Fatal("unexpected error")
	}
	if !compareReader(r, []byte("b")) {
		t.Fatal("unexpected reader output")
	}
}
func TestTooLong(t *testing.T) {
	str := strings.Repeat("a", 33)
	err := SeekAfter32(reader(str), []byte(str))
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, BadLen) {
		t.Fatal("incorrect error")
	}
}
func TestEOF(t *testing.T) {
	r := reader("abc")
	err := SeekAfter32(r, []byte("abcd"))
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, io.EOF) {
		t.Fatal("incorrect error")
	}
}
func TestNormal(t *testing.T) {
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	after := "fghjklzxcvbnm"
	r := reader(strings.Join([]string{"qwertyuiopasd", after}, upper))
	err := SeekAfter32(r, []byte(upper))
	if err != nil {
		t.Fatal("unexpected error")
	}
	if !compareReader(r, []byte(after)) {
		t.Fatal("unexpected reader output")
	}
}
func TestInsidePartialMatch(t *testing.T) {
	r := reader("aaab aab")
	err := SeekAfter32(r, []byte("aab"))
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !compareReader(r, []byte(" aab")) {
		t.Fatal("unexpected reader output")
	}
}
