package main

import (
	"fmt"

	"github.com/taeho-io/auth"
	"golang.org/x/net/context"
)

func main() {
	cli := auth.GetAuthClient()
	resp, err := cli.Auth(context.Background(), &auth.AuthRequest{UserId: 1234})
	fmt.Println(resp)
	fmt.Println(err)
}
