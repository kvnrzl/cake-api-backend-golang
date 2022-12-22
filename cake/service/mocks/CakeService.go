// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	model "backend-engineer-test-privy/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CakeService is an autogenerated mock type for the CakeService type
type CakeService struct {
	mock.Mock
}

// CreateCake provides a mock function with given fields: ctx, cake
func (_m *CakeService) CreateCake(ctx context.Context, cake *model.Cake) (*model.Cake, error) {
	ret := _m.Called(ctx, cake)

	var r0 *model.Cake
	if rf, ok := ret.Get(0).(func(context.Context, *model.Cake) *model.Cake); ok {
		r0 = rf(ctx, cake)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Cake)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Cake) error); ok {
		r1 = rf(ctx, cake)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCake provides a mock function with given fields: ctx, id
func (_m *CakeService) DeleteCake(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCakes provides a mock function with given fields: ctx
func (_m *CakeService) GetAllCakes(ctx context.Context) ([]*model.Cake, error) {
	ret := _m.Called(ctx)

	var r0 []*model.Cake
	if rf, ok := ret.Get(0).(func(context.Context) []*model.Cake); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Cake)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCakeByID provides a mock function with given fields: ctx, id
func (_m *CakeService) GetCakeByID(ctx context.Context, id uint) (*model.Cake, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Cake
	if rf, ok := ret.Get(0).(func(context.Context, uint) *model.Cake); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Cake)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCake provides a mock function with given fields: ctx, cake
func (_m *CakeService) UpdateCake(ctx context.Context, cake *model.Cake) (*model.Cake, error) {
	ret := _m.Called(ctx, cake)

	var r0 *model.Cake
	if rf, ok := ret.Get(0).(func(context.Context, *model.Cake) *model.Cake); ok {
		r0 = rf(ctx, cake)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Cake)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Cake) error); ok {
		r1 = rf(ctx, cake)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCakeService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCakeService creates a new instance of CakeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCakeService(t mockConstructorTestingTNewCakeService) *CakeService {
	mock := &CakeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}