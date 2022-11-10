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
	v1beta3 "k8s.tars.io/apis/tars/v1beta3"
)

// FakeTExitedRecords implements TExitedRecordInterface
type FakeTExitedRecords struct {
	Fake *FakeTarsV1beta3
	ns   string
}

var texitedrecordsResource = schema.GroupVersionResource{Group: "tars.k8s.tars.io", Version: "v1beta3", Resource: "texitedrecords"}

var texitedrecordsKind = schema.GroupVersionKind{Group: "tars.k8s.tars.io", Version: "v1beta3", Kind: "TExitedRecord"}

// Get takes name of the tExitedRecord, and returns the corresponding tExitedRecord object, and an error if there is any.
func (c *FakeTExitedRecords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta3.TExitedRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(texitedrecordsResource, c.ns, name), &v1beta3.TExitedRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TExitedRecord), err
}

// List takes label and field selectors, and returns the list of TExitedRecords that match those selectors.
func (c *FakeTExitedRecords) List(ctx context.Context, opts v1.ListOptions) (result *v1beta3.TExitedRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(texitedrecordsResource, texitedrecordsKind, c.ns, opts), &v1beta3.TExitedRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta3.TExitedRecordList{ListMeta: obj.(*v1beta3.TExitedRecordList).ListMeta}
	for _, item := range obj.(*v1beta3.TExitedRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tExitedRecords.
func (c *FakeTExitedRecords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(texitedrecordsResource, c.ns, opts))

}

// Create takes the representation of a tExitedRecord and creates it.  Returns the server's representation of the tExitedRecord, and an error, if there is any.
func (c *FakeTExitedRecords) Create(ctx context.Context, tExitedRecord *v1beta3.TExitedRecord, opts v1.CreateOptions) (result *v1beta3.TExitedRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(texitedrecordsResource, c.ns, tExitedRecord), &v1beta3.TExitedRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TExitedRecord), err
}

// Update takes the representation of a tExitedRecord and updates it. Returns the server's representation of the tExitedRecord, and an error, if there is any.
func (c *FakeTExitedRecords) Update(ctx context.Context, tExitedRecord *v1beta3.TExitedRecord, opts v1.UpdateOptions) (result *v1beta3.TExitedRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(texitedrecordsResource, c.ns, tExitedRecord), &v1beta3.TExitedRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TExitedRecord), err
}

// Delete takes name of the tExitedRecord and deletes it. Returns an error if one occurs.
func (c *FakeTExitedRecords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(texitedrecordsResource, c.ns, name), &v1beta3.TExitedRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTExitedRecords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(texitedrecordsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta3.TExitedRecordList{})
	return err
}

// Patch applies the patch and returns the patched tExitedRecord.
func (c *FakeTExitedRecords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.TExitedRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(texitedrecordsResource, c.ns, name, pt, data, subresources...), &v1beta3.TExitedRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TExitedRecord), err
}