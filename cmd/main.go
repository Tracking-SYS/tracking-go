package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"

	"github.com/Tracking-SYS/tracking-go/utils/shutdown"
)

func main() {
	eg, ctx := errgroup.WithContext(shutdown.NewCtx())
	err := godotenv.Load("env.yaml")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server, err := buildServer(ctx)
	if err != nil {
		fmt.Println("Can not create server: ", err)
		os.Exit(1)
	}
	eg.Go(func() error {
		return server.StartAll(ctx)
	})

	defer func() {
		err := server.CloseAll()
		if err != nil {
			fmt.Println("Server Close has problem", err)
		}
	}()

	if err := eg.Wait(); err != nil {
		fmt.Println("Exit Application: ", err)
		os.Exit(1)
	} else {
		fmt.Println("Close Application")
	}
}
