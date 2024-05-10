package app

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	SearchParams struct {
		AuthorID uuid.UUID
		Text     string
	}

	Twit struct {
		AuthorID  uuid.UUID
		ID        uuid.UUID
		Text      string
		CreatedAt time.Time
		UpdatedAt time.Time
		Is_banned bool
	}
)
