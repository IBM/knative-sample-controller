# Knative CRD implementation demo

This project demonstrates how to implemented CRDs controllers using the knative Etcd event source linked to
Knative service.

This demo defines one CRD defining cloudant database. Currently no databases is created, just simulated
by printing a message into the console.

## Prerequisites

- [ytt](https://github.com/k14s/ytt)
- [kapp](https://github.com/k14s/kapp)
- [ko](https://github.com/google/ko)
- [kail](hhttps://github.com/boz/kail)

## Installation

1. Clone this repo
1. run these commands:

```shell
ytt t -R -f app/ | ko resolve -f - | kapp deploy -y -a cloudantop -f -
```

## Verifying

1. In one shell do:
```shell
kail -l serving.knative.dev/service=k8s-cloudant-crd-demo-reconcile
```

2. Then apply the sample:
```shell
kubectl apply -f sample/
```

You should see something like this:
```
default/k8s-cloudant-crd-demo-reconcile-w4bjp-deployment-8648c946fg4sfc[user-container]: 2019/04/02 17:48:47 receiving event
default/k8s-cloudant-crd-demo-reconcile-w4bjp-deployment-8648c946fg4sfc[user-container]: 2019/04/02 17:48:47 {{cloudant-credentials knative-cloudant-database}}
```

## Cleaning

1. Delete the sample

```shell
kubectl delete -f sample/
```

2. Delete the cloudant operator:

```shell
kapp delete -a cloudantop
```