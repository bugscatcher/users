package main

import (
	"fmt"
	"net"

	"github.com/bugscatcher/users/application"
	"github.com/bugscatcher/users/config"
	"github.com/bugscatcher/users/grpcHelper"
	"github.com/bugscatcher/users/server/grpc/users"
	"github.com/bugscatcher/users/services"
	"github.com/rs/zerolog/log"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.New()
	if err != nil {
		xErr := xerrors.Errorf("while reading config: %w", err)
		log.Fatal().Err(xErr).Msg("Read config")
	}

	app, err := application.New(&conf)
	if err != nil {
		xErr := xerrors.Errorf("while creating app: %w", err)
		log.Fatal().Err(xErr).Msg("Create app")
	}

	defer app.Close()

	go startGRPCServer(&conf.PublicGRPCServer, app)

}

func startGRPCServer(grpcConf *grpcHelper.ServerConf, app *application.App) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcConf.Port))
	if err != nil {
		xErr := xerrors.Errorf("while listening GRPC port: %w", err)
		log.Fatal().Err(xErr).Msgf("Listen GRPC port: %s", grpcConf.Addr())
	}
	s := grpc.NewServer()

	handler := users.New(app)
	services.RegisterUsersServiceServer(s, handler)
	if err := s.Serve(lis); err != nil {
		xErr := xerrors.Errorf("while serving: %w", err)
		log.Fatal().Err(xErr).Msg("Serve")
	}
	log.Info().Msgf("GRPC server started %s", grpcConf.Addr())
}
