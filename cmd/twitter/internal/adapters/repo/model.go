package repo

import (
	"github.com/ZergsLaw/back-template/cmd/twitter/internal/app"
	"github.com/gofrs/uuid"
	"time"
)

type (
	twit struct {
		AuthorID  uuid.UUID `db:"author_id" json:"author_id"`
		ID        uuid.UUID `db:"id" json:"id"`
		Text      string    `db:"text" json:"text"`
		CreatedAt time.Time `db:"created_at" json:"createdAt"`
		UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	}
)

func convert(t app.Twit) *twit {
	return &twit{
		AuthorID:  t.AuthorID,
		ID:        t.ID,
		Text:      t.Text,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func (t twit) convert() *app.Twit {
	return &app.Twit{
		AuthorID:  t.AuthorID,
		ID:        t.ID,
		Text:      t.Text,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
