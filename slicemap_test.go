package slicemap

import (
	"bytes"
	"testing"
)

func TestAddGet(t *testing.T) {
	key1, value1 := []byte("foo"), []byte("bar")
	key2, value2 := []byte("hello"), []byte("world")

	m := Borrow()
	m.Add(key1, value1)
	m.Add(key2, value2)

	bar := m.Get(key1)
	world := m.Get(key2)

	t.Log(string(bar))
	t.Log(string(world))

	if !bytes.Equal(value1, bar) {
		t.Fatal("get foo error")
	}
	if !bytes.Equal(value2, world) {
		t.Fatal("get hello error")
	}

	GiveBack(m)

	m = Borrow()
	if m.Get(key1) != nil {
		t.Fatal("borrow fail")
	}
	m.Add(key1, value1)
	m.Add(key2, value2)
	if !bytes.Equal(value1, bar) {
		t.Fatal("get foo again error")
	}

	data := m.MarshalJSON()
	t.Log(string(data))
}
