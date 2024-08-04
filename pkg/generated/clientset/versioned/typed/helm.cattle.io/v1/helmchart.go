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

package v1

import (
	"context"
	"time"

	v1 "github.com/k3s-io/helm-controller/pkg/apis/helm.cattle.io/v1"
	scheme "github.com/llmos-ai/llmos-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HelmChartsGetter has a method to return a HelmChartInterface.
// A group's client should implement this interface.
type HelmChartsGetter interface {
	HelmCharts(namespace string) HelmChartInterface
}

// HelmChartInterface has methods to work with HelmChart resources.
type HelmChartInterface interface {
	Create(ctx context.Context, helmChart *v1.HelmChart, opts metav1.CreateOptions) (*v1.HelmChart, error)
	Update(ctx context.Context, helmChart *v1.HelmChart, opts metav1.UpdateOptions) (*v1.HelmChart, error)
	UpdateStatus(ctx context.Context, helmChart *v1.HelmChart, opts metav1.UpdateOptions) (*v1.HelmChart, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.HelmChart, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.HelmChartList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HelmChart, err error)
	HelmChartExpansion
}

// helmCharts implements HelmChartInterface
type helmCharts struct {
	client rest.Interface
	ns     string
}

// newHelmCharts returns a HelmCharts
func newHelmCharts(c *HelmV1Client, namespace string) *helmCharts {
	return &helmCharts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the helmChart, and returns the corresponding helmChart object, and an error if there is any.
func (c *helmCharts) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.HelmChart, err error) {
	result = &v1.HelmChart{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("helmcharts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HelmCharts that match those selectors.
func (c *helmCharts) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HelmChartList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.HelmChartList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("helmcharts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested helmCharts.
func (c *helmCharts) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("helmcharts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a helmChart and creates it.  Returns the server's representation of the helmChart, and an error, if there is any.
func (c *helmCharts) Create(ctx context.Context, helmChart *v1.HelmChart, opts metav1.CreateOptions) (result *v1.HelmChart, err error) {
	result = &v1.HelmChart{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("helmcharts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmChart).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a helmChart and updates it. Returns the server's representation of the helmChart, and an error, if there is any.
func (c *helmCharts) Update(ctx context.Context, helmChart *v1.HelmChart, opts metav1.UpdateOptions) (result *v1.HelmChart, err error) {
	result = &v1.HelmChart{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("helmcharts").
		Name(helmChart.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmChart).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *helmCharts) UpdateStatus(ctx context.Context, helmChart *v1.HelmChart, opts metav1.UpdateOptions) (result *v1.HelmChart, err error) {
	result = &v1.HelmChart{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("helmcharts").
		Name(helmChart.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmChart).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the helmChart and deletes it. Returns an error if one occurs.
func (c *helmCharts) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("helmcharts").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *helmCharts) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("helmcharts").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched helmChart.
func (c *helmCharts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HelmChart, err error) {
	result = &v1.HelmChart{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("helmcharts").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
