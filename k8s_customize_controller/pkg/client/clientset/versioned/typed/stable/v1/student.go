/*
Copyright 2021 The cmp authors .

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
	v1 "k8s_customize_controller/pkg/apis/stable/v1"
	scheme "k8s_customize_controller/pkg/client/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StudentsGetter has a method to return a StudentInterface.
// A group's client should implement this interface.
type StudentsGetter interface {
	Students(namespace string) StudentInterface
}

// StudentInterface has methods to work with Student resources.
type StudentInterface interface {
	Create(ctx context.Context, student *v1.Student, opts metav1.CreateOptions) (*v1.Student, error)
	Update(ctx context.Context, student *v1.Student, opts metav1.UpdateOptions) (*v1.Student, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Student, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.StudentList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Student, err error)
	StudentExpansion
}

// students implements StudentInterface
type students struct {
	client rest.Interface
	ns     string
}

// newStudents returns a Students
func newStudents(c *StableV1Client, namespace string) *students {
	return &students{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the student, and returns the corresponding student object, and an error if there is any.
func (c *students) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Student, err error) {
	result = &v1.Student{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("students").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Students that match those selectors.
func (c *students) List(ctx context.Context, opts metav1.ListOptions) (result *v1.StudentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.StudentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("students").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested students.
func (c *students) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("students").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a student and creates it.  Returns the server's representation of the student, and an error, if there is any.
func (c *students) Create(ctx context.Context, student *v1.Student, opts metav1.CreateOptions) (result *v1.Student, err error) {
	result = &v1.Student{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("students").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(student).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a student and updates it. Returns the server's representation of the student, and an error, if there is any.
func (c *students) Update(ctx context.Context, student *v1.Student, opts metav1.UpdateOptions) (result *v1.Student, err error) {
	result = &v1.Student{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("students").
		Name(student.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(student).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the student and deletes it. Returns an error if one occurs.
func (c *students) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("students").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *students) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("students").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched student.
func (c *students) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Student, err error) {
	result = &v1.Student{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("students").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
