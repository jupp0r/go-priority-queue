package pq_test

import (
	"sort"
	"testing"

	lib "quantime.ai/go-priority-queue/lib/pq"
)

func TestPriorityQueue(t *testing.T) {
	pq := lib.New()
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

func TestPriorityQueueUpdate(t *testing.T) {
	pq := lib.New()
	pq.Insert("foo", 3)
	pq.Insert("bar", 4)
	pq.UpdatePriority("bar", 2)

	item, err := pq.Pop()
	if err != nil {
		t.Fatal(err.Error())
	}

	if item.(string) != "bar" {
		t.Fatal("priority update failed")
	}
}

func TestPriorityQueueLen(t *testing.T) {
	pq := lib.New()
	if pq.Len() != 0 {
		t.Fatal("empty queue should have length of 0")
	}

	pq.Insert("foo", 1)
	pq.Insert("bar", 1)
	if pq.Len() != 2 {
		t.Fatal("queue should have lenght of 2 after 2 inserts")
	}
}

func TestDoubleAddition(t *testing.T) {
	pq := lib.New()
	pq.Insert("foo", 2)
	pq.Insert("bar", 3)
	pq.Insert("bar", 1)

	if pq.Len() != 2 {
		t.Fatal("queue should ignore inserting the same element twice")
	}

	item, _ := pq.Pop()
	if item.(string) != "foo" {
		t.Fatal("queue should ignore duplicate insert, not update existing item")
	}
}

func TestPopEmptyQueue(t *testing.T) {
	pq := lib.New()
	_, err := pq.Pop()
	if err == nil {
		t.Fatal("should produce error when performing pop on empty queue")
	}
}

func TestUpdateNonExistingItem(t *testing.T) {
	pq := lib.New()

	pq.Insert("foo", 4)
	pq.UpdatePriority("bar", 5)

	if pq.Len() != 1 {
		t.Fatal("update should not add items")
	}

	item, _ := pq.Pop()
	if item.(string) != "foo" {
		t.Fatalf("update should not overwrite item, expected \"foo\", got \"%v\"", item.(string))
	}
}
