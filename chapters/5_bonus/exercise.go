package exercise

func Fold[T any](slice []T, f func(T, T) T, init T) T
