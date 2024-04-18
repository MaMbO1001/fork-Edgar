package repo

import (
	"database/sql"
	"errors"
	"github.com/ZergsLaw/back-template/cmd/twitter/internal/app"
	"github.com/lib/pq"
)

func convertErr(err error) error {
	var pqErr *pq.Error

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return app.ErrNotFound
	case errors.As(err, &pqErr):
		return constraint(pqErr)
	default:
		return err
	}
}

func constraint(pqErr *pq.Error) error {}
