# ByteStream Library Documentation

The `bstream` library provides a lightweight and efficient solution for serializing and deserializing data in Go. Built on Go's `gob` package, it supports encoding and decoding of strings, byte slices, and complex types like structs and slices. Additionally, it offers a utility for loading file contents into byte slices. This library is designed for simplicity and performance, making it ideal for projects requiring data serialization or file handling.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage Examples](#usage-examples)
  - [Encoding and Decoding a String](#1-encoding-and-decoding-a-string)
  - [Encoding and Decoding a Struct](#2-encoding-and-decoding-a-struct)
  - [Loading and Encoding File Data](#3-loading-and-encoding-file-data)
- [API Documentation](#api-documentation)
  - [`Encode(v interface{}) ([]byte, error)`](#encodev-interface-byte-error)
  - [`Decode(data []byte, v interface{}) error`](#decodedata-byte-v-interface-error)
  - [`LoadFile(path string) ([]byte, error)`](#loadfilepath-string-byte-error)
- [Error Handling](#error-handling)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Notes](#notes)

---

## Features

- **Serialization/Deserialization**: Encode and decode strings, byte slices, and custom structs with ease.
- **Raw Data Handling**: Process raw bytes and strings directly, avoiding unnecessary encoding overhead.
- **File Loading**: Read file contents into byte slices with a simple utility function.
- **Robust Error Handling**: Comprehensive checks for edge cases like nil values or empty data.

---

## Installation

To integrate the `bstream` library into your Go project, ensure you have Go 1.11 or later (for module support). Follow these steps:

1. **Add the Library**:
   If hosted at `github.com/cryptrunner49/bytestream`, run:

   ```bash
   go get github.com/cryptrunner49/bytestream
   ```

2. **Import in Your Code**:

   ```go
   import "github.com/cryptrunner49/bytestream/bstream"
   ```

3. **Verify Installation**:
   Test the library by running the sample code in `cmd/main.go` (see [Usage Examples](#usage-examples)).

---

## Usage Examples

Below are practical examples showcasing the library's functionality, derived from `cmd/main.go`.

### 1. Encoding and Decoding a String

```go
package main

import (
    "fmt"
    "github.com/cryptrunner49/bytestream/bstream"
)

func main() {
    // Encode a string
    str := "Alice"
    data, err := bstream.Encode(str)
    if err != nil {
        fmt.Printf("Error encoding string: %v\n", err)
        return
    }
    fmt.Println("Serialized string:", data)

    // Decode back to a string
    var decodedStr string
    err = bstream.Decode(data, &decodedStr)
    if err != nil {
        fmt.Printf("Error decoding string: %v\n", err)
        return
    }
    fmt.Println("Deserialized string:", decodedStr)
}
```

**Output:**

```text
Serialized string: [65 108 105 99 101]
Deserialized string: Alice
```

### 2. Encoding and Decoding a Struct

```go
type User struct {
    Name  string
    Age   int
    Email string
}

func main() {
    user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
    data, err := bstream.Encode(user)
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
    fmt.Printf("Deserialized user: %+v\n", decodedUser)
}
```

**Output:**

```text
Serialized user: [some byte array]
Deserialized user: {Name:Alice Age:30 Email:alice@example.com}
```

### 3. Loading and Encoding File Data

```go
func main() {
    fileData, err := bstream.LoadFile("example.txt")
    if err != nil {
        fmt.Printf("Error loading file: %v\n", err)
        return
    }
    data, err := bstream.Encode(fileData)
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
```

**Output (assuming `example.txt` contains "Hello"):**

```text
Serialized file data: [72 101 108 108 111]
Deserialized file data: [72 101 108 108 111]
```

---

## API Documentation

The `bstream` package exposes the following functions:

### `Encode(v interface{}) ([]byte, error)`

Serializes a value into a byte slice.

- **Parameters**:
  - `v interface{}`: The value to encode.
    - `[]byte`: Returned as-is.
    - `string`: Converted to `[]byte`.
    - Other types: Encoded with `gob`.
- **Returns**:
  - `[]byte`: Serialized data.
  - `error`: Non-nil if encoding fails (e.g., `nil` input).
- **Example**:

  ```go
  data, err := bstream.Encode("Hello")
  if err != nil {
      fmt.Println(err)
  }
  ```

### `Decode(data []byte, v interface{}) error`

Deserializes a byte slice into a target value.

- **Parameters**:
  - `data []byte`: The serialized data.
  - `v interface{}`: Pointer to the target variable.
    - `*[]byte`: Set to `data`.
    - `*string`: Set to `string(data)`.
    - Other types: Decoded with `gob`.
- **Returns**:
  - `error`: Non-nil if decoding fails (e.g., `nil` target, empty `data`).
- **Notes**:
  - `v` must be a pointer.
- **Example**:

  ```go
  var str string
  err := bstream.Decode([]byte("Hello"), &str)
  if err != nil {
      fmt.Println(err)
  }
  fmt.Println(str) // "Hello"
  ```

### `LoadFile(path string) ([]byte, error)`

Reads a file into a byte slice.

- **Parameters**:
  - `path string`: File path.
- **Returns**:
  - `[]byte`: File contents.
  - `error`: Non-nil if the file is missing or unreadable.
- **Example**:

  ```go
  data, err := bstream.LoadFile("example.txt")
  if err != nil {
      fmt.Println(err)
  }
  fmt.Println(data)
  ```

---

## Error Handling

The library provides detailed error messages:

- **`Encode`**: `"cannot encode nil value"` for `nil` inputs.
- **`Decode`**:
  - `"cannot decode into nil value"` for `nil` targets.
  - `"cannot decode empty data"` for empty `data`.
- **`LoadFile`**: `"file does not exist: <path>"` for missing files, or OS-specific errors.

---

## Testing

Run the test suite in `bstream/bstream_test.go` with:

```bash
go test ./bstream
```

Tests cover:

- Encoding/decoding of strings, byte slices, structs, and slices.
- Edge cases (e.g., `nil` values, empty data, missing files).

---

## Project Structure

```text
.
├── cmd
│   └── main.go          # Sample usage
├── bstream
│   ├── bstream.go       # Core implementation
│   └── bstream_test.go  # Unit tests
└── DOCS.md              # This file
```

---

## Notes

- Uses `encoding/gob` for complex types; ensure struct fields are exported (capitalized).
- Raw `[]byte` and `string` data are processed without additional encoding for efficiency.
- Check error returns to handle edge cases gracefully.
