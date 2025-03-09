# DOCS.md

## ByteStream Library

The `bstream` library provides a simple and efficient way to serialize and deserialize data in Go. It supports encoding/decoding strings, byte slices, and complex types (e.g., structs, slices) using Go's `gob` package, as well as loading file contents into byte slices. This library is lightweight and designed for ease of use in projects requiring data serialization or file handling.

### Features

- Encode and decode strings, byte slices, and custom structs.
- Direct handling of raw bytes and strings without additional encoding overhead.
- File loading utility to read file contents into a byte slice.
- Comprehensive error handling for edge cases (e.g., nil values, empty data).

---

## Installation

To use the `bstream` library in your Go project, ensure you have Go installed (version 1.11+ recommended for module support). Then, follow these steps:

1. **Add the library to your project**:
   If the library is hosted in a repository (e.g., `github.com/cryptrunner49/bytestream`), use:

   ```bash
   go get github.com/cryptrunner49/bytestream
   ```

2. **Import the library** in your code:

   ```go
   import "github.com/cryptrunner49/bytestream/bstream"
   ```

3. **Verify installation**:
   Run the sample code in `cmd/main.go` (see Usage Examples below) to confirm the library works as expected.

---

## Usage Examples

Below are examples demonstrating how to use the `bstream` library. These are based on the `cmd/main.go` file and can be adapted to your project.

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
fmt.Println("Deserialized user:", fmt.Sprintf("%+v", decodedUser))
```

**Output:**

```text
Serialized user: [some byte array]
Deserialized user: {Name:Alice Age:30 Email:alice@example.com}
```

### 3. Loading and Encoding File Data

```go
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
```

**Output (assuming `example.txt` contains "Hello"):**

```text
Serialized file data: [72 101 108 108 111]
Deserialized file data: [72 101 108 108 111]
```

---

## API Documentation

The `bstream` package provides the following functions:

### `Encode(v interface{}) ([]byte, error)`

Serializes a value into a byte slice.

- **Parameters**:
  - `v interface{}`: The value to encode. Supported types include:
    - `[]byte`: Returned as-is.
    - `string`: Converted to `[]byte`.
    - Other types: Encoded using `gob`.
- **Returns**:
  - `[]byte`: The serialized data.
  - `error`: Non-nil if encoding fails (e.g., `nil` input).
- **Notes**:
  - Returns an error if `v` is `nil`.

**Example**:

```go
data, err := bstream.Encode("Hello")
if err != nil {
    fmt.Println(err)
}
```

### `Decode(data []byte, v interface{}) error`

Deserializes a byte slice into a provided value.

- **Parameters**:
  - `data []byte`: The serialized data to decode.
  - `v interface{}`: A pointer to the target variable. Supported types include:
    - `*[]byte`: Set directly to `data`.
    - `*string`: Set to `string(data)`.
    - Other types: Decoded using `gob`.
- **Returns**:
  - `error`: Non-nil if decoding fails (e.g., `nil` target, empty `data`).
- **Notes**:
  - `v` must be a pointer to the target type.
  - Returns an error if `v` is `nil` or `data` is empty.

**Example**:

```go
var str string
err := bstream.Decode([]byte("Hello"), &str)
if err != nil {
    fmt.Println(err)
}
fmt.Println(str) // "Hello"
```

### `LoadFile(path string) ([]byte, error)`

Reads a file and returns its contents as a byte slice.

- **Parameters**:
  - `path string`: The path to the file.
- **Returns**:
  - `[]byte`: The file contents.
  - `error`: Non-nil if the file doesn't exist or cannot be read.
- **Notes**:
  - Returns a custom error message for non-existent files.

**Example**:

```go
data, err := bstream.LoadFile("example.txt")
if err != nil {
    fmt.Println(err)
}
fmt.Println(data)
```

---

## Error Handling

The library includes robust error handling:

- `Encode`: Returns `"cannot encode nil value"` for `nil` inputs.
- `Decode`: Returns `"cannot decode into nil value"` or `"cannot decode empty data"` for invalid inputs.
- `LoadFile`: Returns `"file does not exist: <path>"` for missing files, or other OS-specific errors.

---

## Testing

The library includes a comprehensive test suite in `bstream/bstream_test.go`. To run the tests:

```bash
go test ./bstream
```

Tests cover:

- Encoding/decoding strings, byte slices, structs, and slices.
- Edge cases like `nil` values, empty data, and non-existent files.

---

## Project Structure

```text
.
├── cmd
│   └── main.go          # Sample usage of the library
├── bstream
│   ├── bstream.go       # Core library implementation
│   └── bstream_test.go  # Unit tests
└── DOCS.md              # This documentation
```

---

## Notes

- The library uses `encoding/gob` for complex types, so ensure your structs are compatible (e.g., exported fields).
- For raw `[]byte` or `string` data, no additional encoding is applied, making it efficient for simple use cases.
