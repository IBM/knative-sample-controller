module github.com/lionelvillard/knative-sample-controller

require (
	github.com/cloudevents/sdk-go v0.0.0-20190402205943-58f0318fe886
	github.com/imdario/mergo v0.3.7 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	k8s.io/api v0.0.0-20190424052528-e8c4fd2c9be3
	k8s.io/apimachinery v0.0.0-20190424052434-11f1676e3da4
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.0.0-20190405172246-9a4d48088f6a
	k8s.io/klog v0.3.0
	k8s.io/utils v0.0.0-20190308190857-21c4ce38f2a7 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190424052529-7fd04442e4f5
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190424052434-11f1676e3da4
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190424052710-157c3d454138
)
