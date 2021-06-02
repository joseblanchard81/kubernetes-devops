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

// xCode generated by client-gen. DO NOT EDIT.

package v1alpha3

import (
	rest "k8s.io/client-go/rest"
	v1alpha3 "kubesphere.io/devops/api/v1alpha3"
	"kubesphere.io/devops/pkg/client/clientset/versioned/scheme"
)

type DevopsV1alpha3Interface interface {
	RESTClient() rest.Interface
	DevOpsProjectsGetter
	PipelinesGetter
}

// DevopsV1alpha3Client is used to interact with features provided by the devops.kubesphere.io group.
type DevopsV1alpha3Client struct {
	restClient rest.Interface
}

func (c *DevopsV1alpha3Client) DevOpsProjects() DevOpsProjectInterface {
	return newDevOpsProjects(c)
}

func (c *DevopsV1alpha3Client) Pipelines(namespace string) PipelineInterface {
	return newPipelines(c, namespace)
}

// NewForConfig creates a new DevopsV1alpha3Client for the given config.
func NewForConfig(c *rest.Config) (*DevopsV1alpha3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DevopsV1alpha3Client{client}, nil
}

// NewForConfigOrDie creates a new DevopsV1alpha3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *DevopsV1alpha3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new DevopsV1alpha3Client for the given RESTClient.
func New(c rest.Interface) *DevopsV1alpha3Client {
	return &DevopsV1alpha3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha3.GroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *DevopsV1alpha3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
