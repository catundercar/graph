package mmkv

import (
	"github.com/dominikbraun/graph"
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func testNewMMKVStore[_, T any](t *testing.T) *mmkv[string, T] {
	store, err := NewMMKVStore[string, T](WithDataPath("testdata"))
	if err != nil {
		t.Fatal(err)
	}
	return store.(*mmkv[string, T])
}

func Test_mmkv_AddEdge(t *testing.T) {
	// TODO
}

func Test_mmkv_AddVertex(t *testing.T) {
	store := testNewMMKVStore[string, string](t)
	hash := "hello"
	value := "world"
	properties := graph.VertexProperties{
		Attributes: map[string]string{
			"hello": "world",
		},
	}

	Convey("TestMMKVStore", t, func() {
		Convey("Add Vertex", func() {
			So(store.AddVertex(hash, value, properties) == nil, ShouldBeTrue)
		})
		Convey("Get Vertex", func() {
			vertex, prop, err := store.Vertex(hash)
			So(err == nil, ShouldBeTrue)
			ShouldEqual(vertex, value)
			ShouldEqual(prop.Attributes["hello"], value)
		})
		Convey("Delete Vertex", func() {
			So(store.RemoveVertex(hash) == nil, ShouldBeTrue)
			_, _, err := store.Vertex(hash)
			ShouldEqual(err, graph.ErrVertexNotFound)
		})
	})
}

func Test_mmkv_Edge(t *testing.T) {
	// TODO
}

func Test_mmkv_EdgeCount(t *testing.T) {
	// TODO
}

func Test_mmkv_ListEdges(t *testing.T) {
	// TODO
}

func Test_mmkv_ListVertices(t *testing.T) {
	// TODO
}

func Test_mmkv_RemoveEdge(t *testing.T) {
	// TODO
}
