package main

import (
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apiserver/pkg/storage"
	_ "k8s.io/client-go/kubernetes"

	_ "k8s.io/kube-aggregator/pkg/apis/apiregistration/install"
	_ "k8s.io/kube-aggregator/pkg/apis/apiregistration/validation"
	_ "k8s.io/kube-aggregator/pkg/client/clientset_generated/internalclientset"
	_ "k8s.io/kube-aggregator/pkg/client/listers/apiregistration/internalversion"
	_ "k8s.io/kube-aggregator/pkg/client/listers/apiregistration/v1beta1"
	_ "k8s.io/kube-aggregator/pkg/cmd/server"
)

func main() {
}
