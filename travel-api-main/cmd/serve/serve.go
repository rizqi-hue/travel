/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package serve

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/travel-api-main/api"
	def "github.com/travel-api-main/cmd/config"
	"github.com/travel-api-main/pkg/config"
	"github.com/travel-api-main/pkg/database"
	"github.com/travel-api-main/pkg/logger"
)

// serveCmd represents the serve command
var (
	Config string

	ServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "configration & connection",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
)

func startServer() {
	ctx := context.Background()

	if err := config.Load(def.DefaultConfig, Config); err != nil {
		log.Fatal(err)
	}

	logger.Configure()
	database.PostgresConnection(ctx)

	// Start RPC server in a goroutine
	go NewGrpcServer(ctx)

	// Start Fiber server in a goroutine
	go api.NewFiberServer()

	select {}
}
