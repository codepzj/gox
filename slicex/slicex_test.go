package slicex

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3}
	if !Contains(arr, 2) {
		t.Error("Contains failed: expected true for 2")
	}
	if Contains(arr, 5) {
		t.Error("Contains failed: expected false for 5")
	}
}

func TestIndexOf(t *testing.T) {
	arr := []string{"a", "b", "c"}
	if idx := IndexOf(arr, "b"); idx != 1 {
		t.Errorf("IndexOf failed: expected 1 got %d", idx)
	}
	if idx := IndexOf(arr, "x"); idx != -1 {
		t.Errorf("IndexOf failed: expected -1 got %d", idx)
	}
}

func TestUnique(t *testing.T) {
	arr := []int{1, 2, 2, 3, 1}
	expected := []int{1, 2, 3}
	result := Unique(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unique failed: expected %v got %v", expected, result)
	}

	arrStr := []string{"a", "b", "a"}
	expectedStr := []string{"a", "b"}
	resultStr := Unique(arrStr)
	if !reflect.DeepEqual(resultStr, expectedStr) {
		t.Errorf("Unique failed: expected %v got %v", expectedStr, resultStr)
	}
}

func TestConcat(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	expected := []int{1, 2, 3, 4}
	result := Concat(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Concat failed: expected %v got %v", expected, result)
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{4, 3, 2, 1}
	result := Reverse(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Reverse failed: expected %v got %v", expected, result)
	}

	empty := []int{}
	resultEmpty := Reverse(empty)
	if len(resultEmpty) != 0 {
		t.Errorf("Reverse failed on empty slice, expected empty got %v", resultEmpty)
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := []int{2, 4, 6}
	result := Map(arr, func(v int) int {
		return v * 2
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map failed: expected %v got %v", expected, result)
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}
	result := Filter(arr, func(v int) bool {
		return v%2 == 0
	})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter failed: expected %v got %v", expected, result)
	}
}

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	sum := Reduce(arr, func(acc, v int) int {
		return acc + v
	}, 0)
	if sum != 10 {
		t.Errorf("Reduce sum failed: expected 10 got %v", sum)
	}

	words := []string{"a", "b", "c"}
	concat := Reduce(words, func(acc, v string) string {
		return acc + v
	}, "")
	if concat != "abc" {
		t.Errorf("Reduce concat failed: expected 'abc' got '%v'", concat)
	}
}
