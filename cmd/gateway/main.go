package main

import (
	"net"

	"github.com/kazmerdome/muzz/internal/actor/db"
	"github.com/kazmerdome/muzz/internal/module/decision"
	"github.com/kazmerdome/muzz/internal/module/explore"
	explore_grpc "github.com/kazmerdome/muzz/internal/module/explore/explore-grpc"
	"github.com/kazmerdome/muzz/internal/util/config"
	"github.com/kazmerdome/muzz/internal/util/logger"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Init Config
	c := config.NewConfig()
	err := c.LoadConfigFile(".", "env", ".env")
	if err != nil {
		log.
			Error().
			Msg(err.Error())
	}

	// Init Logger
	logger.InitLogger(
		c.GetString("LOG_LEVEL"),
		c.GetString("ENVIRONMENT"),
	)

	// Init Database
	pdb := db.NewPostgresDB(
		c.GetString("POSTGRES_DATABASE"),
		c.GetString("POSTGRES_URI"),
		c.GetBool("POSTGRES_IS_SSL_DISABLED"),
	).Connect()
	defer pdb.Disconnect()

	// Load modules
	decisionModule := decision.NewDecisionModule(pdb)
	exploreModule := explore.NewExploreModule(decisionModule.GetRepository())

	// Add Controller(s) to grpc Server
	server := grpc.NewServer()
	explore_grpc.RegisterExploreServiceServer(server, exploreModule.GetController())
	reflection.Register(server)

	// Start grpc server
	con, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatal().Err(err)
	}
	err = server.Serve(con)
	if err != nil {
		log.Fatal().Err(err)
	}
}
