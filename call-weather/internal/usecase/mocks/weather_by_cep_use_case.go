// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/fabioods/go-expert-call-weather/internal/domain"
	mock "github.com/stretchr/testify/mock"

	usecase "github.com/fabioods/go-expert-call-weather/internal/usecase"
)

// WeatherByCepUseCase is an autogenerated mock type for the WeatherByCepUseCase type
type WeatherByCepUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, input
func (_m *WeatherByCepUseCase) Execute(_a0 context.Context, input usecase.InputDTO) (domain.Cep, error) {
	ret := _m.Called(_a0, input)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 domain.Cep
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, usecase.InputDTO) (domain.Cep, error)); ok {
		return rf(_a0, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, usecase.InputDTO) domain.Cep); ok {
		r0 = rf(_a0, input)
	} else {
		r0 = ret.Get(0).(domain.Cep)
	}

	if rf, ok := ret.Get(1).(func(context.Context, usecase.InputDTO) error); ok {
		r1 = rf(_a0, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewWeatherByCepUseCase creates a new instance of WeatherByCepUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWeatherByCepUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *WeatherByCepUseCase {
	mock := &WeatherByCepUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
