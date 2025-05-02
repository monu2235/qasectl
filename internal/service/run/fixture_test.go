package run

import (
	"github.com/soulfulnerv/qasectl/internal/service/run/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

type fixture struct {
	client *mocks.Mockclient
}

func newFixture(t *testing.T) *fixture {
	ctrl := gomock.NewController(t)

	return &fixture{
		client: mocks.NewMockclient(ctrl),
	}
}

type baseArgs struct {
	err    error
	isUsed bool
}
