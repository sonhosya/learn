package hash_map

type Pair struct {
	Key int
	Val string
}
type ArrayHashMap struct {
	Buckets []*Pair
}

func NewArrayHashMap() *ArrayHashMap {
	return &ArrayHashMap{
		Buckets: make([]*Pair, 100),
	}
}

func (h *ArrayHashMap) Put(key int, val string) {
	idx := h.hashFunc(key)
	h.Buckets[idx] = &Pair{Key: key, Val: val}
}

func (h *ArrayHashMap) hashFunc(key int) int {
	return key % len(h.Buckets)
}

func (h *ArrayHashMap) Get(key int) string {
	idx := h.hashFunc(key)
	pair := h.Buckets[idx]
	if pair != nil {
		return pair.Val
	}
	return ""
}

func (h *ArrayHashMap) Remove(key int) {
	idx := h.hashFunc(key)
	h.Buckets[idx] = nil
}
