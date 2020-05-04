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
	v1 "github.com/slok/bilrost/pkg/kubernetes/gen/clientset/versioned/typed/auth/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAuthV1 struct {
	*testing.Fake
}

func (c *FakeAuthV1) AuthBackends() v1.AuthBackendInterface {
	return &FakeAuthBackends{c}
}

func (c *FakeAuthV1) IngressAuths(namespace string) v1.IngressAuthInterface {
	return &FakeIngressAuths{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAuthV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
