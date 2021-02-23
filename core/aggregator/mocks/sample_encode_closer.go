// Code generated by mockery v1.0.0
package aggregatemock

import (
	mock "github.com/stretchr/testify/mock"

	core "github.com/phsm/pandora/core"
)

// SampleEncodeCloser is an autogenerated mock type for the SampleEncodeCloser type
type SampleEncodeCloser struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *SampleEncodeCloser) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Encode provides a mock function with given fields: s
func (_m *SampleEncodeCloser) Encode(s core.Sample) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Sample) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Flush provides a mock function with given fields:
func (_m *SampleEncodeCloser) Flush() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
