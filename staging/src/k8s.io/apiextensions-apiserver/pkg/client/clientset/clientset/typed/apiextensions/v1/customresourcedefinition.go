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

package v1

import (
	"context"
	"time"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	scheme "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CustomResourceDefinitionsGetter has a method to return a CustomResourceDefinitionInterface.
// A group's client should implement this interface.
type CustomResourceDefinitionsGetter interface {
	CustomResourceDefinitions() CustomResourceDefinitionInterface
}

// CustomResourceDefinitionInterface has methods to work with CustomResourceDefinition resources.
type CustomResourceDefinitionInterface interface {
	Create(*v1.CustomResourceDefinition) (*v1.CustomResourceDefinition, error)
	Update(*v1.CustomResourceDefinition) (*v1.CustomResourceDefinition, error)
	UpdateStatus(*v1.CustomResourceDefinition) (*v1.CustomResourceDefinition, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.CustomResourceDefinition, error)
	List(opts metav1.ListOptions) (*v1.CustomResourceDefinitionList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CustomResourceDefinition, err error)
	CustomResourceDefinitionExpansion
}

// customResourceDefinitions implements CustomResourceDefinitionInterface
type customResourceDefinitions struct {
	client rest.Interface
}

// newCustomResourceDefinitions returns a CustomResourceDefinitions
func newCustomResourceDefinitions(c *ApiextensionsV1Client) *customResourceDefinitions {
	return &customResourceDefinitions{
		client: c.RESTClient(),
	}
}

// Get takes name of the customResourceDefinition, and returns the corresponding customResourceDefinition object, and an error if there is any.
func (c *customResourceDefinitions) Get(name string, options metav1.GetOptions) (result *v1.CustomResourceDefinition, err error) {
	result = &v1.CustomResourceDefinition{}
	err = c.client.Get().
		Resource("customresourcedefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CustomResourceDefinitions that match those selectors.
func (c *customResourceDefinitions) List(opts metav1.ListOptions) (result *v1.CustomResourceDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CustomResourceDefinitionList{}
	err = c.client.Get().
		Resource("customresourcedefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested customResourceDefinitions.
func (c *customResourceDefinitions) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("customresourcedefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a customResourceDefinition and creates it.  Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (c *customResourceDefinitions) Create(customResourceDefinition *v1.CustomResourceDefinition) (result *v1.CustomResourceDefinition, err error) {
	result = &v1.CustomResourceDefinition{}
	err = c.client.Post().
		Resource("customresourcedefinitions").
		Body(customResourceDefinition).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a customResourceDefinition and updates it. Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (c *customResourceDefinitions) Update(customResourceDefinition *v1.CustomResourceDefinition) (result *v1.CustomResourceDefinition, err error) {
	result = &v1.CustomResourceDefinition{}
	err = c.client.Put().
		Resource("customresourcedefinitions").
		Name(customResourceDefinition.Name).
		Body(customResourceDefinition).
		Do(context.TODO()).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *customResourceDefinitions) UpdateStatus(customResourceDefinition *v1.CustomResourceDefinition) (result *v1.CustomResourceDefinition, err error) {
	result = &v1.CustomResourceDefinition{}
	err = c.client.Put().
		Resource("customresourcedefinitions").
		Name(customResourceDefinition.Name).
		SubResource("status").
		Body(customResourceDefinition).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the customResourceDefinition and deletes it. Returns an error if one occurs.
func (c *customResourceDefinitions) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("customresourcedefinitions").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *customResourceDefinitions) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("customresourcedefinitions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched customResourceDefinition.
func (c *customResourceDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CustomResourceDefinition, err error) {
	result = &v1.CustomResourceDefinition{}
	err = c.client.Patch(pt).
		Resource("customresourcedefinitions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}
