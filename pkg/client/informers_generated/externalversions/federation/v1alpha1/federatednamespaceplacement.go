/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	federation_v1alpha1 "github.com/marun/fnord/pkg/apis/federation/v1alpha1"
	clientset "github.com/marun/fnord/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/marun/fnord/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/marun/fnord/pkg/client/listers_generated/federation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// FederatedNamespacePlacementInformer provides access to a shared informer and lister for
// FederatedNamespacePlacements.
type FederatedNamespacePlacementInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FederatedNamespacePlacementLister
}

type federatedNamespacePlacementInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewFederatedNamespacePlacementInformer constructs a new informer for FederatedNamespacePlacement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedNamespacePlacementInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedNamespacePlacementInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedNamespacePlacementInformer constructs a new informer for FederatedNamespacePlacement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedNamespacePlacementInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederationV1alpha1().FederatedNamespacePlacements().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederationV1alpha1().FederatedNamespacePlacements().Watch(options)
			},
		},
		&federation_v1alpha1.FederatedNamespacePlacement{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedNamespacePlacementInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedNamespacePlacementInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedNamespacePlacementInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federation_v1alpha1.FederatedNamespacePlacement{}, f.defaultInformer)
}

func (f *federatedNamespacePlacementInformer) Lister() v1alpha1.FederatedNamespacePlacementLister {
	return v1alpha1.NewFederatedNamespacePlacementLister(f.Informer().GetIndexer())
}
