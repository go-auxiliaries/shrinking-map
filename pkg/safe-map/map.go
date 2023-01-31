package safe_map

import (
	"sync"
)

// Map - map implementation that shrinks automatically once limit of deleted elements reaches the limit
// It is made thread safe via `sync.RWMutex`.
// In general, it is faster than sync.Map, there are some benchmarks to prove that
type Map[T comparable, V any] struct {
	values  map[T]V
	deleted uint64
	limit   uint64
	lock    sync.RWMutex
}

func New[T comparable, V any](limit uint64) *Map[T, V] {
	return &Map[T, V]{
		values:  make(map[T]V),
		deleted: 0,
		limit:   limit,
	}
}

func (m *Map[T, V]) Set(key T, val V) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.values[key] = val
}

func (m *Map[T, V]) GetOrSet(key T, val V) V {
	m.lock.Lock()
	defer m.lock.Unlock()
	v, ok := m.values[key]
	if ok {
		return v
	}
	m.values[key] = val
	return val
}

func (m *Map[T, V]) GetAndDelete(key T) (V, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	v, ok := m.values[key]
	if ok {
		delete(m.values, key)
		return v, true
	}
	return v, false
}

func (m *Map[T, V]) Delete(keys ...T) {
	m.lock.Lock()
	d := uint64(0)
	defer func() {
		m.addDeleted(d)
		m.lock.Unlock()
	}()

	for _, key := range keys {
		d++
		delete(m.values, key)
	}
}

func (m *Map[T, V]) Get(key T) V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.values[key]
}

func (m *Map[T, V]) Get2(key T) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	item, ok := m.values[key]
	return item, ok
}

func (m *Map[T, V]) Shrink() {
	m.lock.Lock()
	defer m.lock.Unlock()
	var newValues map[T]V
	if uint64(len(m.values)) >= m.deleted {
		newValues = make(map[T]V, uint64(len(m.values))-m.deleted)
	} else {
		newValues = make(map[T]V, 0)
	}

	for key, val := range m.values {
		newValues[key] = val
	}
	m.values = newValues
	m.deleted = 0
}

func (m *Map[T, V]) LockSession(body func(map[T]V) uint64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.addDeleted(body(m.values))
}

func (m *Map[T, V]) RLockSession(body func(map[T]V)) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	body(m.values)
}

func (m *Map[T, V]) RLock() {
	m.lock.RLock()
}

func (m *Map[T, V]) RUnlock() {
	m.lock.RUnlock()
}

func (m *Map[T, V]) Lock() {
	m.lock.Lock()
}

func (m *Map[T, V]) ULock() {
	m.lock.Unlock()
}

func (m *Map[T, V]) Values() map[T]V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.values
}

func (m *Map[T, V]) addDeleted(n uint64) {
	m.deleted += n
	if m.limit != 0 && m.deleted > m.limit {
		go m.Shrink()
	}
}
