package main

import (
	"log"

	"github.com/evankanderson/efgh"
)

// HandleEvent reads generic data, and returns an event with the data reversed.
func HandleEvent(data []byte) []byte {
	log.Printf("Read event: %q. ", data)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	log.Printf("Reversed: %q\n", data)
	return data
}

func main() {
	efgh.Start(HandleEvent)
}
