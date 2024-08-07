/*
Copyright 2024 llmos.ai.

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
// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/llmos-ai/llmos-operator/pkg/apis/ml.llmos.ai/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeModelFiles implements ModelFileInterface
type FakeModelFiles struct {
	Fake *FakeMlV1
}

var modelfilesResource = v1.SchemeGroupVersion.WithResource("modelfiles")

var modelfilesKind = v1.SchemeGroupVersion.WithKind("ModelFile")

// Get takes name of the modelFile, and returns the corresponding modelFile object, and an error if there is any.
func (c *FakeModelFiles) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ModelFile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(modelfilesResource, name), &v1.ModelFile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ModelFile), err
}

// List takes label and field selectors, and returns the list of ModelFiles that match those selectors.
func (c *FakeModelFiles) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ModelFileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(modelfilesResource, modelfilesKind, opts), &v1.ModelFileList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ModelFileList{ListMeta: obj.(*v1.ModelFileList).ListMeta}
	for _, item := range obj.(*v1.ModelFileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested modelFiles.
func (c *FakeModelFiles) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(modelfilesResource, opts))
}

// Create takes the representation of a modelFile and creates it.  Returns the server's representation of the modelFile, and an error, if there is any.
func (c *FakeModelFiles) Create(ctx context.Context, modelFile *v1.ModelFile, opts metav1.CreateOptions) (result *v1.ModelFile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(modelfilesResource, modelFile), &v1.ModelFile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ModelFile), err
}

// Update takes the representation of a modelFile and updates it. Returns the server's representation of the modelFile, and an error, if there is any.
func (c *FakeModelFiles) Update(ctx context.Context, modelFile *v1.ModelFile, opts metav1.UpdateOptions) (result *v1.ModelFile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(modelfilesResource, modelFile), &v1.ModelFile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ModelFile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeModelFiles) UpdateStatus(ctx context.Context, modelFile *v1.ModelFile, opts metav1.UpdateOptions) (*v1.ModelFile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(modelfilesResource, "status", modelFile), &v1.ModelFile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ModelFile), err
}

// Delete takes name of the modelFile and deletes it. Returns an error if one occurs.
func (c *FakeModelFiles) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(modelfilesResource, name, opts), &v1.ModelFile{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeModelFiles) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(modelfilesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ModelFileList{})
	return err
}

// Patch applies the patch and returns the patched modelFile.
func (c *FakeModelFiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ModelFile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(modelfilesResource, name, pt, data, subresources...), &v1.ModelFile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ModelFile), err
}
