// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "challenge/alerts/src/api/alerts/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AlertsRepository is an autogenerated mock type for the AlertsRepository type
type AlertsRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, alert
func (_m *AlertsRepository) Create(ctx context.Context, alert domain.Alert) (*domain.Alert, error) {
	ret := _m.Called(ctx, alert)

	var r0 *domain.Alert
	if rf, ok := ret.Get(0).(func(context.Context, domain.Alert) *domain.Alert); ok {
		r0 = rf(ctx, alert)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Alert) error); ok {
		r1 = rf(ctx, alert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAlertsByType provides a mock function with given fields: ctx, typeInput
func (_m *AlertsRepository) GetAlertsByType(ctx context.Context, typeInput string) ([]domain.Alert, error) {
	ret := _m.Called(ctx, typeInput)

	var r0 []domain.Alert
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Alert); ok {
		r0 = rf(ctx, typeInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, typeInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *AlertsRepository) GetAll(ctx context.Context) ([]domain.Alert, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Alert
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Alert); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Alert)
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

// GetMetrics provides a mock function with given fields: ctx
func (_m *AlertsRepository) GetMetrics(ctx context.Context) ([]domain.Metrics, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Metrics
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Metrics); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Metrics)
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

// Search provides a mock function with given fields: ctx, input
func (_m *AlertsRepository) Search(ctx context.Context, input domain.AlertSearch) ([]domain.Alert, error) {
	ret := _m.Called(ctx, input)

	var r0 []domain.Alert
	if rf, ok := ret.Get(0).(func(context.Context, domain.AlertSearch) []domain.Alert); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.AlertSearch) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAlertsRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAlertsRepository creates a new instance of AlertsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAlertsRepository(t mockConstructorTestingTNewAlertsRepository) *AlertsRepository {
	mock := &AlertsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
