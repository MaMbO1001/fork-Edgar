package app_test

import (
	"context"
	"github.com/ZergsLaw/back-template/cmd/twitter/internal/app"
	"github.com/ZergsLaw/back-template/internal/testhelper"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

type mocks struct {
	repo *MockRepo
}

func start(t *testing.T) (context.Context, *app.App, *mocks, *require.Assertions) {
	t.Helper()
	ctrl := gomock.NewController(t)

	mockRepo := NewMockRepo(ctrl)

	module := app.New(mockRepo)

	mocks := &mocks{
		repo: mockRepo,
	}

	return testhelper.Context(t), module, mocks, require.New(t)
}
