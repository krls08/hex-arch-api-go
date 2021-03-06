package server

import (
	"fmt"
	"log"

	mooc "github.com/krls08/hex-arch-api-go/hex_arch/internal"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/creating"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/server/handler/courses"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/server/handler/health"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	creatingCourseService creating.CourseService
}

func New(host string, port uint, creatingCourseService creating.CourseService, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		httpAddr: fmt.Sprintf(host + ":" + fmt.Sprint(port)),
		engine:   gin.New(),

		creatingCourseService: creatingCourseService,
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
	s.engine.POST("/courses", courses.CreateHandler(s.creatingCourseService))
	s.engine.GET("/courses", courses.GetHandler(s.creatingCourseService))
}
