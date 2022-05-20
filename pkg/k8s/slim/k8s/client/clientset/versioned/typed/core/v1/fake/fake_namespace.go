// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	corev1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNamespaces implements NamespaceInterface
type FakeNamespaces struct {
	Fake *FakeCoreV1
}

var namespacesResource = schema.GroupVersionResource{Group: "core", Version: "v1", Resource: "namespaces"}

var namespacesKind = schema.GroupVersionKind{Group: "core", Version: "v1", Kind: "Namespace"}

// Get takes name of the namespace, and returns the corresponding namespace object, and an error if there is any.
func (c *FakeNamespaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(namespacesResource, name), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

// List takes label and field selectors, and returns the list of Namespaces that match those selectors.
func (c *FakeNamespaces) List(ctx context.Context, opts v1.ListOptions) (result *corev1.NamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(namespacesResource, namespacesKind, opts), &corev1.NamespaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.NamespaceList{ListMeta: obj.(*corev1.NamespaceList).ListMeta}
	for _, item := range obj.(*corev1.NamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested namespaces.
func (c *FakeNamespaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(namespacesResource, opts))
}

// Create takes the representation of a namespace and creates it.  Returns the server's representation of the namespace, and an error, if there is any.
func (c *FakeNamespaces) Create(ctx context.Context, namespace *corev1.Namespace, opts v1.CreateOptions) (result *corev1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(namespacesResource, namespace), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

// Update takes the representation of a namespace and updates it. Returns the server's representation of the namespace, and an error, if there is any.
func (c *FakeNamespaces) Update(ctx context.Context, namespace *corev1.Namespace, opts v1.UpdateOptions) (result *corev1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(namespacesResource, namespace), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}

// Delete takes name of the namespace and deletes it. Returns an error if one occurs.
func (c *FakeNamespaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(namespacesResource, name, opts), &corev1.Namespace{})
	return err
}

// Patch applies the patch and returns the patched namespace.
func (c *FakeNamespaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(namespacesResource, name, pt, data, subresources...), &corev1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Namespace), err
}
