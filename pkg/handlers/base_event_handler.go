package handlers

type BaseEventHandler[T any] struct {
	measurement T
	measured    chan T
}
