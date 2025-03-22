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

    chunks, err := chunker.ChunkSlice(people, 100)
    if err != nil {
        log.Fatal(err)
    }

    for i, chunk := range chunks {
        fmt.Printf("Chunk %d: %v\n", i+1, chunk)
    }
}