/*
 * License: GPL-3.0 license | GNU General Public License v3.0
 * Purpose: This file demonstrates the usage of the bstream package for encoding and decoding data.
 * Key Components:
 * - Main function that showcases encoding and decoding of strings, structs, and file data.
 * Role within the Codebase: Provides examples of how to use the bstream package in a real-world scenario.
 */

 package main

 import (
	 "fmt"
 
	 "github.com/cryptrunner49/bytestream/bstream"
 )
 
 /*
  * ======================================================================
  * Section: Main Function
  * Purpose: Demonstrates the encoding and decoding of various data types using the bstream package.
  * ======================================================================
  */
 
 func main() {
	 // Example 1: Encoding and decoding a string
	 str := "Alice"
	 data, err := bstream.Encode(str)
	 if err != nil {
		 fmt.Printf("Error encoding string: %v\n", err)
		 return
	 }
	 fmt.Println("Serialized string:", data)
 
	 var decodedStr string
	 err = bstream.Decode(data, &decodedStr)
	 if err != nil {
		 fmt.Printf("Error decoding string: %v\n", err)
		 return
	 }
	 fmt.Println("Deserialized string:", decodedStr)
 
	 // Example 2: Encoding and decoding a struct
	 type User struct {
		 Name  string
		 Age   int
		 Email string
	 }
	 user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
	 data, err = bstream.Encode(user)
	 if err != nil {
		 fmt.Printf("Error encoding user: %v\n", err)
		 return
	 }
	 fmt.Println("Serialized user:", data)
 
	 var decodedUser User
	 err = bstream.Decode(data, &decodedUser)
	 if err != nil {
		 fmt.Printf("Error decoding user: %v\n", err)
		 return
	 }
	 fmt.Println("Deserialized user:", fmt.Sprintf("%+v", decodedUser))
 
	 // Example 3: Loading and encoding a file
	 fileData, err := bstream.LoadFile("example.txt")
	 if err != nil {
		 fmt.Printf("Error loading file: %v\n", err)
		 return
	 }
	 data, err = bstream.Encode(fileData) // Simply returns fileData as []byte
	 if err != nil {
		 fmt.Printf("Error encoding file data: %v\n", err)
		 return
	 }
	 fmt.Println("Serialized file data:", data)
 
	 var decodedFileData []byte
	 err = bstream.Decode(data, &decodedFileData)
	 if err != nil {
		 fmt.Printf("Error decoding file data: %v\n", err)
		 return
	 }
	 fmt.Println("Deserialized file data:", decodedFileData)
 }
 