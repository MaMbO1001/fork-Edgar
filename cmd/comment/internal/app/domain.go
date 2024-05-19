package app

import (
	"github.com/gofrs/uuid"
	"time"
)

type Comment struct {
	TwitID    uuid.UUID
	AuthorID  uuid.UUID
	ID        uuid.UUID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
