/*
Copyright The Godel Scheduler Authors.

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

	nodev1alpha1 "github.com/gaodb1210/escheduler-api/pkg/apis/node/v1alpha1"
	versioned "github.com/gaodb1210/escheduler-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/gaodb1210/escheduler-api/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gaodb1210/escheduler-api/pkg/client/listers/node/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VNodeInformer provides access to a shared informer and lister for
// VNodes.
type VNodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VNodeLister
}

type vNodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVNodeInformer constructs a new informer for VNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVNodeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVNodeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVNodeInformer constructs a new informer for VNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVNodeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NodeV1alpha1().VNodes().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NodeV1alpha1().VNodes().Watch(context.TODO(), options)
			},
		},
		&nodev1alpha1.VNode{},
		resyncPeriod,
		indexers,
	)
}

func (f *vNodeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVNodeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vNodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&nodev1alpha1.VNode{}, f.defaultInformer)
}

func (f *vNodeInformer) Lister() v1alpha1.VNodeLister {
	return v1alpha1.NewVNodeLister(f.Informer().GetIndexer())
}
