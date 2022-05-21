package health

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("CheckHandler")
		ctx.String(http.StatusOK, "everything is ok!")
	}

}
