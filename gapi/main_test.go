package gapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	db "github.com/zura-t/simplebank/db/sqlc"
	"github.com/zura-t/simplebank/utils"
	"github.com/zura-t/simplebank/worker"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}