package server

import (
	"fmt"
	"log"

	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/creating"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/platform/server/handler/courses"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/platform/server/handler/health"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/kit/command"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	creatingCourseService creating.CourseService
	commandBus            command.Bus
}

func New(host string, port uint, commandBus command.Bus, creatingCourseService creating.CourseService) Server {
	srv := Server{
		httpAddr: fmt.Sprintf(host + ":" + fmt.Sprint(port)),
		engine:   gin.New(),

		creatingCourseService: creatingCourseService,
		commandBus:            commandBus,
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
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
	s.engine.GET("/courses", courses.GetHandler(s.creatingCourseService))
}
