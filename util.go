package common_util

import "sync"

// LockItem 通过锁保护 item 的 访问
type LockItem[T any] struct {
	item T
	lock sync.Mutex
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
	}
}

// Action 对锁保护的内容进行更新
func (s *LockItem[T]) Action(act func(t *T)) {
	defer s.Unlock()
	act(s.Lock())
}

// Get 获取内容
func (s *LockItem[T]) Get() T {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.item
}

// Set 设置内容
func (s *LockItem[T]) Set(t T) {
	s.lock.Lock()
	defer s.lock.Unlock()
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
