// slabs-forge

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/slabs-forge/klabs-api/apis/slabs.forge/v1"
	scheme "github.com/slabs-forge/klabs-api/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ModuleLinksGetter has a method to return a ModuleLinkInterface.
// A group's client should implement this interface.
type ModuleLinksGetter interface {
	ModuleLinks(namespace string) ModuleLinkInterface
}

// ModuleLinkInterface has methods to work with ModuleLink resources.
type ModuleLinkInterface interface {
	Create(ctx context.Context, moduleLink *v1.ModuleLink, opts metav1.CreateOptions) (*v1.ModuleLink, error)
	Update(ctx context.Context, moduleLink *v1.ModuleLink, opts metav1.UpdateOptions) (*v1.ModuleLink, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ModuleLink, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ModuleLinkList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ModuleLink, err error)
	ModuleLinkExpansion
}

// moduleLinks implements ModuleLinkInterface
type moduleLinks struct {
	client rest.Interface
	ns     string
}

// newModuleLinks returns a ModuleLinks
func newModuleLinks(c *SlabsV1Client, namespace string) *moduleLinks {
	return &moduleLinks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the moduleLink, and returns the corresponding moduleLink object, and an error if there is any.
func (c *moduleLinks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ModuleLink, err error) {
	result = &v1.ModuleLink{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("modulelinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ModuleLinks that match those selectors.
func (c *moduleLinks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ModuleLinkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ModuleLinkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("modulelinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested moduleLinks.
func (c *moduleLinks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("modulelinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a moduleLink and creates it.  Returns the server's representation of the moduleLink, and an error, if there is any.
func (c *moduleLinks) Create(ctx context.Context, moduleLink *v1.ModuleLink, opts metav1.CreateOptions) (result *v1.ModuleLink, err error) {
	result = &v1.ModuleLink{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("modulelinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(moduleLink).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a moduleLink and updates it. Returns the server's representation of the moduleLink, and an error, if there is any.
func (c *moduleLinks) Update(ctx context.Context, moduleLink *v1.ModuleLink, opts metav1.UpdateOptions) (result *v1.ModuleLink, err error) {
	result = &v1.ModuleLink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("modulelinks").
		Name(moduleLink.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(moduleLink).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the moduleLink and deletes it. Returns an error if one occurs.
func (c *moduleLinks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("modulelinks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *moduleLinks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("modulelinks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched moduleLink.
func (c *moduleLinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ModuleLink, err error) {
	result = &v1.ModuleLink{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("modulelinks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
