package fetching

import (
	"context"
	"errors"
	"fmt"
	"testing"

	mooc "github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CourseService_FetchCourses_RepositoryError(t *testing.T) {
	courseID := "cbb20d26-e1b0-11ec-8fea-0242ac120002"
	courseName := "Test Course"
	courseDuration := "10 hours"

	var courses []mooc.Course
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)
	courses = append(courses, course)

	courseRespositoryMock := new(storagemocks.CourseRepository)
	courseRespositoryMock.On("GetCourses", mock.Anything /* corresponds to context.Background()*/).Return(courses, errors.New("something unexpected happened"))
	//courseRespositoryMock.On("GetCourses", context.Background()).Return(courses, errors.New("something unexpected happened"))

	courseService := NewCourseService(courseRespositoryMock)
	_, err = courseService.GetAllCourses(context.Background())

	courseRespositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_FetchCourses_RepositorySuccess(t *testing.T) {
	courseID := "cbb20d26-e1b0-11ec-8fea-0242ac120002"
	courseName := "Test Course"
	courseDuration := "10 hours"

	var courses []mooc.Course
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)
	courses = append(courses, course)

	courseRespositoryMock := new(storagemocks.CourseRepository)
	courseRespositoryMock.On("GetCourses", mock.Anything).Return(courses, nil)

	courseService := NewCourseService(courseRespositoryMock)

	result, err := courseService.GetAllCourses(context.Background())
	fmt.Println("------------------ result from courseServices getall ->", result)

	courseRespositoryMock.AssertExpectations(t)
	assert.NoError(t, err)

}
