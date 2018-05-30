/*
Copyright 2018 The Attic Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/appscode/messenger/apis/messenger/v1alpha1"
	scheme "github.com/appscode/messenger/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NotifiersGetter has a method to return a NotifierInterface.
// A group's client should implement this interface.
type NotifiersGetter interface {
	Notifiers(namespace string) NotifierInterface
}

// NotifierInterface has methods to work with Notifier resources.
type NotifierInterface interface {
	Create(*v1alpha1.Notifier) (*v1alpha1.Notifier, error)
	Update(*v1alpha1.Notifier) (*v1alpha1.Notifier, error)
	UpdateStatus(*v1alpha1.Notifier) (*v1alpha1.Notifier, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Notifier, error)
	List(opts v1.ListOptions) (*v1alpha1.NotifierList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Notifier, err error)
	NotifierExpansion
}

// notifiers implements NotifierInterface
type notifiers struct {
	client rest.Interface
	ns     string
}

// newNotifiers returns a Notifiers
func newNotifiers(c *MessengerV1alpha1Client, namespace string) *notifiers {
	return &notifiers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the notifier, and returns the corresponding notifier object, and an error if there is any.
func (c *notifiers) Get(name string, options v1.GetOptions) (result *v1alpha1.Notifier, err error) {
	result = &v1alpha1.Notifier{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notifiers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Notifiers that match those selectors.
func (c *notifiers) List(opts v1.ListOptions) (result *v1alpha1.NotifierList, err error) {
	result = &v1alpha1.NotifierList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notifiers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested notifiers.
func (c *notifiers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("notifiers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a notifier and creates it.  Returns the server's representation of the notifier, and an error, if there is any.
func (c *notifiers) Create(notifier *v1alpha1.Notifier) (result *v1alpha1.Notifier, err error) {
	result = &v1alpha1.Notifier{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("notifiers").
		Body(notifier).
		Do().
		Into(result)
	return
}

// Update takes the representation of a notifier and updates it. Returns the server's representation of the notifier, and an error, if there is any.
func (c *notifiers) Update(notifier *v1alpha1.Notifier) (result *v1alpha1.Notifier, err error) {
	result = &v1alpha1.Notifier{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notifiers").
		Name(notifier.Name).
		Body(notifier).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *notifiers) UpdateStatus(notifier *v1alpha1.Notifier) (result *v1alpha1.Notifier, err error) {
	result = &v1alpha1.Notifier{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notifiers").
		Name(notifier.Name).
		SubResource("status").
		Body(notifier).
		Do().
		Into(result)
	return
}

// Delete takes name of the notifier and deletes it. Returns an error if one occurs.
func (c *notifiers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notifiers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *notifiers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notifiers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched notifier.
func (c *notifiers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Notifier, err error) {
	result = &v1alpha1.Notifier{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("notifiers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}