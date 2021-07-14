package main

import (
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"

	"github.com/Tracking-SYS/tracking-go/utils/shutdown"
)

func main() {
	eg, ctx := errgroup.WithContext(shutdown.NewCtx())
	server, err := buildServer(ctx)
	if err != nil {
		fmt.Println("Can not create server: ", err)
		os.Exit(1)
	}
	eg.Go(func() error {
		return server.StartAll(ctx)
	})

	defer func() {
		server.CloseAll()
	}()

	if err := eg.Wait(); err != nil {
		fmt.Println("Exit Application: ", err)
		os.Exit(1)
	} else {
		fmt.Println("Close Application")
	}
}
