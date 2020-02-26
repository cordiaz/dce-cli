// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import os "os"

// TerraformBinFileSystemUtil is an autogenerated mock type for the TerraformBinFileSystemUtil type
type TerraformBinFileSystemUtil struct {
	mock.Mock
}

// GetConfigDir provides a mock function with given fields:
func (_m *TerraformBinFileSystemUtil) GetConfigDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetLocalTFModuleDir provides a mock function with given fields:
func (_m *TerraformBinFileSystemUtil) GetLocalTFModuleDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTerraformBin provides a mock function with given fields:
func (_m *TerraformBinFileSystemUtil) GetTerraformBin() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTerraformBinDir provides a mock function with given fields:
func (_m *TerraformBinFileSystemUtil) GetTerraformBinDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsExistingFile provides a mock function with given fields: path
func (_m *TerraformBinFileSystemUtil) IsExistingFile(path string) bool {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OpenFileWriter provides a mock function with given fields: path
func (_m *TerraformBinFileSystemUtil) OpenFileWriter(path string) (*os.File, error) {
	ret := _m.Called(path)

	var r0 *os.File
	if rf, ok := ret.Get(0).(func(string) *os.File); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveAll provides a mock function with given fields: path
func (_m *TerraformBinFileSystemUtil) RemoveAll(path string) {
	_m.Called(path)
}

// Unarchive provides a mock function with given fields: source, destination
func (_m *TerraformBinFileSystemUtil) Unarchive(source string, destination string) error {
	ret := _m.Called(source, destination)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(source, destination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
