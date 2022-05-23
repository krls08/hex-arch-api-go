// Code generated by mockery v2.12.2. DO NOT EDIT.

package storagemocks

import (
	context "context"

	mooc "github.com/krls08/hex-arch-api-go/hex_arch/internal"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// CourseRepository is an autogenerated mock type for the CourseRepository type
type CourseRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, course
func (_m *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	ret := _m.Called(ctx, course)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mooc.Course) error); ok {
		r0 = rf(ctx, course)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCourseRepository creates a new instance of CourseRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCourseRepository(t testing.TB) *CourseRepository {
	mock := &CourseRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
