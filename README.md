# Serverless Operator running in Knative

This project demonstrates how to implement CustomResourceDefinition (CRD) controllers using the
Knative [API server event source](https://github.com/knative/eventing) feeding Knative services.

This is a Knative variant of the [sample controller](https://github.com/kubernetes/sample-controller) project.

## Prerequisites

- A Kubernetes cluster with [Knative 0.6.0](https://knative.dev)
- envsubst installed locally. This is installed by the gettext package. If not installed it can be installed by a Linux package manager, or by Homebrew on OS X.

## Running

1. Clone this repository somewhere not under $GOPATH
1. Set $DOCKER_USER
1. Build and publish the sample-controller docker image

```sh
go get ./...
CGO_ENABLED=0 GOOS=linux go build -o sample-controller cmd/reconcile/main.go
docker build -t $DOCKER_USER/sample-controller .
docker push $DOCKER_USER/sample-controller
```

4. Deploy the sample controller in the default namespace:

```sh
envsubst < config/template/ksvc-example.yaml | kubectl apply -f -
kubectl apply -Rf config/default
```

Two pods are created: one watching for events (`apiserver-example-foo-XXX`) and another one containing the reconciling loop (`example-foo-reconcile-XXX-deployment-YYY-ZZZ`). The first pod always stays alive, whereas the second one scale down to zero when there is no events.


## Verifying

1. In one terminal, watch for pods.
```sh
watch kubectl get pods
```

2. In another terminal, create a `Foo` object:
```sh
kubectl apply -f sample
```

Observe the pod named `example-foo-reconcile-XXX` being created.

3. Check the controller for `Foo` is working:
```sh
kubectl get deployments example-foo
NAME          DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
example-foo   1         1         1            1           4m48s
```

Observe the pod named `example-foo-reconcile-XXX` being deleted after about 1mn30s (default Knative scale down period).

## Cleaning

1. Delete the sample

```shell
kubectl delete -f sample/
```

2. Delete the controller:

```she
kubectl delete -Rf config
```
