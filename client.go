package auth

import (
	"sync"

	"github.com/taeho-io/taeho-go/interceptor"
	"google.golang.org/grpc"
)

const (
	ServiceURL = "auth:80"
)

var (
	cm     = &sync.Mutex{}
	Client AuthClient
)

func GetAuthClient() AuthClient {
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

	cli := NewAuthClient(conn)
	Client = cli
	return cli
}
