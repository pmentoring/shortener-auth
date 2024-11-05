package grpc

import (
	"flag"
	pb "github.com/pmentoring/shortener-protoc/gen/go/shortener"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "go-app:8000", "the address to connect to")
)

func NewGrpc() pb.ShortenerClient {
	flag.Parse()

	var opts []grpc.DialOption
	conn, err := grpc.NewClient(*addr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewShortenerClient(conn)

	return client
}
