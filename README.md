# chunker

The `chunker` package provides a utility function to chunk a slice of any type into smaller chunks based on the serialized size in bytes. This is useful when you need to process or transmit large slices in smaller, manageable pieces.

## Installation

To install the package, use `go get`:

```sh
go get github.com/zachary-walters/slice-chunker
```

## Usage

Import the package into your Go code

```sh
import "github.com/zachary-walters/slice-chunker"
```

## Function 

### ChunkSlice

The `ChunkSlice` function takes a slice of any type and chunks it by `maxChunkSize` in bytes.

```go
func ChunkSlice[T any](items []T, maxChunkSize int) ([][]T, error)
```

#### Parameters
- `items`: The input slice of items to be chunked.
- `maxChunkSize`: The maximum size of each chunk in bytes.

#### Returns
- A slice of chunks, where each chunk is a slice of items.
- An error if any occurred during the serialization of items.

#### Example

You can find the example [here](example/example.go).

```go 
package main

import (
    "fmt"
    "log"
    chunker "github.com/zachary-walters/slice-chunker"
)

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    people := []Person{
        {FirstName: "Laura", LastName: "Palmer", Age: 17},
		{FirstName: "Bobby", LastName: "Briggs", Age: 17},
		{FirstName: "Dale", LastName: "Cooper", Age: 17},
        // Add more people as needed
    }

    chunks, err := chunker.ChunkSlice(people, 500*1024)
    if err != nil {
        log.Fatal(err)
    }

    for i, chunk := range chunks {
        fmt.Printf("Chunk %d: %v\n", i+1, chunk)
    }
}
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.