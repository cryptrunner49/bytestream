# 💨 ByteStream

**Effortless serialization in Go – Convert your data into `[]byte` with ease!**

![Go Badge](https://img.shields.io/badge/Go-1.21-blue?style=flat)
![License](https://img.shields.io/badge/License-MIT-green)

## 🚀 Overview

ByteStream is a lightweight and efficient Go library for serializing and deserializing data into `[]byte`. It provides a simple API for transforming structures, primitives, and more into byte slices, making it ideal for storage, networking, or caching.

## ✨ Features

✅ **Easy-to-use**: Simple API for encoding/decoding.  
✅ **Lightweight & Fast**: Optimized for performance.  
✅ **Flexible**: Supports multiple encoding formats.  
✅ **Go idiomatic**: Designed with Go best practices in mind.  

## 📦 Installation

```sh
go get github.com/cryptrunner49/bytestream
```

## 🛠 Usage

### 🔹 Encoding Data

```go
package main

import (
 "fmt"
 "https://github.com/cryptrunner49/bytestream/bstream"
)

type User struct {
 Name  string
 Age   int
 Email string
}

func main() {
 user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
 data, err := bytestream.Encode(user)
 if err != nil {
  panic(err)
 }

 fmt.Println("Serialized Data:", data)
}
```

### 🔹 Decoding Data

```go
var decodedUser User
err := bytestream.Decode(data, &decodedUser)
if err != nil {
    panic(err)
}

fmt.Println("Decoded User:", decodedUser)
```

## ⚡ Supported Types

- ✅ Structs
- ✅ Maps
- ✅ Slices
- ✅ Strings
- ✅ Integers & Floats
- ✅ Custom types

## 📜 License

ByteStream is licensed under the [GPL-3.0 License](LICENSE).

## 🤝 Contributing

We welcome contributions! Feel free to open issues and submit pull requests.
