/*
Copyright The Flagger Authors.

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

package v1alpha3

import (
	"context"
	"time"

	v1alpha3 "github.com/weaveworks/flagger/pkg/apis/istio/v1alpha3"
	scheme "github.com/weaveworks/flagger/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EnvoyFiltersGetter has a method to return a EnvoyFilterInterface.
// A group's client should implement this interface.
type EnvoyFiltersGetter interface {
	EnvoyFilters(namespace string) EnvoyFilterInterface
}

// EnvoyFilterInterface has methods to work with EnvoyFilter resources.
type EnvoyFilterInterface interface {
	Create(ctx context.Context, envoyFilter *v1alpha3.EnvoyFilter, opts v1.CreateOptions) (*v1alpha3.EnvoyFilter, error)
	Update(ctx context.Context, envoyFilter *v1alpha3.EnvoyFilter, opts v1.UpdateOptions) (*v1alpha3.EnvoyFilter, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha3.EnvoyFilter, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha3.EnvoyFilterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.EnvoyFilter, err error)
	EnvoyFilterExpansion
}

// envoyFilters implements EnvoyFilterInterface
type envoyFilters struct {
	client rest.Interface
	ns     string
}

// newEnvoyFilters returns a EnvoyFilters
func newEnvoyFilters(c *NetworkingV1alpha3Client, namespace string) *envoyFilters {
	return &envoyFilters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the envoyFilter, and returns the corresponding envoyFilter object, and an error if there is any.
func (c *envoyFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.EnvoyFilter, err error) {
	result = &v1alpha3.EnvoyFilter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("envoyfilters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EnvoyFilters that match those selectors.
func (c *envoyFilters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.EnvoyFilterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha3.EnvoyFilterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("envoyfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested envoyFilters.
func (c *envoyFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("envoyfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a envoyFilter and creates it.  Returns the server's representation of the envoyFilter, and an error, if there is any.
func (c *envoyFilters) Create(ctx context.Context, envoyFilter *v1alpha3.EnvoyFilter, opts v1.CreateOptions) (result *v1alpha3.EnvoyFilter, err error) {
	result = &v1alpha3.EnvoyFilter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("envoyfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(envoyFilter).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a envoyFilter and updates it. Returns the server's representation of the envoyFilter, and an error, if there is any.
func (c *envoyFilters) Update(ctx context.Context, envoyFilter *v1alpha3.EnvoyFilter, opts v1.UpdateOptions) (result *v1alpha3.EnvoyFilter, err error) {
	result = &v1alpha3.EnvoyFilter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("envoyfilters").
		Name(envoyFilter.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(envoyFilter).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the envoyFilter and deletes it. Returns an error if one occurs.
func (c *envoyFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("envoyfilters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *envoyFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("envoyfilters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched envoyFilter.
func (c *envoyFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.EnvoyFilter, err error) {
	result = &v1alpha3.EnvoyFilter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("envoyfilters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
