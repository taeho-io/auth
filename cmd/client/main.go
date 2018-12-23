package main

import (
	"fmt"

	"github.com/taeho-io/auth"
	"golang.org/x/net/context"
)

func main() {
	cli := auth.GetAuthClient()
	resp, err := cli.Auth(context.Background(), &auth.AuthRequest{UserId: "taeho"})
	fmt.Println(resp)
	fmt.Println(err)
}
