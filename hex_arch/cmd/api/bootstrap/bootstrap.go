package bootstrap

import "github.com/krls08/hex-arch-api-go/hex_arch/internal/platform/server"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)

	return srv.Run()
}
