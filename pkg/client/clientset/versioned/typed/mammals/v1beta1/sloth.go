// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "kubebuilder-demo/pkg/apis/mammals/v1beta1"
	scheme "kubebuilder-demo/pkg/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SlothsGetter has a method to return a SlothInterface.
// A group's client should implement this interface.
type SlothsGetter interface {
	Sloths(namespace string) SlothInterface
}

// SlothInterface has methods to work with Sloth resources.
type SlothInterface interface {
	Create(*v1beta1.Sloth) (*v1beta1.Sloth, error)
	Update(*v1beta1.Sloth) (*v1beta1.Sloth, error)
	UpdateStatus(*v1beta1.Sloth) (*v1beta1.Sloth, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Sloth, error)
	List(opts v1.ListOptions) (*v1beta1.SlothList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Sloth, err error)
	SlothExpansion
}

// sloths implements SlothInterface
type sloths struct {
	client rest.Interface
	ns     string
}

// newSloths returns a Sloths
func newSloths(c *MammalsV1beta1Client, namespace string) *sloths {
	return &sloths{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sloth, and returns the corresponding sloth object, and an error if there is any.
func (c *sloths) Get(name string, options v1.GetOptions) (result *v1beta1.Sloth, err error) {
	result = &v1beta1.Sloth{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sloths").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Sloths that match those selectors.
func (c *sloths) List(opts v1.ListOptions) (result *v1beta1.SlothList, err error) {
	result = &v1beta1.SlothList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sloths").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sloths.
func (c *sloths) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sloths").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sloth and creates it.  Returns the server's representation of the sloth, and an error, if there is any.
func (c *sloths) Create(sloth *v1beta1.Sloth) (result *v1beta1.Sloth, err error) {
	result = &v1beta1.Sloth{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sloths").
		Body(sloth).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sloth and updates it. Returns the server's representation of the sloth, and an error, if there is any.
func (c *sloths) Update(sloth *v1beta1.Sloth) (result *v1beta1.Sloth, err error) {
	result = &v1beta1.Sloth{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sloths").
		Name(sloth.Name).
		Body(sloth).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sloths) UpdateStatus(sloth *v1beta1.Sloth) (result *v1beta1.Sloth, err error) {
	result = &v1beta1.Sloth{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sloths").
		Name(sloth.Name).
		SubResource("status").
		Body(sloth).
		Do().
		Into(result)
	return
}

// Delete takes name of the sloth and deletes it. Returns an error if one occurs.
func (c *sloths) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sloths").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sloths) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sloths").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sloth.
func (c *sloths) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Sloth, err error) {
	result = &v1beta1.Sloth{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sloths").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
