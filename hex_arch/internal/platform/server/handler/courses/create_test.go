package courses

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Create(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", CreateHandler(courseRepository))

	t.Run("given and invalid request it returns 400", func(t *testing.T) {

	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {

	})
}
