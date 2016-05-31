package pq

import (
	"sort"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := New()
	elements := []float64{5, 3, 7, 8, 6, 2, 9}
	for _, e := range elements {
		pq.Insert(e, e)
	}

	sort.Float64s(elements)
	for _, e := range elements {
		item, err := pq.Pop()
		if err != nil {
			t.Fatalf(err.Error())
		}

		i := item.(float64)
		if e != i {
			t.Fatalf("expected %v, got %v", e, i)
		}
	}
}
