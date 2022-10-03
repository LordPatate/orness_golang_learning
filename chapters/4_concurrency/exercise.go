package exercise

type Comparable interface {
	int | float32 | float64 | uint
}

func MergeSort[T Comparable](slice []T) []T
