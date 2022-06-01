package creating

import (
	"context"
	"fmt"

	mooc "github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService
type CourseService struct {
	courseRepository mooc.CourseRepository
}

// NewCourseService returns de default Service interface implementation
func NewCourseSerivce(courseRepository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository: courseRepository,
	}
}

// CreateCourse implements the creating.CourseService interface
func (s CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	fmt.Println("create course application layer")
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}
	return s.courseRepository.Save(ctx, course)
}

func (s CourseService) GetAllCourses(ctx context.Context) ([]mooc.Course, error) {
	fmt.Println("Get all courses application layer")
	return s.courseRepository.GetCourses(ctx)
}
