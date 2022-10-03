package exercise

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	slice := []int{
		1, 200, 3, 40, 5, 600, 7, 80, 9, 0,
	}
	expected := []int{
		0, 1, 3, 5, 7, 9, 40, 80, 200, 600,
	}
	slice = MergeSort(slice)
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("got %v", slice)
	}
}
