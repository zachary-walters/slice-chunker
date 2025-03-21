package chunker

import "encoding/json"

// ChunkSlice takes a slice of any type and chunks it by maxChunkSize in bytes.
//
// The function iterates through the input slice and groups the items into chunks
// based on their serialized size in bytes. Each chunk is ensured to be less than
// or equal to maxChunkSize. The function returns a slice of chunks, where each
// chunk is a slice of items.
//
// Parameters:
//
//	items - the input slice of items to be chunked
//	maxChunkSize - the maximum size of each chunk in bytes
//
// Returns:
//
//	A slice of chunks, where each chunk is a slice of items, and an error if any
//	occurred during the serialization of items.
func ChunkSlice[T any](items []T, maxChunkSize int) ([][]T, error) {
	var chunks [][]T
	var currentChunk []T
	var currentChunkSize int

	for _, item := range items {
		itemSize, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		if currentChunkSize+len(itemSize) > maxChunkSize {
			chunks = append(chunks, currentChunk)
			currentChunk = []T{}
			currentChunkSize = 0
		}

		currentChunk = append(currentChunk, item)
		currentChunkSize += len(itemSize)
	}

	if len(currentChunk) > 0 {
		chunks = append(chunks, currentChunk)
	}

	return chunks, nil
}
