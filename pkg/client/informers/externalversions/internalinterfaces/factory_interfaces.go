/*
Copyright 2018 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// This file was automatically generated by informer-gen

package internalinterfaces

import (
	versioned "github.com/openfaas-incubator/openfaas-operator/pkg/client/clientset/versioned"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

type NewInformerFunc func(versioned.Interface, time.Duration) cache.SharedIndexInformer

// SharedInformerFactory a small interface to allow for adding an informer without an import cycle
type SharedInformerFactory interface {
	Start(stopCh <-chan struct{})
	InformerFor(obj runtime.Object, newFunc NewInformerFunc) cache.SharedIndexInformer
}
