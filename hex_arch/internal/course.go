package mooc

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

var ErrInvalidCopurseID = errors.New("invalid Course ID")
var ErrEmptyCourseName = errors.New("the field Course Name can not be empty")
var ErrEmptyDuration = errors.New("the field Duration can not be empty")
var ErrMissingHours = errors.New("duration must contain hours")

// CourseID represents the course uinique identifier
type CourseID struct {
	value string
}

func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCopurseID, value)
	}
	return CourseID{
		value: v.String(),
	}, nil
}

// String method in CourseID, returns the CourseID type into string.
func (id CourseID) String() string {
	return id.value
}

// CourseName represents the course name.
type CourseName struct {
	value string
}

// NewCourseName instantiate VO for CourseName
func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}

	return CourseName{
		value: value,
	}, nil
}

// String type converts the CourseName into string.
func (name CourseName) String() string {
	return name.value
}

// CourseDuration represents the course duration.
type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyDuration
	}
	// must contain the hours
	isOk, err := regexp.MatchString("hours", value)
	if err != nil {
		return CourseDuration{}, err
	}
	if !isOk {
		return CourseDuration{}, ErrMissingHours
	}

	return CourseDuration{
		value: value,
	}, nil
}

// String type converts the CourseDuration into string.
func (duration CourseDuration) String() string {
	return duration.value
}

// CourseRepo defines the xepected behavieour form a course storage
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
	GetCourses(ctx context.Context) ([]Course, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// Course is the data structure that represents a course.
type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new course
func NewCourse(id, name, duration string) (Course, error) {
	idVO, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}
	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil
}

// ID returns the course unique identifier.
func (c Course) ID() CourseID {
	return c.id
}

// Name returns the course name.
func (c Course) Name() CourseName {
	return c.name
}

// Duration returns the course duration.
func (c Course) Duration() CourseDuration {
	return c.duration
}
