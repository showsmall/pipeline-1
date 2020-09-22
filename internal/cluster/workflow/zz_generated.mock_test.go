// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package workflow

import "github.com/stretchr/testify/mock"

// MockK8sConfigGetter is an autogenerated mock for the K8sConfigGetter type.
type MockK8sConfigGetter struct {
	mock.Mock
}

// Get provides a mock function.
func (_m *MockK8sConfigGetter) Get(organizationID uint, k8sSecretID string) (_result_0 []uint8, _result_1 error) {
	ret := _m.Called(organizationID, k8sSecretID)

	var r0 []uint8
	if rf, ok := ret.Get(0).(func(uint, string) []uint8); ok {
		r0 = rf(organizationID, k8sSecretID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint8)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(organizationID, k8sSecretID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}