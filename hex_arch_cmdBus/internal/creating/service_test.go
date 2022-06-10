package creating

import (
	"context"
	"errors"
	"testing"

	mooc "github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/platform/storage/storagemocks"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CourseService_CreateCourse_RepositoryError(t *testing.T) {
	courseID := "cbb20d26-e1b0-11ec-8fea-0242ac120002"
	courseName := "Test Course"
	courseDuration := "10 hours"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(errors.New("something unexpected happened"))

	eventBusMock := new(eventmocks.Bus)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)
	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)
	courseRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Succeed(t *testing.T) {
	courseID := "cbb20d26-e1b0-11ec-8fea-0242ac120002"
	courseName := "Test Course"
	courseDuration := "10 hours"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(nil)
	eventBusMock := new(eventmocks.Bus)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
