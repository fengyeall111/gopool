package gopool

import (
	"arena"
	"sync"
)

// T use new to init it
type Pool[T any] interface {
	Get() T
	Put(T)
	Destory() // you should not use pool after destory it,even the object is alloc from pool
}

type pool[T any] struct {
	*arena.Arena
	*sync.Pool
}

func (p *pool[T]) Destory() {
	p.Pool = nil
	p.Arena.Free()
}

func (p *pool[T]) Get() T {
	return p.Pool.Get().(T)
}
func (p *pool[T]) Put(t T) {
	p.Pool.Put(t)
}

type New[T any] func() T

func NewPool[T any](f New[T]) Pool[T] {
	a := arena.NewArena()
	p := arena.New[sync.Pool](a)
	p.New = func() any {
		return f()
	}
	return &pool[T]{
		a,
		p,
	}
}
