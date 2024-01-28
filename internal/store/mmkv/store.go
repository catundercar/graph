package mmkv

import (
	"encoding/json"
	"github.com/dominikbraun/graph"
	mmkv2 "github.com/dominikbraun/graph/third_party/tencent.com/mmkv"
	"strings"
)

type options struct {
	dataPath string // the path to save mmkv resources.
}

// Option configures how we set up the mmkv store.
type Option interface {
	apply(*options)
}

type funcOption struct {
	fn func(*options)
}

func (fo *funcOption) apply(opt *options) {
	fo.fn(opt)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{fn: f}
}

func WithDataPath(path string) Option {
	return newFuncOption(func(opt *options) {
		opt.dataPath = path
	})
}

// mmkv implement the graph.Store by mmkv[https://github.com/Tencent/MMKV].
// Note: the hash value only support string type.
type mmkv[K string, T any] struct {
	opts *options
	shim KV[string, []byte]
}

func NewMMKVStore[K string, T any](opts ...Option) (graph.Store[string, T], error) {
	store := &mmkv[string, T]{
		opts: &options{},
		shim: &mmkvWrapper[string, []byte]{},
	}
	for _, op := range opts {
		op.apply(store.opts)
	}
	mmkv2.InitializeMMKV(store.opts.dataPath)
	//store.MMKV = mmkv2.MMKVWithIDAndMode("vertex", mmkv2.MMKV_MULTI_PROCESS)
	//store.MMKV = mmkv2.MMKVWithIDAndMode("edge_to", mmkv2.MMKV_MULTI_PROCESS)
	//store.MMKV = mmkv2.MMKVWithIDAndMode("edge_source", mmkv2.MMKV_MULTI_PROCESS)
	//store.MMKV = mmkv2.MMKVWithIDAndMode("predicates", mmkv2.MMKV_MULTI_PROCESS)
	return store, nil
}

func (m *mmkv[K, T]) AddVertex(hash string, value T, properties graph.VertexProperties) (err error) {

	pred := predicate{properties}
	for _, pair := range pred.PairsWithSubject(hash) {
		if err = m.shim.WithNS(Vertex).Set(pair.Key, pair.Value); err != nil {
			return
		}
	}

	// properties
	predicateNames := pred.Predicates()
	if err = m.shim.WithNS(Predicates).Set(hash, []byte(strings.Join(predicateNames, PredicatePlaceholder))); err != nil {
		return
	}

	var byt []byte
	byt, err = json.Marshal(value)
	if err != nil {
		return
	}
	if err = m.shim.WithNS(Vertex).Set(hash, byt); err != nil {
		return
	}
	return
}

func (m *mmkv[K, T]) Vertex(hash string) (v T, vps graph.VertexProperties, err error) {
	byt, err := m.shim.WithNS(Vertex).Get(hash)
	if err != nil {
		return
	}
	if err = json.Unmarshal(byt, &v); err != nil {
		return
	}

	vps = graph.VertexProperties{
		Attributes: make(map[string]string),
	}
	predicateNames, err := m.shim.WithNS(Predicates).Get(hash)
	if err != nil {
		return
	}
	for _, name := range strings.Split(string(predicateNames), PredicatePlaceholder) {
		var value []byte
		value, err = m.shim.WithNS(Vertex).Get(PredicateName(hash, name))
		if err != nil {
			return v, graph.VertexProperties{}, err
		}
		vps.Attributes[name] = string(value)
	}
	return
}

func (m *mmkv[K, T]) RemoveVertex(hash string) error {
	if err := m.shim.WithNS(Vertex).Del(hash); err != nil {
		return err
	}

	predicateNames, err := m.shim.WithNS(Predicates).Get(hash)
	if err != nil {
		return err
	}
	for _, name := range strings.Split(string(predicateNames), PredicatePlaceholder) {
		if err = m.shim.WithNS(Vertex).Del(PredicateName(hash, name)); err != nil {
			return err
		}
	}
	return nil
}

func (m *mmkv[K, T]) ListVertices() ([]K, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mmkv[K, T]) VertexCount() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mmkv[K, T]) AddEdge(sourceHash, targetHash K, edge graph.Edge[K]) error {
	//TODO implement me
	panic("implement me")
}

func (m *mmkv[K, T]) UpdateEdge(sourceHash, targetHash K, edge graph.Edge[K]) error {
	//TODO implement me
	panic("implement me")
}

func (m *mmkv[K, T]) RemoveEdge(sourceHash, targetHash K) error {
	//TODO implement me
	panic("implement me")
}

func (m *mmkv[K, T]) Edge(sourceHash, targetHash K) (graph.Edge[K], error) {
	//TODO implement me
	panic("implement me")
}

func (m *mmkv[K, T]) ListEdges() ([]graph.Edge[K], error) {
	//TODO implement me
	panic("implement me")
}

func (m *mmkv[K, T]) EdgeCount() (int, error) {
	//TODO implement me
	panic("implement me")
}
