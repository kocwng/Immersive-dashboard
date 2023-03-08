// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	classes "immersiveApp/features/classes"

	mock "github.com/stretchr/testify/mock"
)

// ClassServiceInterface is an autogenerated mock type for the ClassServiceInterface type
type ClassServiceInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: classEntity
func (_m *ClassServiceInterface) Create(classEntity classes.ClassEntity) (classes.ClassEntity, error) {
	ret := _m.Called(classEntity)

	var r0 classes.ClassEntity
	if rf, ok := ret.Get(0).(func(classes.ClassEntity) classes.ClassEntity); ok {
		r0 = rf(classEntity)
	} else {
		r0 = ret.Get(0).(classes.ClassEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(classes.ClassEntity) error); ok {
		r1 = rf(classEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ClassServiceInterface) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ClassServiceInterface) GetAll() ([]classes.ClassEntity, error) {
	ret := _m.Called()

	var r0 []classes.ClassEntity
	if rf, ok := ret.Get(0).(func() []classes.ClassEntity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]classes.ClassEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *ClassServiceInterface) GetById(id uint) (classes.ClassEntity, error) {
	ret := _m.Called(id)

	var r0 classes.ClassEntity
	if rf, ok := ret.Get(0).(func(uint) classes.ClassEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(classes.ClassEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: classEntity, id
func (_m *ClassServiceInterface) Update(classEntity classes.ClassEntity, id uint) (classes.ClassEntity, error) {
	ret := _m.Called(classEntity, id)

	var r0 classes.ClassEntity
	if rf, ok := ret.Get(0).(func(classes.ClassEntity, uint) classes.ClassEntity); ok {
		r0 = rf(classEntity, id)
	} else {
		r0 = ret.Get(0).(classes.ClassEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(classes.ClassEntity, uint) error); ok {
		r1 = rf(classEntity, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClassServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewClassServiceInterface creates a new instance of ClassServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClassServiceInterface(t mockConstructorTestingTNewClassServiceInterface) *ClassServiceInterface {
	mock := &ClassServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
