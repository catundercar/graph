package mmkv

import (
	"github.com/dominikbraun/graph"
	mmkv2 "github.com/dominikbraun/graph/third_party/tencent.com/mmkv"
)

type NS string

func (ns NS) String() string {
	return string(ns)
}

var (
	Vertex     NS = "vertex"
	EdgeTo     NS = "edge_to"
	EdgeSrc    NS = "edge_src"
	Predicates NS = "predicates"
)

// KV is a wrapper of kv storage.
type KV[K comparable, V any] interface {
	WithNS(ns NS) KV[K, V]
	Set(key K, value V) error
	Get(key K) (V, error)
	Del(key K) error
}

type mmkvWrapper[K string, V []byte] struct {
	ns NS
	mmkv2.MMKV
}

func (m *mmkvWrapper[K, V]) WithNS(ns NS) KV[K, V] {
	if m.ns == ns {
		return m
	}
	ctol := mmkv2.MMKVWithIDAndMode(ns.String(), mmkv2.MMKV_MULTI_PROCESS)
	m2 := &mmkvWrapper[K, V]{
		ns:   ns,
		MMKV: ctol,
	}
	return m2
}

func (m *mmkvWrapper[K, V]) Set(key K, value V) error {
	m.SetBytes(value, string(key))
	return nil
}

func (m *mmkvWrapper[K, V]) Get(key K) (V, error) {
	if m.Contains(string(key)) {
		return m.GetBytes(string(key)), nil
	}
	return nil, graph.ErrVertexNotFound
}

func (m *mmkvWrapper[K, V]) Del(key K) error {
	if m.Contains(string(key)) {
		m.RemoveKey(string(key))
		return nil
	}
	return graph.ErrVertexNotFound
}
