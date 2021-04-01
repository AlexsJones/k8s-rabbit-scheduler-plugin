module github.com/AlexsJones/k8s-rabbit-scheduler-plugin

go 1.14

require (
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.0 // indirect
	github.com/onsi/ginkgo v1.11.0 // indirect
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/prometheus/client_golang v0.9.3 // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/sys v0.0.0-20190922100055-0a153f010e69 // indirect
	golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.16.12
	k8s.io/apimachinery v0.16.12
	k8s.io/client-go v0.16.12
	k8s.io/klog v1.0.0
	k8s.io/kube-scheduler v0.0.0
	k8s.io/kubernetes v1.16.12
)

replace (
	k8s.io/api => k8s.io/api v0.16.12
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.12
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.12
	k8s.io/apiserver => k8s.io/apiserver v0.16.12
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.16.12
	k8s.io/client-go => k8s.io/client-go v0.16.12
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.16.12
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.16.12
	k8s.io/code-generator => k8s.io/code-generator v0.16.12
	k8s.io/component-base => k8s.io/component-base v0.16.12
	k8s.io/cri-api => k8s.io/cri-api v0.16.12
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.16.12
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.16.12
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.16.12
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.16.12
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.16.12
	k8s.io/kubectl => k8s.io/kubectl v0.16.12
	k8s.io/kubelet => k8s.io/kubelet v0.16.12
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.16.12
	k8s.io/metrics => k8s.io/metrics v0.16.12
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.16.12
)
