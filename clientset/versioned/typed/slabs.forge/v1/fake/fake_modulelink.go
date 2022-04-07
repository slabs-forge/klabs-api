// slabs-forge

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	slabsforgev1 "github.com/slabs-forge/klabs-api/apis/slabs.forge/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeModuleLinks implements ModuleLinkInterface
type FakeModuleLinks struct {
	Fake *FakeSlabsV1
	ns   string
}

var modulelinksResource = schema.GroupVersionResource{Group: "slabs.forge", Version: "v1", Resource: "modulelinks"}

var modulelinksKind = schema.GroupVersionKind{Group: "slabs.forge", Version: "v1", Kind: "ModuleLink"}

// Get takes name of the moduleLink, and returns the corresponding moduleLink object, and an error if there is any.
func (c *FakeModuleLinks) Get(ctx context.Context, name string, options v1.GetOptions) (result *slabsforgev1.ModuleLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(modulelinksResource, c.ns, name), &slabsforgev1.ModuleLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*slabsforgev1.ModuleLink), err
}

// List takes label and field selectors, and returns the list of ModuleLinks that match those selectors.
func (c *FakeModuleLinks) List(ctx context.Context, opts v1.ListOptions) (result *slabsforgev1.ModuleLinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(modulelinksResource, modulelinksKind, c.ns, opts), &slabsforgev1.ModuleLinkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &slabsforgev1.ModuleLinkList{ListMeta: obj.(*slabsforgev1.ModuleLinkList).ListMeta}
	for _, item := range obj.(*slabsforgev1.ModuleLinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested moduleLinks.
func (c *FakeModuleLinks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(modulelinksResource, c.ns, opts))

}

// Create takes the representation of a moduleLink and creates it.  Returns the server's representation of the moduleLink, and an error, if there is any.
func (c *FakeModuleLinks) Create(ctx context.Context, moduleLink *slabsforgev1.ModuleLink, opts v1.CreateOptions) (result *slabsforgev1.ModuleLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(modulelinksResource, c.ns, moduleLink), &slabsforgev1.ModuleLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*slabsforgev1.ModuleLink), err
}

// Update takes the representation of a moduleLink and updates it. Returns the server's representation of the moduleLink, and an error, if there is any.
func (c *FakeModuleLinks) Update(ctx context.Context, moduleLink *slabsforgev1.ModuleLink, opts v1.UpdateOptions) (result *slabsforgev1.ModuleLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(modulelinksResource, c.ns, moduleLink), &slabsforgev1.ModuleLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*slabsforgev1.ModuleLink), err
}

// Delete takes name of the moduleLink and deletes it. Returns an error if one occurs.
func (c *FakeModuleLinks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(modulelinksResource, c.ns, name, opts), &slabsforgev1.ModuleLink{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeModuleLinks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(modulelinksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &slabsforgev1.ModuleLinkList{})
	return err
}

// Patch applies the patch and returns the patched moduleLink.
func (c *FakeModuleLinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *slabsforgev1.ModuleLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(modulelinksResource, c.ns, name, pt, data, subresources...), &slabsforgev1.ModuleLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*slabsforgev1.ModuleLink), err
}
