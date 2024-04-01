package app

import "github.com/gofrs/uuid"

type (
	SearchParams struct {
		AuthorID uuid.UUID
		Text     string
	}

	Twit struct {
		AuthorID uuid.UUID
		ID       uuid.UUID
		Text     string
	}
)
