/*
 * License: GPL-3.0 license | GNU General Public License v3.0
 * Purpose: This file provides functionality for encoding and decoding various data types using Gob encoding.
 * Key Components:
 * - Encode: Serializes data into []byte.
 * - Decode: Deserializes []byte into data.
 * - LoadFile: Reads a file and returns its contents as []byte.
 * Role within the Codebase: Facilitates the serialization and deserialization of data for storage or transmission.
 */

 package bstream

 import (
	 "bytes"
	 "encoding/gob"
	 "errors"
	 "os"
 )
 
 /*
  * ======================================================================
  * Section: Encoding Function
  * Purpose: Serializes data into []byte using Gob encoding.
  * ======================================================================
  */
 
 // Encode serializes the given value into []byte.
 // If v is []byte or string, it returns the raw bytes.
 // For other types, it uses Gob encoding.
 func Encode(v interface{}) ([]byte, error) {
	 switch val := v.(type) {
	 case []byte:
		 return val, nil
	 case string:
		 return []byte(val), nil
	 case nil:
		 return nil, errors.New("cannot encode nil value")
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
 
 /*
  * ======================================================================
  * Section: Decoding Function
  * Purpose: Deserializes []byte into data using Gob decoding.
  * ======================================================================
  */
 
 // Decode deserializes the given []byte into v.
 // If v is *[]byte or *string, it sets the value directly.
 // For other types, it uses Gob decoding.
 func Decode(data []byte, v interface{}) error {
	 if v == nil {
		 return errors.New("cannot decode into nil value")
	 }
	 switch val := v.(type) {
	 case *[]byte:
		 *val = data
		 return nil
	 case *string:
		 *val = string(data)
		 return nil
	 default:
		 if len(data) == 0 {
			 return errors.New("cannot decode empty data")
		 }
		 buf := bytes.NewBuffer(data)
		 dec := gob.NewDecoder(buf)
		 return dec.Decode(v)
	 }
 }
 
 /*
  * ======================================================================
  * Section: File Loading Function
  * Purpose: Reads a file and returns its contents as []byte.
  * ======================================================================
  */
 
 // LoadFile reads the file at the given path and returns its contents as []byte.
 // Returns an error if the file does not exist or cannot be read.
 func LoadFile(path string) ([]byte, error) {
	 data, err := os.ReadFile(path)
	 if err != nil {
		 if os.IsNotExist(err) {
			 return nil, errors.New("file does not exist: " + path)
		 }
		 return nil, err
	 }
	 return data, nil
 }
 