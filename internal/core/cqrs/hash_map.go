package cqrs

import "strings"

const separator = "."

type hashMap map[string]interface{}

func (h hashMap) has(key string) bool {
	parts := strings.Split(key, separator)
	if len(parts) > 1 && h.has(parts[0]) {
		v := h.get(parts[0])
		if innerHash, ok := v.(hashMap); ok {
			return innerHash.has(strings.Join(parts[1:], separator))
		}

		return false
	}

	_, ok := h[key]
	return ok
}

func (h hashMap) get(key string, defaultVal ...interface{}) interface{} {
	var def interface{}
	if len(defaultVal) > 0 {
		def = defaultVal[0]
	}

	parts := strings.Split(key, separator)
	if len(parts) > 1 {
		v := h.get(parts[0])

		if innerHash, ok := v.(hashMap); ok {
			return innerHash.get(strings.Join(parts[1:], separator))
		}

		return def
	}

	v, ok := h[key]
	if !ok {
		return def
	}

	return v
}

func (h hashMap) add(key string, val interface{}) {
	h[key] = val
}
