package app

import (
	"context"
	"github.com/gofrs/uuid"
)

type Jopa interface {
	ByTwit(ctx context.Context, twitID uuid.UUID, limit int, offset int) (com []Comment, total int, err error)
	CreateComment(ctx context.Context, comment Comment) (*Comment, error)
	UpdateComment(ctx context.Context, comment Comment) (*Comment, error)
	DeleteComment(ctx context.Context, commentID uuid.UUID) error
	ByCommentID(ctx context.Context, commentID uuid.UUID) (*Comment, error)
}
