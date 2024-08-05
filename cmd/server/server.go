package server

import (
	"github.com/poc/internal/grpc"
)

func Init() {
	grpc.InitServer()
}
