package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport/http"
)

// CloudantDatabase is ...
type CloudantDatabase struct {
	Spec CloudantDatabaseSpec `json:"spec"`
}

// CloudantDatabaseSpec is ...
type CloudantDatabaseSpec struct {
	SecretRef string `json:"secretRef"`
	Name      string `json:"name"`
}

func reconcile(ctx context.Context, event cloudevents.Event) {
	log.Println("receiving event")
	valid := event.Validate()
	if valid != nil {
		log.Printf("event not valid: %v\n", valid)
		return
	}

	data := event.Data.([]byte)
	obj := CloudantDatabase{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		log.Printf("data not valid: %v\n", err)
		return
	}

	log.Println(obj)
}

func main() {
	c, err := newDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	log.Fatal(c.StartReceiver(context.Background(), reconcile))
}

func newDefaultClient(target ...string) (client.Client, error) {
	tOpts := []http.Option{http.WithBinaryEncoding()}
	if len(target) > 0 && target[0] != "" {
		tOpts = append(tOpts, http.WithTarget(target[0]))
	}

	// Make an http transport for the CloudEvents client.
	t, err := http.New(tOpts...)
	if err != nil {
		return nil, err
	}
	// Use the transport to make a new CloudEvents client.
	c, err := client.New(t,
		client.WithUUIDs(),
		client.WithTimeNow(),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
