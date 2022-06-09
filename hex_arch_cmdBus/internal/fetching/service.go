package fetching

import (
	"context"
	"fmt"

	mooc "github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal"
)

type CourseService struct {
	courseRepository mooc.CourseRepository
}

func NewCourseService(courseRepository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository: courseRepository,
	}
}

func (s CourseService) GetAllCourses(ctx context.Context) ([]mooc.Course, error) {
	fmt.Println("Get all courses application layer")
	return s.courseRepository.GetCourses(ctx)
}
