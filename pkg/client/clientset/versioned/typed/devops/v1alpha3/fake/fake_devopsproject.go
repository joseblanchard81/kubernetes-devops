/*
Copyright 2020 The KubeSphere Authors.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha3 "kubesphere.io/devops/pkg/api/devops/v1alpha3"
)

// FakeDevOpsProjects implements DevOpsProjectInterface
type FakeDevOpsProjects struct {
	Fake *FakeDevopsV1alpha3
}

var devopsprojectsResource = schema.GroupVersionResource{Group: "devops.kubesphere.io", Version: "v1alpha3", Resource: "devopsprojects"}

var devopsprojectsKind = schema.GroupVersionKind{Group: "devops.kubesphere.io", Version: "v1alpha3", Kind: "DevOpsProject"}

// Get takes name of the devOpsProject, and returns the corresponding devOpsProject object, and an error if there is any.
func (c *FakeDevOpsProjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.DevOpsProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(devopsprojectsResource, name), &v1alpha3.DevOpsProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.DevOpsProject), err
}

// List takes label and field selectors, and returns the list of DevOpsProjects that match those selectors.
func (c *FakeDevOpsProjects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.DevOpsProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(devopsprojectsResource, devopsprojectsKind, opts), &v1alpha3.DevOpsProjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.DevOpsProjectList{ListMeta: obj.(*v1alpha3.DevOpsProjectList).ListMeta}
	for _, item := range obj.(*v1alpha3.DevOpsProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devOpsProjects.
func (c *FakeDevOpsProjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(devopsprojectsResource, opts))
}

// Create takes the representation of a devOpsProject and creates it.  Returns the server's representation of the devOpsProject, and an error, if there is any.
func (c *FakeDevOpsProjects) Create(ctx context.Context, devOpsProject *v1alpha3.DevOpsProject, opts v1.CreateOptions) (result *v1alpha3.DevOpsProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(devopsprojectsResource, devOpsProject), &v1alpha3.DevOpsProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.DevOpsProject), err
}

// Update takes the representation of a devOpsProject and updates it. Returns the server's representation of the devOpsProject, and an error, if there is any.
func (c *FakeDevOpsProjects) Update(ctx context.Context, devOpsProject *v1alpha3.DevOpsProject, opts v1.UpdateOptions) (result *v1alpha3.DevOpsProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(devopsprojectsResource, devOpsProject), &v1alpha3.DevOpsProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.DevOpsProject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevOpsProjects) UpdateStatus(ctx context.Context, devOpsProject *v1alpha3.DevOpsProject, opts v1.UpdateOptions) (*v1alpha3.DevOpsProject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(devopsprojectsResource, "status", devOpsProject), &v1alpha3.DevOpsProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.DevOpsProject), err
}

// Delete takes name of the devOpsProject and deletes it. Returns an error if one occurs.
func (c *FakeDevOpsProjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(devopsprojectsResource, name), &v1alpha3.DevOpsProject{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevOpsProjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(devopsprojectsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha3.DevOpsProjectList{})
	return err
}

// Patch applies the patch and returns the patched devOpsProject.
func (c *FakeDevOpsProjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.DevOpsProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(devopsprojectsResource, name, pt, data, subresources...), &v1alpha3.DevOpsProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.DevOpsProject), err
}
