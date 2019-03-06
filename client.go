package auth

import (
	"sync"

	"github.com/taeho-io/go-taeho/interceptor"
	"github.com/taeho-io/idl/gen/go/auth"
	"google.golang.org/grpc"
)

const (
	ServiceURL = "auth:80"
)

var (
	cm     = &sync.Mutex{}
	Client auth.AuthClient
)

func GetAuthClient() auth.AuthClient {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client
	}

	// We don't need to error here, as this creates a pool and connections
	// will happen later
	conn, _ := grpc.Dial(
		ServiceURL,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			interceptor.ContextUnaryClientInterceptor(),
		),
	)

	cli := auth.NewAuthClient(conn)
	Client = cli
	return cli
}
