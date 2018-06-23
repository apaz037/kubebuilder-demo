// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kubernetes-sigs/kubebuilder/test/projects/validations/pkg/apis/apps/v1"
	scheme "github.com/kubernetes-sigs/kubebuilder/test/projects/validations/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TestKindsGetter has a method to return a TestKindInterface.
// A group's client should implement this interface.
type TestKindsGetter interface {
	TestKinds(namespace string) TestKindInterface
}

// TestKindInterface has methods to work with TestKind resources.
type TestKindInterface interface {
	Create(*v1.TestKind) (*v1.TestKind, error)
	Update(*v1.TestKind) (*v1.TestKind, error)
	UpdateStatus(*v1.TestKind) (*v1.TestKind, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.TestKind, error)
	List(opts meta_v1.ListOptions) (*v1.TestKindList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TestKind, err error)
	TestKindExpansion
}

// testKinds implements TestKindInterface
type testKinds struct {
	client rest.Interface
	ns     string
}

// newTestKinds returns a TestKinds
func newTestKinds(c *AppsV1Client, namespace string) *testKinds {
	return &testKinds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the testKind, and returns the corresponding testKind object, and an error if there is any.
func (c *testKinds) Get(name string, options meta_v1.GetOptions) (result *v1.TestKind, err error) {
	result = &v1.TestKind{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testkinds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TestKinds that match those selectors.
func (c *testKinds) List(opts meta_v1.ListOptions) (result *v1.TestKindList, err error) {
	result = &v1.TestKindList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testkinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested testKinds.
func (c *testKinds) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("testkinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a testKind and creates it.  Returns the server's representation of the testKind, and an error, if there is any.
func (c *testKinds) Create(testKind *v1.TestKind) (result *v1.TestKind, err error) {
	result = &v1.TestKind{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("testkinds").
		Body(testKind).
		Do().
		Into(result)
	return
}

// Update takes the representation of a testKind and updates it. Returns the server's representation of the testKind, and an error, if there is any.
func (c *testKinds) Update(testKind *v1.TestKind) (result *v1.TestKind, err error) {
	result = &v1.TestKind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testkinds").
		Name(testKind.Name).
		Body(testKind).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *testKinds) UpdateStatus(testKind *v1.TestKind) (result *v1.TestKind, err error) {
	result = &v1.TestKind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testkinds").
		Name(testKind.Name).
		SubResource("status").
		Body(testKind).
		Do().
		Into(result)
	return
}

// Delete takes name of the testKind and deletes it. Returns an error if one occurs.
func (c *testKinds) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testkinds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *testKinds) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testkinds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched testKind.
func (c *testKinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TestKind, err error) {
	result = &v1.TestKind{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("testkinds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
