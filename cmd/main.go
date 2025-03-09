package main

import (
	"fmt"
	"github.com/cryptrunner49/bytestream/bstream" // Replace with your actual module path
)

func main() {
	// Example 1: Encoding and decoding a string
	str := "Alice"
	data, err := bstream.Encode(str)
	if err != nil {
		panic(err)
	}
	fmt.Println("Serialized string:", data)

	var decodedStr string
	err = bstream.Decode(data, &decodedStr)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	fmt.Println("Serialized user:", data)

	var decodedUser User
	err = bstream.Decode(data, &decodedUser)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deserialized user:", fmt.Sprintf("%+v", decodedUser))

	// Example 3: Loading and encoding a file
	fileData, err := bstream.LoadFile("example.txt")
	if err != nil {
		panic(err)
	}
	data, err = bstream.Encode(fileData) // Simply returns fileData as []byte
	if err != nil {
		panic(err)
	}
	fmt.Println("Serialized file data:", data)

	var decodedFileData []byte
	err = bstream.Decode(data, &decodedFileData)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deserialized file data:", decodedFileData)
}