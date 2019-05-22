package main

import (
	"context"
	"flag"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"

	"github.com/lionelvillard/knative-sample-controller/pkg/controller"
	clientset "github.com/lionelvillard/knative-sample-controller/pkg/generated/clientset/versioned"
	"github.com/lionelvillard/knative-sample-controller/pkg/informer"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	flag.Parse()

	ctx := context.Background()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}

	// Create Cloud Event informer. It creates an HTTP server listening to cloud events on port 8080
	// and put events into the workqueue
	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Foos")
	cloudEventInformer, err := informer.NewInformer(queue)
	if err != nil {
		klog.Fatalf("Error creating Cloud Event Informer: %s", err.Error())
	}

	// Start this informer in a go routine
	go cloudEventInformer.Start(ctx)

	// Create the controller for Foo.
	ctrl := controller.NewController(kubeClient, exampleClient, queue)

	// And run it.
	if err = ctrl.Run(2, ctx.Done()); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
