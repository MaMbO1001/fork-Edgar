package app

import (
	"context"
	"github.com/gofrs/uuid"
)

type (
	Repo interface {
		ByAuthor(context.Context, uuid.UUID, int, int) (twit []Twit, total int, err error)
		Update(ctx context.Context, text string) (*Twit, error)
		Delete(ctx context.Context, id uuid.UUID) error
		Create(context.Context, string, uuid.UUID) (*Twit, error)
		GetTwitByID(ctx context.Context, Id uuid.UUID) (twit *Twit, err error)
	}
)
