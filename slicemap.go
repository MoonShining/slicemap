package slicemap

import (
	"bytes"
	"sync"
)

func Borrow() *SliceMap {
	sm := *pool.Get().(*SliceMap)
	for i, kv := range sm {
		kv.Key = kv.Key[:0]
		kv.Value = kv.Value[:0]
		sm[i] = kv
	}
	return &sm
}

func GiveBack(sm *SliceMap) {
	pool.Put(sm)
}

var pool = sync.Pool{
	New: func() interface{} {
		return &SliceMap{}
	},
}

type KV struct {
	Key   []byte
	Value []byte
}

type SliceMap []KV

func (sm *SliceMap) Add(k, v []byte) {
	kvs := *sm
	if cap(kvs) > len(kvs) {
		kvs = kvs[:len(kvs)+1]
	} else {
		kvs = append(kvs, KV{})
	}
	kv := &kvs[len(kvs)-1]
	kv.Key = append(kv.Key[:0], k...)
	kv.Value = append(kv.Value[:0], v...)
	*sm = kvs
}

func (sm *SliceMap) Get(k []byte) []byte {
	for _, kv := range *sm {
		if bytes.Equal(kv.Key, k) {
			return kv.Value
		}
	}
	return nil
}

func (sm *SliceMap) MarshalJSON() []byte {
	var buffer bytes.Buffer

	buffer.WriteString("{")

	for i, kv := range *sm {
		if len(kv.Key) == 0 {
			continue
		}
		buffer.WriteString(`"`)
		buffer.Write(kv.Key)
		buffer.WriteString(`":"`)
		buffer.Write(kv.Value)
		buffer.WriteString(`"`)
		if i != len(*sm)-1 {
			buffer.WriteString(`,`)
		}
	}

	buffer.WriteString("}")
	return buffer.Bytes()
}
