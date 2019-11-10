#### memory reused map

```
m := slicemap.Borrow()
defer slicemap.GiveBack(m)

m.Add([]byte("foo"), []byte("bar"))

bar := m.Get([]byte("foo"))
```
