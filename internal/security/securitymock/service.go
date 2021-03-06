// Code generated by mockery v1.0.0. DO NOT EDIT.

package securitymock

import (
	context "context"

	model "github.com/slok/bilrost/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// RollbackAppSecurity provides a mock function with given fields: ctx, app
func (_m *Service) RollbackAppSecurity(ctx context.Context, app model.App) error {
	ret := _m.Called(ctx, app)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.App) error); ok {
		r0 = rf(ctx, app)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SecureApp provides a mock function with given fields: ctx, app
func (_m *Service) SecureApp(ctx context.Context, app model.App) error {
	ret := _m.Called(ctx, app)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.App) error); ok {
		r0 = rf(ctx, app)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
