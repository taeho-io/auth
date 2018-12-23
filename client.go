package auth

import (
	"sync"

	"google.golang.org/grpc"
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

	serviceURL := "auth"

	// We don't need to error here, as this creates a pool and connections
	// will happen later
	conn, _ := grpc.Dial(
		serviceURL,
		grpc.WithInsecure(),
	)

	cli := NewAuthClient(conn)
	Client = cli
	return cli
}
