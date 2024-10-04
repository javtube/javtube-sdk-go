package maps

import (
	"bytes"
	"encoding/json"

	"github.com/elliotchance/orderedmap/v2"
	jsoniter "github.com/json-iterator/go"
)

var (
	_ json.Marshaler   = (*OrderedMap[int, any])(nil)
	_ json.Unmarshaler = (*OrderedMap[int, any])(nil)
)

type OrderedMap[K comparable, V any] struct {
	*orderedmap.OrderedMap[K, V]
	escapeHTML bool
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{OrderedMap: orderedmap.NewOrderedMap[K, V]()}
}

func (m *OrderedMap[K, V]) SetEscapeHTML(on bool) {
	m.escapeHTML = on
}

func (m *OrderedMap[K, V]) Set(key K, value V) bool {
	return m.OrderedMap.Set(key, value)
}

func (m *OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('{')
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(m.escapeHTML)
	for el := m.Front(); el != nil; el = el.Next() {
		if el != m.Front() {
			buf.WriteByte(',')
		}
		// add key
		if err := enc.Encode(el.Key); err != nil {
			return nil, err
		}
		buf.WriteByte(':')
		// add value
		if err := enc.Encode(el.Value); err != nil {
			return nil, err
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (m *OrderedMap[K, V]) UnmarshalJSON(data []byte) error {
	if m.OrderedMap == nil {
		m.OrderedMap = orderedmap.NewOrderedMap[K, V]()
	}
	temp := make(map[K]V)
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	root := jsoniter.Get(data)
	for _, key := range root.Keys() {
		k := any(key).(K)
		m.Set(k, temp[k])
	}
	return nil
}