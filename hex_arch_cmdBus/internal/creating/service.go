package creating

import (
	"context"
	"fmt"

	mooc "github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/kit/event"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService
type CourseService struct {
	courseRepository mooc.CourseRepository
	eventBus         event.Bus
}

// NewCourseService returns de default Service interface implementation
func NewCourseService(courseRepository mooc.CourseRepository, eventBus event.Bus) CourseService {
	return CourseService{
		courseRepository: courseRepository,
		eventBus:         eventBus,
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
