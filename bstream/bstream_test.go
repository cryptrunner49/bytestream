package bstream

import (
	"bytes"
	"testing"
)

func TestEncodeDecodeString(t *testing.T) {
	original := "hello"
	data, err := Encode(original)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	var decoded string
	err = Decode(data, &decoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if decoded != original {
		t.Errorf("expected %q, got %q", original, decoded)
	}
}

func TestEncodeDecodeBytes(t *testing.T) {
	original := []byte{1, 2, 3}
	data, err := Encode(original)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	var decoded []byte
	err = Decode(data, &decoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if !bytes.Equal(decoded, original) {
		t.Errorf("expected %v, got %v", original, decoded)
	}
}

func TestEncodeDecodeStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	original := User{"Alice", 30}
	data, err := Encode(original)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	var decoded User
	err = Decode(data, &decoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if decoded != original {
		t.Errorf("expected %+v, got %+v", original, decoded)
	}
}

func TestEncodeDecodeSlice(t *testing.T) {
	original := []int{1, 2, 3}
	data, err := Encode(original)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	var decoded []int
	err = Decode(data, &decoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if len(decoded) != len(original) || decoded[0] != original[0] || decoded[1] != original[1] || decoded[2] != original[2] {
		t.Errorf("expected %v, got %v", original, decoded)
	}
}