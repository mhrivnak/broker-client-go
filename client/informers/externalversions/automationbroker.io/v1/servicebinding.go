/*
Copyright (c) 2018 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen

package v1

import (
	time "time"

	versioned "github.com/automationbroker/broker-client-go/client/clientset/versioned"
	internalinterfaces "github.com/automationbroker/broker-client-go/client/informers/externalversions/internalinterfaces"
	v1 "github.com/automationbroker/broker-client-go/client/listers/automationbroker.io/v1"
	automationbroker_io_v1 "github.com/automationbroker/broker-client-go/pkg/apis/automationbroker.io/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceBindingInformer provides access to a shared informer and lister for
// ServiceBindings.
type ServiceBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ServiceBindingLister
}

type serviceBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceBindingInformer constructs a new informer for ServiceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceBindingInformer constructs a new informer for ServiceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutomationbrokerV1().ServiceBindings(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutomationbrokerV1().ServiceBindings(namespace).Watch(options)
			},
		},
		&automationbroker_io_v1.ServiceBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&automationbroker_io_v1.ServiceBinding{}, f.defaultInformer)
}

func (f *serviceBindingInformer) Lister() v1.ServiceBindingLister {
	return v1.NewServiceBindingLister(f.Informer().GetIndexer())
}
