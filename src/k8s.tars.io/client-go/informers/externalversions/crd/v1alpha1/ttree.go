/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	crdv1alpha1 "k8s.tars.io/api/crd/v1alpha1"
	versioned "k8s.tars.io/client-go/clientset/versioned"
	internalinterfaces "k8s.tars.io/client-go/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.tars.io/client-go/listers/crd/v1alpha1"
)

// TTreeInformer provides access to a shared informer and lister for
// TTrees.
type TTreeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TTreeLister
}

type tTreeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTTreeInformer constructs a new informer for TTree type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTTreeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTTreeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTTreeInformer constructs a new informer for TTree type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTTreeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().TTrees(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().TTrees(namespace).Watch(context.TODO(), options)
			},
		},
		&crdv1alpha1.TTree{},
		resyncPeriod,
		indexers,
	)
}

func (f *tTreeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTTreeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tTreeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crdv1alpha1.TTree{}, f.defaultInformer)
}

func (f *tTreeInformer) Lister() v1alpha1.TTreeLister {
	return v1alpha1.NewTTreeLister(f.Informer().GetIndexer())
}