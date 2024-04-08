package app

import (
	"context"
	"github.com/gofrs/uuid"
)

type (
	Repo interface {
		ByAuthor(context.Context, uuid.UUID) (total int, twit []Twit, err error)
		Update(ctx context.Context, twit Twit) (*Twit, error)
		Delete(ctx context.Context, ID uuid.UUID) error
		Create(context.Context, Twit) (*Twit, error)
	}
)
