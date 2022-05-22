package server

import (
	"fmt"
	"log"

	mooc "github.com/krls08/hex-arch-api-go/hex_arch/internal"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/server/handler/courses"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/server/handler/health"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	courseRepo mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		httpAddr: fmt.Sprintf(host + ":" + fmt.Sprint(port)),
		engine:   gin.New(),

		courseRepo: courseRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	fmt.Println("Engine routes ...")
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepo))
}
