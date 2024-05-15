package hash_map

import (
	"fmt"
	"testing"
)

func TestArrayHashMap_Get(t *testing.T) {
	m := NewArrayHashMap()
	m.Put(1, "1")
	fmt.Printf("ArrayHashMap get: key: %d, value: %+v\n", 1, m.Get(1))
}

func TestArrayHashMap_Put(t *testing.T) {
	m := NewArrayHashMap()
	m.Put(1, "1")
	fmt.Printf("ArrayHashMap put: %+v\n", m)
}

func TestArrayHashMap_Remove(t *testing.T) {
	m := NewArrayHashMap()
	m.Put(1, "1")
	fmt.Printf("ArrayHashMap remove: %+v\n", m)
	m.Remove(1)
	fmt.Printf("ArrayHashMap remove: %+v\n", m)

}

func TestNewArrayHashMap(t *testing.T) {
	m := NewArrayHashMap()
	fmt.Printf("NewArrayHashMap: %+v", m)
}
