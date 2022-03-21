package pq

import (
	"math"
	"sort"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := New[float64]()
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

		i := item
		if e != i {
			t.Fatalf("expected %v, got %v", e, i)
		}
	}
}

func TestPriorityQueueUpdate(t *testing.T) {
	pq := New[string]()
	pq.Insert("foo", 3)
	pq.Insert("bar", 4)
	pq.UpdatePriority("bar", 2)

	item, err := pq.Pop()
	if err != nil {
		t.Fatal(err.Error())
	}

	if item != "bar" {
		t.Fatal("priority update failed")
	}
}

func TestPriorityQueueLen(t *testing.T) {
	pq := New[string]()
	if pq.Len() != 0 {
		t.Fatal("empty queue should have length of 0")
	}

	pq.Insert("foo", 1)
	pq.Insert("bar", 1)
	if pq.Len() != 2 {
		t.Fatal("queue should have lenght of 2 after 2 inserts")
	}
}

func TestItemHeapLess(t *testing.T) {
	h := itemHeap[int]{
		&item[int]{priority: math.Inf(1)},
		&item[int]{priority: math.Inf(1)},
		&item[int]{priority: math.Inf(-1)},
		&item[int]{priority: math.Inf(-1)},
	}

	// test all pairwise elements: 0,1 then 2,3 then...
	for i := 0; i < len(h); i += 2 {
		if h.Less(i, i+1) {
			t.Fatalf("%v should not have less priority than %v", h[i], h[i+1])
		}
	}
}

func TestDoubleAddition(t *testing.T) {
	pq := New[string]()
	pq.Insert("foo", 2)
	pq.Insert("bar", 3)
	pq.Insert("bar", 1)

	if pq.Len() != 2 {
		t.Fatal("queue should ignore inserting the same element twice")
	}

	item, _ := pq.Pop()
	if item != "foo" {
		t.Fatal("queue should ignore duplicate insert, not update existing item")
	}
}

func TestPopEmptyQueue(t *testing.T) {
	pq := New[float32]()
	_, err := pq.Pop()
	if err == nil {
		t.Fatal("should produce error when performing pop on empty queue")
	} else if err != ErrEmptyQueue {
		t.Fatalf("error should be equal to %v, got %v", ErrEmptyQueue, err)
	}
}

func TestUpdateNonExistingItem(t *testing.T) {
	pq := New[string]()

	pq.Insert("foo", 4)
	pq.UpdatePriority("bar", 5)

	if pq.Len() != 1 {
		t.Fatal("update should not add items")
	}

	item, _ := pq.Pop()
	if item != "foo" {
		t.Fatalf("update should not overwrite item, expected \"foo\", got \"%v\"", item)
	}
}
