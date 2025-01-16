// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScyllaDBDatacenters implements ScyllaDBDatacenterInterface
type FakeScyllaDBDatacenters struct {
	Fake *FakeScyllaV1alpha1
	ns   string
}

var scylladbdatacentersResource = v1alpha1.SchemeGroupVersion.WithResource("scylladbdatacenters")

var scylladbdatacentersKind = v1alpha1.SchemeGroupVersion.WithKind("ScyllaDBDatacenter")

// Get takes name of the scyllaDBDatacenter, and returns the corresponding scyllaDBDatacenter object, and an error if there is any.
func (c *FakeScyllaDBDatacenters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	emptyResult := &v1alpha1.ScyllaDBDatacenter{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(scylladbdatacentersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScyllaDBDatacenter), err
}

// List takes label and field selectors, and returns the list of ScyllaDBDatacenters that match those selectors.
func (c *FakeScyllaDBDatacenters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScyllaDBDatacenterList, err error) {
	emptyResult := &v1alpha1.ScyllaDBDatacenterList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(scylladbdatacentersResource, scylladbdatacentersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScyllaDBDatacenterList{ListMeta: obj.(*v1alpha1.ScyllaDBDatacenterList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScyllaDBDatacenterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scyllaDBDatacenters.
func (c *FakeScyllaDBDatacenters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(scylladbdatacentersResource, c.ns, opts))

}

// Create takes the representation of a scyllaDBDatacenter and creates it.  Returns the server's representation of the scyllaDBDatacenter, and an error, if there is any.
func (c *FakeScyllaDBDatacenters) Create(ctx context.Context, scyllaDBDatacenter *v1alpha1.ScyllaDBDatacenter, opts v1.CreateOptions) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	emptyResult := &v1alpha1.ScyllaDBDatacenter{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(scylladbdatacentersResource, c.ns, scyllaDBDatacenter, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScyllaDBDatacenter), err
}

// Update takes the representation of a scyllaDBDatacenter and updates it. Returns the server's representation of the scyllaDBDatacenter, and an error, if there is any.
func (c *FakeScyllaDBDatacenters) Update(ctx context.Context, scyllaDBDatacenter *v1alpha1.ScyllaDBDatacenter, opts v1.UpdateOptions) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	emptyResult := &v1alpha1.ScyllaDBDatacenter{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(scylladbdatacentersResource, c.ns, scyllaDBDatacenter, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScyllaDBDatacenter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScyllaDBDatacenters) UpdateStatus(ctx context.Context, scyllaDBDatacenter *v1alpha1.ScyllaDBDatacenter, opts v1.UpdateOptions) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	emptyResult := &v1alpha1.ScyllaDBDatacenter{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(scylladbdatacentersResource, "status", c.ns, scyllaDBDatacenter, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScyllaDBDatacenter), err
}

// Delete takes name of the scyllaDBDatacenter and deletes it. Returns an error if one occurs.
func (c *FakeScyllaDBDatacenters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(scylladbdatacentersResource, c.ns, name, opts), &v1alpha1.ScyllaDBDatacenter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScyllaDBDatacenters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(scylladbdatacentersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScyllaDBDatacenterList{})
	return err
}

// Patch applies the patch and returns the patched scyllaDBDatacenter.
func (c *FakeScyllaDBDatacenters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	emptyResult := &v1alpha1.ScyllaDBDatacenter{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(scylladbdatacentersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScyllaDBDatacenter), err
}
