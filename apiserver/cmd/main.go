package main

import (
	"context"
	"fmt"
	"github.com/ashleywang1/capture-the-flag-apiserver/apiserver/server"
	"github.com/ashleywang1/capture-the-flag-apiserver/apiserver/server/handler"
)

func main() {
	ctx := context.Background()

	apiServer := server.NewGrpcServer(ctx, handler.NewHandler())

	if err := apiServer.Run(ctx); err != nil {
		fmt.Printf("Exited with error: %v\n", err)
	}
	fmt.Printf("Done\n")
}