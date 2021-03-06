// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	"time"

	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumBGPLoadBalancerIPPoolsGetter has a method to return a CiliumBGPLoadBalancerIPPoolInterface.
// A group's client should implement this interface.
type CiliumBGPLoadBalancerIPPoolsGetter interface {
	CiliumBGPLoadBalancerIPPools() CiliumBGPLoadBalancerIPPoolInterface
}

// CiliumBGPLoadBalancerIPPoolInterface has methods to work with CiliumBGPLoadBalancerIPPool resources.
type CiliumBGPLoadBalancerIPPoolInterface interface {
	Create(ctx context.Context, ciliumBGPLoadBalancerIPPool *v2alpha1.CiliumBGPLoadBalancerIPPool, opts v1.CreateOptions) (*v2alpha1.CiliumBGPLoadBalancerIPPool, error)
	Update(ctx context.Context, ciliumBGPLoadBalancerIPPool *v2alpha1.CiliumBGPLoadBalancerIPPool, opts v1.UpdateOptions) (*v2alpha1.CiliumBGPLoadBalancerIPPool, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.CiliumBGPLoadBalancerIPPool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.CiliumBGPLoadBalancerIPPoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumBGPLoadBalancerIPPool, err error)
	CiliumBGPLoadBalancerIPPoolExpansion
}

// ciliumBGPLoadBalancerIPPools implements CiliumBGPLoadBalancerIPPoolInterface
type ciliumBGPLoadBalancerIPPools struct {
	client rest.Interface
}

// newCiliumBGPLoadBalancerIPPools returns a CiliumBGPLoadBalancerIPPools
func newCiliumBGPLoadBalancerIPPools(c *CiliumV2alpha1Client) *ciliumBGPLoadBalancerIPPools {
	return &ciliumBGPLoadBalancerIPPools{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumBGPLoadBalancerIPPool, and returns the corresponding ciliumBGPLoadBalancerIPPool object, and an error if there is any.
func (c *ciliumBGPLoadBalancerIPPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumBGPLoadBalancerIPPool, err error) {
	result = &v2alpha1.CiliumBGPLoadBalancerIPPool{}
	err = c.client.Get().
		Resource("ciliumbgploadbalancerippools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumBGPLoadBalancerIPPools that match those selectors.
func (c *ciliumBGPLoadBalancerIPPools) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumBGPLoadBalancerIPPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.CiliumBGPLoadBalancerIPPoolList{}
	err = c.client.Get().
		Resource("ciliumbgploadbalancerippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumBGPLoadBalancerIPPools.
func (c *ciliumBGPLoadBalancerIPPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumbgploadbalancerippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciliumBGPLoadBalancerIPPool and creates it.  Returns the server's representation of the ciliumBGPLoadBalancerIPPool, and an error, if there is any.
func (c *ciliumBGPLoadBalancerIPPools) Create(ctx context.Context, ciliumBGPLoadBalancerIPPool *v2alpha1.CiliumBGPLoadBalancerIPPool, opts v1.CreateOptions) (result *v2alpha1.CiliumBGPLoadBalancerIPPool, err error) {
	result = &v2alpha1.CiliumBGPLoadBalancerIPPool{}
	err = c.client.Post().
		Resource("ciliumbgploadbalancerippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumBGPLoadBalancerIPPool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciliumBGPLoadBalancerIPPool and updates it. Returns the server's representation of the ciliumBGPLoadBalancerIPPool, and an error, if there is any.
func (c *ciliumBGPLoadBalancerIPPools) Update(ctx context.Context, ciliumBGPLoadBalancerIPPool *v2alpha1.CiliumBGPLoadBalancerIPPool, opts v1.UpdateOptions) (result *v2alpha1.CiliumBGPLoadBalancerIPPool, err error) {
	result = &v2alpha1.CiliumBGPLoadBalancerIPPool{}
	err = c.client.Put().
		Resource("ciliumbgploadbalancerippools").
		Name(ciliumBGPLoadBalancerIPPool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumBGPLoadBalancerIPPool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciliumBGPLoadBalancerIPPool and deletes it. Returns an error if one occurs.
func (c *ciliumBGPLoadBalancerIPPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumbgploadbalancerippools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumBGPLoadBalancerIPPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumbgploadbalancerippools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciliumBGPLoadBalancerIPPool.
func (c *ciliumBGPLoadBalancerIPPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumBGPLoadBalancerIPPool, err error) {
	result = &v2alpha1.CiliumBGPLoadBalancerIPPool{}
	err = c.client.Patch(pt).
		Resource("ciliumbgploadbalancerippools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
