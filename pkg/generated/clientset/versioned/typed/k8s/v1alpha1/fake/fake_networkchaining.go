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
	v1alpha1 "github.com/akraino-edge-stack/icn-nodus/pkg/apis/k8s/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkChainings implements NetworkChainingInterface
type FakeNetworkChainings struct {
	Fake *FakeK8sV1alpha1
	ns   string
}

var networkchainingsResource = schema.GroupVersionResource{Group: "k8s.plugin.opnfv.org", Version: "v1alpha1", Resource: "networkchainings"}

var networkchainingsKind = schema.GroupVersionKind{Group: "k8s.plugin.opnfv.org", Version: "v1alpha1", Kind: "NetworkChaining"}

// Get takes name of the networkChaining, and returns the corresponding networkChaining object, and an error if there is any.
func (c *FakeNetworkChainings) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkChaining, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkchainingsResource, c.ns, name), &v1alpha1.NetworkChaining{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkChaining), err
}

// List takes label and field selectors, and returns the list of NetworkChainings that match those selectors.
func (c *FakeNetworkChainings) List(opts v1.ListOptions) (result *v1alpha1.NetworkChainingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkchainingsResource, networkchainingsKind, c.ns, opts), &v1alpha1.NetworkChainingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkChainingList{ListMeta: obj.(*v1alpha1.NetworkChainingList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkChainingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkChainings.
func (c *FakeNetworkChainings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkchainingsResource, c.ns, opts))

}

// Create takes the representation of a networkChaining and creates it.  Returns the server's representation of the networkChaining, and an error, if there is any.
func (c *FakeNetworkChainings) Create(networkChaining *v1alpha1.NetworkChaining) (result *v1alpha1.NetworkChaining, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkchainingsResource, c.ns, networkChaining), &v1alpha1.NetworkChaining{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkChaining), err
}

// Update takes the representation of a networkChaining and updates it. Returns the server's representation of the networkChaining, and an error, if there is any.
func (c *FakeNetworkChainings) Update(networkChaining *v1alpha1.NetworkChaining) (result *v1alpha1.NetworkChaining, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkchainingsResource, c.ns, networkChaining), &v1alpha1.NetworkChaining{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkChaining), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkChainings) UpdateStatus(networkChaining *v1alpha1.NetworkChaining) (*v1alpha1.NetworkChaining, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkchainingsResource, "status", c.ns, networkChaining), &v1alpha1.NetworkChaining{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkChaining), err
}

// Delete takes name of the networkChaining and deletes it. Returns an error if one occurs.
func (c *FakeNetworkChainings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkchainingsResource, c.ns, name), &v1alpha1.NetworkChaining{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkChainings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkchainingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkChainingList{})
	return err
}

// Patch applies the patch and returns the patched networkChaining.
func (c *FakeNetworkChainings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkChaining, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkchainingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkChaining{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkChaining), err
}
