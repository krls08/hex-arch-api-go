package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/server"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	mysqlURI := fmt.Sprintf("url to connect...")
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)

	return srv.Run()
}
