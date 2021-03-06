// slabs-forge

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	slabsforgev1 "github.com/slabs-forge/klabs-api/apis/slabs.forge/v1"
	versioned "github.com/slabs-forge/klabs-api/clientset/versioned"
	internalinterfaces "github.com/slabs-forge/klabs-api/informers/externalversions/internalinterfaces"
	v1 "github.com/slabs-forge/klabs-api/listers/slabs.forge/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ModuleInformer provides access to a shared informer and lister for
// Modules.
type ModuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ModuleLister
}

type moduleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewModuleInformer constructs a new informer for Module type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewModuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredModuleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredModuleInformer constructs a new informer for Module type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredModuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SlabsV1().Modules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SlabsV1().Modules(namespace).Watch(context.TODO(), options)
			},
		},
		&slabsforgev1.Module{},
		resyncPeriod,
		indexers,
	)
}

func (f *moduleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredModuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *moduleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&slabsforgev1.Module{}, f.defaultInformer)
}

func (f *moduleInformer) Lister() v1.ModuleLister {
	return v1.NewModuleLister(f.Informer().GetIndexer())
}
