package app

import (
	"context"
	"github.com/gofrs/uuid"
)

type (
	Repo interface {
		TwitByAuthors(context.Context, uuid.UUID) (*Twit, error)
		UpdateTwit(context.Context, Twit) (*Twit, error)
		DeleteTwit(ctx context.Context, id uuid.UUID) error
		CreateTwit(context.Context, Twit) (uuid.UUID, error)
	}
)
