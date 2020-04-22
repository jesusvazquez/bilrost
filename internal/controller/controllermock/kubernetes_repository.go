// Code generated by mockery v1.0.0. DO NOT EDIT.

package controllermock

import context "context"

import mock "github.com/stretchr/testify/mock"
import v1beta1 "k8s.io/api/networking/v1beta1"
import watch "k8s.io/apimachinery/pkg/watch"

// KubernetesRepository is an autogenerated mock type for the KubernetesRepository type
type KubernetesRepository struct {
	mock.Mock
}

// GetIngress provides a mock function with given fields: ctx, ns, name
func (_m *KubernetesRepository) GetIngress(ctx context.Context, ns string, name string) (*v1beta1.Ingress, error) {
	ret := _m.Called(ctx, ns, name)

	var r0 *v1beta1.Ingress
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *v1beta1.Ingress); ok {
		r0 = rf(ctx, ns, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Ingress)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ns, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIngresses provides a mock function with given fields: ctx, ns, labelSelector
func (_m *KubernetesRepository) ListIngresses(ctx context.Context, ns string, labelSelector map[string]string) (*v1beta1.IngressList, error) {
	ret := _m.Called(ctx, ns, labelSelector)

	var r0 *v1beta1.IngressList
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) *v1beta1.IngressList); ok {
		r0 = rf(ctx, ns, labelSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.IngressList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, ns, labelSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIngress provides a mock function with given fields: ctx, ingress
func (_m *KubernetesRepository) UpdateIngress(ctx context.Context, ingress *v1beta1.Ingress) error {
	ret := _m.Called(ctx, ingress)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Ingress) error); ok {
		r0 = rf(ctx, ingress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WatchIngresses provides a mock function with given fields: ctx, ns, labelSelector
func (_m *KubernetesRepository) WatchIngresses(ctx context.Context, ns string, labelSelector map[string]string) (watch.Interface, error) {
	ret := _m.Called(ctx, ns, labelSelector)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) watch.Interface); ok {
		r0 = rf(ctx, ns, labelSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, ns, labelSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}