package creating

import (
	"context"
	"errors"
	"testing"

	mooc "github.com/krls08/hex-arch-api-go/hex_arch/internal"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/storage/storagemocks"
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

	courseRespositoryMock := new(storagemocks.CourseRepository)
	courseRespositoryMock.On("Save", mock.Anything, course).Return(errors.New("something unexpected happened"))

	courseService := NewCourseSerivce(courseRespositoryMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)
	courseRespositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Succeed(t *testing.T) {
	courseID := "cbb20d26-e1b0-11ec-8fea-0242ac120002"
	courseName := "Test Course"
	courseDuration := "10 hours"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRespositoryMock := new(storagemocks.CourseRepository)
	courseRespositoryMock.On("Save", mock.Anything, course).Return(nil)

	courseService := NewCourseSerivce(courseRespositoryMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRespositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
