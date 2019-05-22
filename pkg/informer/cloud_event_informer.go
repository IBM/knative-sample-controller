package informer

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport/http"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/util/workqueue"
)

type Informer interface {
	// Start enqueue cloud events
	// This call is blocking.
	Start(context context.Context)
}

type informer struct {
	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface

	client cloudevents.Client
}

// NewInformer returns a controller for populating the store while also
// providing event notifications.
//
// Parameters
//  * objType is an object of the type that you expect to receive.
//  * resyncPeriod: if non-zero, will re-list this often (you will get OnUpdate
//    calls, even if nothing changed). Otherwise, re-list will be delayed as
//    long as possible (until the upstream source closes the watch or times out,
//    or you stop the controller).
//  * h is the object you want notifications sent to.
//  * clientState is the store you want to populate
//
func NewInformer(workqueue workqueue.RateLimitingInterface) (Informer, error) {
	tOpts := []http.Option{cloudevents.WithBinaryEncoding()}

	// Make an http transport for the CloudEvents client.
	t, err := cloudevents.NewHTTPTransport(tOpts...)
	if err != nil {
		return nil, fmt.Errorf("Error creating HTTP transport: %v", err)
	}

	// Use the transport to make a new CloudEvents client.
	c, err := cloudevents.NewClient(t,
		cloudevents.WithUUIDs(),
		cloudevents.WithTimeNow(),
	)
	if err != nil {
		return nil, fmt.Errorf("Error creating CloudEvent client: %v", err)
	}

	return &informer{
		workqueue: workqueue,
		client:    c,
	}, nil
}

func (c *informer) Start(ctx context.Context) {
	c.client.StartReceiver(ctx, c.enqueue)
}

func (c *informer) enqueue(event cloudevents.Event) error {
	key := &corev1.ObjectReference{}
	err := event.DataAs(key)
	if err != nil {
		return err
	}
	c.workqueue.Add(key.Namespace + "/" + key.Name)
	return nil
}
