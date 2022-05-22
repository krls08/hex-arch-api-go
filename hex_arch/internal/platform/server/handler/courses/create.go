package courses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/krls08/hex-arch-api-go/hex_arch/internal"
)

type createRequest struct {
	ID        string `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Durantion string `json:"duration" binding:"required"`
}

func CreateHandler(courseRepo mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("createHandler start return")
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println("validation done")

		course := mooc.NewCourse(req.ID, req.Name, req.Durantion)
		_ = course
		if err := courseRepo.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.Status(http.StatusCreated)

	}
}