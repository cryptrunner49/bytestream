package bstream

import (
	"bytes"
	"encoding/gob"
	"os"
)

// Encode serializes the given value into []byte.
// If v is []byte or string, it returns the raw bytes.
// For other types, it uses Gob encoding.
func Encode(v interface{}) ([]byte, error) {
	switch val := v.(type) {
	case []byte:
		return val, nil
	case string:
		return []byte(val), nil
	default:
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(v)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
}

// Decode deserializes the given []byte into v.
// If v is *[]byte or *string, it sets the value directly.
// For other types, it uses Gob decoding.
func Decode(data []byte, v interface{}) error {
	switch val := v.(type) {
	case *[]byte:
		*val = data
		return nil
	case *string:
		*val = string(data)
		return nil
	default:
		buf := bytes.NewBuffer(data)
		dec := gob.NewDecoder(buf)
		return dec.Decode(v)
	}
}

// LoadFile reads the file at the given path and returns its contents as []byte.
func LoadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}