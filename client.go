package auth

import (
	"sync"

	"google.golang.org/grpc"
)

const (
	serviceURL = "auth:80"
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
		serviceURL,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			ContextUnaryClientInterceptor(),
		),
	)

	cli := NewAuthClient(conn)
	Client = cli
	return cli
}
