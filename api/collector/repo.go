package collector

import (
	"errors"

	"github.com/gopcua/opcua"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *opcua.Client
	logger log.Logger
}
