/*
 * License: GPL-3.0 license | GNU General Public License v3.0
 * Purpose: This file contains unit tests for the bstream package, which handles encoding and decoding of various data types.
 * Key Components:
 * - Tests for encoding and decoding strings, bytes, structs, slices, and error handling.
 * Role within the Codebase: Ensures the correctness and reliability of the bstream package's encoding and decoding functionality.
 */

 package bstream

 import (
	 "bytes"
	 "errors"
	 "os"
	 "testing"
 )
 
 /*
  * ======================================================================
  * Section: String Encoding and Decoding Tests
  * Purpose: Tests the encoding and decoding of string data.
  * ======================================================================
  */
 
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
 
 /*
  * ======================================================================
  * Section: Byte Slice Encoding and Decoding Tests
  * Purpose: Tests the encoding and decoding of byte slice data.
  * ======================================================================
  */
 
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
 
 /*
  * ======================================================================
  * Section: Struct Encoding and Decoding Tests
  * Purpose: Tests the encoding and decoding of struct data.
  * ======================================================================
  */
 
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
 
 /*
  * ======================================================================
  * Section: Slice Encoding and Decoding Tests
  * Purpose: Tests the encoding and decoding of slice data.
  * ======================================================================
  */
 
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
 
 /*
  * ======================================================================
  * Section: File Loading Tests
  * Purpose: Tests the loading of files and handling of file not found errors.
  * ======================================================================
  */
 
 func TestLoadFileNotExists(t *testing.T) {
	 _, err := LoadFile("nonexistentfile.txt")
	 if err == nil {
		 t.Fatal("Expected error for nonexistent file, got nil")
	 }
	 if !errors.Is(err, os.ErrNotExist) && err.Error() != "file does not exist: nonexistentfile.txt" {
		 t.Errorf("Expected specific file not exist error, got: %v", err)
	 }
 }
 
 /*
  * ======================================================================
  * Section: Nil Value Encoding and Decoding Tests
  * Purpose: Tests the handling of nil values during encoding and decoding.
  * ======================================================================
  */
 
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
 
 /*
  * ======================================================================
  * Section: Empty Data Decoding Tests
  * Purpose: Tests the handling of empty data during decoding.
  * ======================================================================
  */
 
 func TestDecodeEmptyData(t *testing.T) {
	 var num int
	 err := Decode([]byte{}, &num)
	 if err == nil || err.Error() != "cannot decode empty data" {
		 t.Errorf("Expected 'cannot decode empty data' error, got: %v", err)
	 }
 }
 