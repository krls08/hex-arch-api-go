package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	mooc "github.com/krls08/hex-arch-api-go/hex_arch/internal"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

func (r *CourseRepository) GetCourses(ctx context.Context) ([]mooc.Course, error) {
	//courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	rows, err := r.db.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []mooc.Course{}
	sqlcourses := []sqlCourse{}
	for rows.Next() {
		var course sqlCourse
		if err := rows.Scan(&course.ID, &course.Name, &course.Duration); err != nil {
			return nil, err
		}
		sqlcourses = append(sqlcourses, course)
		dom_course, err := mooc.NewCourse(course.ID, course.Name, course.Duration)
		if err != nil {
			return courses, err
		}
		courses = append(courses, dom_course)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}
