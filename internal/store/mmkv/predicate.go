package mmkv

import (
	"fmt"
	"github.com/dominikbraun/graph"
)

// Pair is a key-value pair which wrappers a property of vertex.
// Key is PredicateName: Hash(vertex) and the name of property.
// Value is the value of property.
type Pair struct {
	Key   string
	Value []byte
}

type predicate struct {
	graph.VertexProperties
}

func (p predicate) Predicates() []string {
	predicates := make([]string, 0, len(p.Attributes))
	for key := range p.Attributes {
		predicates = append(predicates, key)
	}
	return predicates
}

func (p predicate) PairsWithSubject(subject string) []Pair {
	var pairs = make([]Pair, 0, len(p.Attributes))
	for key, value := range p.Attributes {
		pairs = append(pairs, Pair{Key: PredicateName(subject, key), Value: []byte(value)})
	}
	return pairs
}

const PredicatePlaceholder = " "

func PredicateName(subject, property string) string {
	return fmt.Sprintf("<%s, %s>", property, subject)
}
