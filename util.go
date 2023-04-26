package common_util

import (
	sync "github.com/sasha-s/go-deadlock"
	"sync/atomic"
)

// LockItem 通过锁保护 item 的 访问
type LockItem[T any] struct {
	item T
	lock *sync.Mutex
}

// Lock 获取锁的同时返回保护值的指针
func (s *LockItem[T]) Lock() *T {
	s.lock.Lock()
	return &s.item
}

// Unlock 解锁
func (s *LockItem[T]) Unlock() {
	s.lock.Unlock()
}

// NewLockItem 初始化，参数 t 可以传入初始值，只会解析索引 0 ，不传代表使用零值
func NewLockItem[T any](t ...T) *LockItem[T] {
	return &LockItem[T]{
		item: func() T {
			if len(t) > 0 {
				return t[0]
			}
			return Zero[T]()
		}(),
		lock: new(sync.Mutex),
	}
}

// Action 对锁保护的内容进行更新
func (s *LockItem[T]) Action(act func(t *T)) {
	defer s.Unlock()
	act(s.Lock())
}

// Get 获取内容
func (s *LockItem[T]) Get() T {
	defer s.lock.Unlock()
	s.lock.Lock()
	return s.item
}

// Set 设置内容
func (s *LockItem[T]) Set(t T) {
	defer s.lock.Unlock()
	s.lock.Lock()
	s.item = t
}

// Ptr 返回 t 的指针
func Ptr[T any](t T) *T {
	return &t
}

// Zero 返回类型的零值
func Zero[T any]() (zero T) {
	return zero
}

type Tuple[T1 any, T2 any] struct {
	A T1
	B T2
}

func BuildTuple[T1 any, T2 any](a T1, b T2) Tuple[T1, T2] {
	return Tuple[T1, T2]{
		A: a, B: b,
	}
}

type AtomicVal[T any] struct {
	v atomic.Value
}

func NewAtomicVal[T any](val ...T) *AtomicVal[T] {
	v := &AtomicVal[T]{}
	if len(val) > 0 {
		v.Store(val[0])
	}
	return v
}
func (a *AtomicVal[T]) Load() T {
	val, ok := a.v.Load().(T)
	if ok {
		return val
	}
	return Zero[T]()
}
func (a *AtomicVal[T]) Store(t T) {
	a.v.Store(t)
}
