package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/evankanderson/efgh"
)

// MyData is a JSON-formatted sample structured data from an event.
type MyData struct {
	AField  string `json:"a_field,omitempty"`
	ANumber int64  `json:"a_number,omitempty"`
}

// Exported function.
func HandleEvent(ctx context.Context, data MyData) error {
	cex, ok := efgh.CloudEvent(ctx)
	if !ok {
		log.Printf("Couldn't extract context for event: %+v\n", data)
		return errors.New("Couldn't extract context, failing")
	}
	log.Printf("Read event %s at %s: %+v\n", cex.EventID, cex.EventTime.Format(time.RFC3339), data)
	return nil
}

func main() {
	efgh.Start(HandleEvent)
}
