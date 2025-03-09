package bstream

import (
	"bytes"
	"errors"
	"os"
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

func TestLoadFileNotExists(t *testing.T) {
	_, err := LoadFile("nonexistentfile.txt")
	if err == nil {
		t.Fatal("Expected error for nonexistent file, got nil")
	}
	if !errors.Is(err, os.ErrNotExist) && err.Error() != "file does not exist: nonexistentfile.txt" {
		t.Errorf("Expected specific file not exist error, got: %v", err)
	}
}

func TestEncodeNil(t *testing.T) {
	_, err := Encode(nil)
	if err == nil || err.Error() != "cannot encode nil value" {
		t.Errorf("Expected 'cannot encode nil value' error, got: %v", err)
	}
}

func TestDecodeNil(t *testing.T) {
	err := Decode([]byte("test"), nil)
	if err == nil || err.Error() != "cannot decode into nil value" {
		t.Errorf("Expected 'cannot decode into nil value' error, got: %v", err)
	}
}

func TestDecodeEmptyData(t *testing.T) {
	var num int
	err := Decode([]byte{}, &num)
	if err == nil || err.Error() != "cannot decode empty data" {
		t.Errorf("Expected 'cannot decode empty data' error, got: %v", err)
	}
}