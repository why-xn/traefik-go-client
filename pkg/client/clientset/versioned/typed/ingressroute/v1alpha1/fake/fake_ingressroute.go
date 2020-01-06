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
	v1alpha1 "github.com/why-xn/traefik-go-client/pkg/apis/ingressroute/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIngressRoutes implements IngressRouteInterface
type FakeIngressRoutes struct {
	Fake *FakeIngressrouteV1alpha1
	ns   string
}

var ingressroutesResource = schema.GroupVersionResource{Group: "ingressroute", Version: "v1alpha1", Resource: "ingressroutes"}

var ingressroutesKind = schema.GroupVersionKind{Group: "ingressroute", Version: "v1alpha1", Kind: "IngressRoute"}

// Get takes name of the ingressRoute, and returns the corresponding ingressRoute object, and an error if there is any.
func (c *FakeIngressRoutes) Get(name string, options v1.GetOptions) (result *v1alpha1.IngressRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ingressroutesResource, c.ns, name), &v1alpha1.IngressRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressRoute), err
}

// List takes label and field selectors, and returns the list of IngressRoutes that match those selectors.
func (c *FakeIngressRoutes) List(opts v1.ListOptions) (result *v1alpha1.IngressRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ingressroutesResource, ingressroutesKind, c.ns, opts), &v1alpha1.IngressRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IngressRouteList{ListMeta: obj.(*v1alpha1.IngressRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.IngressRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingressRoutes.
func (c *FakeIngressRoutes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ingressroutesResource, c.ns, opts))

}

// Create takes the representation of a ingressRoute and creates it.  Returns the server's representation of the ingressRoute, and an error, if there is any.
func (c *FakeIngressRoutes) Create(ingressRoute *v1alpha1.IngressRoute) (result *v1alpha1.IngressRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ingressroutesResource, c.ns, ingressRoute), &v1alpha1.IngressRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressRoute), err
}

// Update takes the representation of a ingressRoute and updates it. Returns the server's representation of the ingressRoute, and an error, if there is any.
func (c *FakeIngressRoutes) Update(ingressRoute *v1alpha1.IngressRoute) (result *v1alpha1.IngressRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ingressroutesResource, c.ns, ingressRoute), &v1alpha1.IngressRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressRoute), err
}

// Delete takes name of the ingressRoute and deletes it. Returns an error if one occurs.
func (c *FakeIngressRoutes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ingressroutesResource, c.ns, name), &v1alpha1.IngressRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressRoutes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ingressroutesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IngressRouteList{})
	return err
}

// Patch applies the patch and returns the patched ingressRoute.
func (c *FakeIngressRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IngressRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingressroutesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IngressRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressRoute), err
}
