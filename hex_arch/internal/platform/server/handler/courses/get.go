package courses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/krls08/hex-arch-api-go/hex_arch/internal"
)

type GetRespC struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func GetHandler(courseRepo mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("GetCoursesHandler start return")
		courses, err := courseRepo.GetCourses(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		respCourses := make([]GetRespC, 0, len(courses))
		for _, v := range courses {
			course := GetRespC{
				ID:       v.ID().String(),
				Name:     v.Name().String(),
				Duration: v.Duration().String(),
			}
			respCourses = append(respCourses, course)
		}

		ctx.JSON(http.StatusOK, respCourses)

	}
}
