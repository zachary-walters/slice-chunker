package chunker

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestChunkSlice(t *testing.T) {
	testingStructs1000 := func() []TestStruct {
		s := make([]TestStruct, 1000)
		for i := range 1000 {
			s = append(s, TestStruct{ID: i, Name: "Item " + string(rune(i))})
		}
		return s
	}()

	tests := []struct {
		name         string
		input        []TestStruct
		maxChunkSize int
		expected     int
	}{
		{
			name: "Single chunk",
			input: []TestStruct{
				{ID: 1, Name: "Item 1"},
				{ID: 2, Name: "Item 2"},
			},
			maxChunkSize: 1024,
			expected:     1,
		},
		{
			name:         "Multiple chunks",
			input:        testingStructs1000,
			maxChunkSize: 50,
			expected:     1470,
		},
		{
			name:         "Empty input",
			input:        []TestStruct{},
			maxChunkSize: 50,
			expected:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chunks, err := ChunkSlice(tt.input, tt.maxChunkSize)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, len(chunks))

			// Verify that the total number of items is the same
			totalItems := 0
			for _, chunk := range chunks {
				totalItems += len(chunk)
			}
			assert.Equal(t, len(tt.input), totalItems)

			// Verify that each chunk size is within the limit
			for _, chunk := range chunks {
				chunkSize := 0
				for _, item := range chunk {
					itemSize, err := json.Marshal(item)
					assert.NoError(t, err)
					chunkSize += len(itemSize)
				}
				assert.LessOrEqual(t, chunkSize, tt.maxChunkSize)
			}
		})
	}
}
