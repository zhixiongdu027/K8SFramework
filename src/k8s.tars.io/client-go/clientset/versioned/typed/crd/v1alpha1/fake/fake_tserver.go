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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.tars.io/api/crd/v1alpha1"
)

// FakeTServers implements TServerInterface
type FakeTServers struct {
	Fake *FakeCrdV1alpha1
	ns   string
}

var tserversResource = schema.GroupVersionResource{Group: "crd", Version: "v1alpha1", Resource: "tservers"}

var tserversKind = schema.GroupVersionKind{Group: "crd", Version: "v1alpha1", Kind: "TServer"}

// Get takes name of the tServer, and returns the corresponding tServer object, and an error if there is any.
func (c *FakeTServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tserversResource, c.ns, name), &v1alpha1.TServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TServer), err
}

// List takes label and field selectors, and returns the list of TServers that match those selectors.
func (c *FakeTServers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tserversResource, tserversKind, c.ns, opts), &v1alpha1.TServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TServerList{ListMeta: obj.(*v1alpha1.TServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.TServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tServers.
func (c *FakeTServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tserversResource, c.ns, opts))

}

// Create takes the representation of a tServer and creates it.  Returns the server's representation of the tServer, and an error, if there is any.
func (c *FakeTServers) Create(ctx context.Context, tServer *v1alpha1.TServer, opts v1.CreateOptions) (result *v1alpha1.TServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tserversResource, c.ns, tServer), &v1alpha1.TServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TServer), err
}

// Update takes the representation of a tServer and updates it. Returns the server's representation of the tServer, and an error, if there is any.
func (c *FakeTServers) Update(ctx context.Context, tServer *v1alpha1.TServer, opts v1.UpdateOptions) (result *v1alpha1.TServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tserversResource, c.ns, tServer), &v1alpha1.TServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTServers) UpdateStatus(ctx context.Context, tServer *v1alpha1.TServer, opts v1.UpdateOptions) (*v1alpha1.TServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tserversResource, "status", c.ns, tServer), &v1alpha1.TServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TServer), err
}

// Delete takes name of the tServer and deletes it. Returns an error if one occurs.
func (c *FakeTServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tserversResource, c.ns, name), &v1alpha1.TServer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tserversResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TServerList{})
	return err
}

// Patch applies the patch and returns the patched tServer.
func (c *FakeTServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tserversResource, c.ns, name, pt, data, subresources...), &v1alpha1.TServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TServer), err
}