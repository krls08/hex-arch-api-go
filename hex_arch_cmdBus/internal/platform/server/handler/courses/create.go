package courses

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/krls08/hex-arch-api-go/hex_arch/internal"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/creating"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(creatingCourseService creating.CourseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("createHandler start return")
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println("validation done")

		err := creatingCourseService.CreateCourse(ctx, req.ID, req.Name, req.Duration)
		// Errors handling
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
		ctx.Status(http.StatusCreated)
	}

}
