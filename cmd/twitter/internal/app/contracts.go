package app

import (
	"context"
	"github.com/gofrs/uuid"
)

type (
	Repo interface {
		ByAuthor(context.Context, uuid.UUID) (*Twit, error)
		Update(ctx context.Context, ID uuid.UUID, text string) (*Twit, error)
		Delete(ctx context.Context, ID uuid.UUID) error
		Create(context.Context, Twit) (*Twit, error)
	}
)
