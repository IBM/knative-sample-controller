# Knative Sample Controller

This project demonstrates how to implement CustomResourceDefinition (CRD) controllers using the
Knative [API server event source](https://github.com/knative/eventing-sources/tree/master/contrib/apiserver) linked to Knative service.

This is a Knative variant of the [sample controller](https://github.com/kubernetes/sample-controller) project.

## Prerequisites

- A Kubernetes cluster
- [Knative API Server eventing source](https://github.com/knative/eventing-sources/tree/master/contrib/apiserver/samples) installed in your cluster
- envsubst installed locally. This is installed by the gettext package. If not installed it can be installed by a Linux package manager, or by Homebrew on OS X.

## Running

1. Clone this repository somewhere not under $GOPATH
1. Set $DOCKER_REGISTRY and $DOCKER_USER
1. Build and publish the sample-controller docker image

```sh
go get ./...
CGO_ENABLED=0 GOOS=linux go build -o sample-controller cmd/reconcile/main.go
docker build -t $DOCKER_USER/sample-controller .
docker push $DOCKER_USER/sample-controller
```

4. Deploy:

```sh
envsubst < config/template/ksvc-example.yaml | kubectl apply -f -
kubectl apply -f config/default
```

## Verifying

1. Apply the sample:
```sh
kubectl apply -f sample
```

2. Wait a bit and do:
```sh
kubectl get deployments example-foo
NAME          DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
example-foo   1         1         1            1           4m48s
```


## Cleaning

1. Delete the sample

```shell
kubectl delete -f sample/
```

2. Delete the controller:

```she
kubectl delete -R -f config
```