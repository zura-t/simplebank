package gapi

import (
	"fmt"

	db "github.com/zura-t/simplebank/db/sqlc"
	"github.com/zura-t/simplebank/pb"
	"github.com/zura-t/simplebank/token"
	"github.com/zura-t/simplebank/utils"
	"github.com/zura-t/simplebank/worker"
)

type Server struct {
	pb.UnimplementedMainServiceServer
	store      db.Store
	config     utils.Config
	tokenMaker token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewJwtMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can't create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
