package cloud

import (
	"context"
	"fmt"
	"net/http"
)

// F My Coud Hello World - Http
func F(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello, Crystal Cloud️️ ☁ Factory!")
	w.Write(msg)
}

// G - fn on event
func G(ctx context.Context, e GCSEvent) error {
	fmt.Printf("%s was uploaded to %s\n", e.Name, e.Bucket)
	return nil
}

// GCSEvent type
type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}
