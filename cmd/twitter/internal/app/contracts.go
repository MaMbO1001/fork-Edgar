package app

import (
	"context"
	"github.com/gofrs/uuid"
)

type (
	Repo interface {
		ByAuthor(context.Context, uuid.UUID, int, int) (twit []Twit, total int, err error)
		Update(ctx context.Context, twit Twit) (*Twit, error)
		Delete(ctx context.Context, ID uuid.UUID) error
		Create(context.Context, Twit) (*Twit, error)
	}
)
