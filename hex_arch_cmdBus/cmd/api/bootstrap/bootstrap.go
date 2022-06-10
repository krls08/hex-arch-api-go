package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/creating"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/fetching"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/platform/bus/inmemory"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/platform/server"
	"github.com/krls08/hex-arch-api-go/hex_arch_cmdBus/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	//dbUser = "root"
	//dbPass = "123456"
	//dbHost = "localhost"
	dbUser          = "db_user"
	dbPass          = "password"
	dbHost          = "localhost"
	dbPort          = "3306"
	dbName          = "hex_arch_db"
	shutdownTimeout = 10 * time.Second
	dbTimeout       = 5 * time.Second
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
	)

	courseRepository := mysql.NewCourseRepository(db, dbTimeout)
	creatingCourseService := creating.NewCourseService(courseRepository)
	fetchingCourseService := fetching.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	//ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus, creatingCourseService, fetchingCourseService) //creatingCourseService will be inside command bus
	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus, fetchingCourseService)

	return srv.Run(ctx)
}
