package shrinking_map

// Map - map implementation that shrinks automatically once limit of deleted elements reaches the limit
// It is not thread safe, if you are planning using it in parallel, please use `github.com/go-auxiliaries/shrinking-map/pkg/safe-tight-map` instead
type Map[T comparable, V any] struct {
	values  map[T]V
	deleted uint64
	limit   uint64
}

func New[T comparable, V any](limit uint64) *Map[T, V] {
	return &Map[T, V]{
		values:  make(map[T]V),
		deleted: 0,
		limit:   limit,
	}
}

func (m *Map[T, V]) Set(key T, val V) {
	m.values[key] = val
}

func (m *Map[T, V]) Delete(keys ...T) {
	d := uint64(0)
	defer func() {
		m.deleted += d
		if m.limit != 0 && m.deleted > m.limit {
			m.Shrink()
		}
	}()

	for _, key := range keys {
		d++
		delete(m.values, key)
	}
}

func (m *Map[T, V]) Get(key T) V {
	return m.values[key]
}

func (m *Map[T, V]) Get2(key T) (V, bool) {
	item, ok := m.values[key]
	return item, ok
}

func (m *Map[T, V]) GetOrSet(key T, val V) V {
	v, ok := m.values[key]
	if ok {
		return v
	}
	m.values[key] = val
	return val
}

func (m *Map[T, V]) GetAndDelete(key T) (V, bool) {
	v, ok := m.values[key]
	if ok {
		delete(m.values, key)
		return v, true
	}
	return v, false
}

func (m *Map[T, V]) Values() map[T]V {
	return m.values
}

func (m *Map[T, V]) Shrink() {
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

func (m *Map[T, V]) SetLimit(limit uint64) {
	m.limit = limit
}

func (m *Map[T, V]) GetLimit() uint64 {
	return m.limit
}
