package models

import "encoding/json"

type Metadata map[string]any

func (m Metadata) Get(key string) (any, bool) {
	if m == nil {
		return nil, false
	}
	val, ok := m[key]

	return val, ok
}

func (m Metadata) Set(key string, value any) {
	if m == nil {
		return
	}
	m[key] = value
}

func (m Metadata) Delete(key string) {
	if m == nil {
		return
	}
	delete(m, key)
}

func (m Metadata) ToJSON() (string, error) {
	if m == nil {
		return "{}", nil
	}

	bytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
