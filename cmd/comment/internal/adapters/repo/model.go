package repo

import (
	"github.com/ZergsLaw/back-template/cmd/comment/internal/app"
	"github.com/gofrs/uuid"
	"time"
)

type (
	comment struct {
		TwitID    uuid.UUID `db:"twit_id" json:"twit_id"`
		AuthorID  uuid.UUID `db:"author_id" json:"author_id"`
		ID        uuid.UUID `db:"id" json:"id"`
		Text      string    `db:"text" json:"text"`
		CreatedAt time.Time `db:"created_at" json:"createdAt"`
		UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	}
)

func convert(c app.Comment) *comment {
	return &comment{
		TwitID:    c.TwitID,
		AuthorID:  c.AuthorID,
		ID:        c.ID,
		Text:      c.Text,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func (c comment) convert() *app.Comment {
	return &app.Comment{
		TwitID:    c.TwitID,
		AuthorID:  c.AuthorID,
		ID:        c.ID,
		Text:      c.Text,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
