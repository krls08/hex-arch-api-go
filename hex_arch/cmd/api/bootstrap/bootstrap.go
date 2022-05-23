package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/server"
	"github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "krls"
	dbPass = "123456"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "hex_arch_db"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)

	return srv.Run()
}
