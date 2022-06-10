package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/fetching"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/platform/server/handler/courses"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/platform/server/handler/health"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/kit/command"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	// deps
	//creatingCourseService creating.CourseService -> this is inside command bus
	fetchingCourseService fetching.CourseService
	commandBus            command.Bus
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus,
	//creatingCourseService creating.CourseService,
	fetchingCourseService fetching.CourseService) (context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf(host + ":" + fmt.Sprint(port)),
		engine:   gin.New(),

		shutdownTimeout: shutdownTimeout,
		//creatingCourseService: creatingCourseService,
		fetchingCourseService: fetchingCourseService,

		commandBus: commandBus,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)
	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatal("Server shutdown", err)
		}
	}()

	<-ctx.Done()

	return srv.Shutdown(ctx)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}

func (s *Server) registerRoutes() {
	fmt.Println("Engine routes ...")
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
	s.engine.GET("/courses", courses.GetHandler(s.fetchingCourseService))
}
