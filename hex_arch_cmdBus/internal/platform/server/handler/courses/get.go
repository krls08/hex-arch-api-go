package courses

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/fetching"
)

type GetRespCourse struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func GetHandler(getCourseService fetching.CourseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("GetCoursesHandler start return")
		courses, err := getCourseService.GetAllCourses(ctx.Request.Context())
		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCopurseID), errors.Is(err, mooc.ErrEmptyCourseName),
				errors.Is(err, mooc.ErrEmptyDuration), errors.Is(err, mooc.ErrMissingHours):

				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return

			}
		}

		respCourses := make([]GetRespCourse, 0, len(courses))
		for _, v := range courses {
			course := GetRespCourse{
				ID:       v.ID().String(),
				Name:     v.Name().String(),
				Duration: v.Duration().String(),
			}
			respCourses = append(respCourses, course)
		}

		ctx.JSON(http.StatusOK, respCourses)

	}
}
