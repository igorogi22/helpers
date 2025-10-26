package arrays_test

import (
	"testing"

	helpers "github.com/igorogi22/helpers"
)

func TestGenerateArray(t *testing.T) {
	t.Run("should generate array starting from default 1", func(t *testing.T) {
		got := helpers.CreateNumericArray(5)
		want := []int{1, 2, 3, 4, 5}

		for i := range want {
			if got[i] != want[i] {
				t.Errorf("expected %v, got %v", want, got)
			}
		}
	})

	t.Run("should generate array starting from custom value", func(t *testing.T) {
		got := helpers.CreateNumericArray(3, 10)
		want := []int{10, 11, 12}

		for i := range want {
			if got[i] != want[i] {
				t.Errorf("expected %v, got %v", want, got)
			}
		}
	})
}

func TestGetRandomItem(t *testing.T) {
	arr := []string{"Go", "Node", "Rust", "Python"}
	item := helpers.GetRandomItem(arr)

	found := false
	for _, v := range arr {
		if v == item {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("item %v não encontrado no array original", item)
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	helpers.Reverse(arr)
	want := []int{5, 4, 3, 2, 1}

	for i := range want {
		if arr[i] != want[i] {
			t.Errorf("expected %v, got %v", want, arr)
		}
	}
}
