package repo

import (
	"database/sql"
	"errors"
	"github.com/ZergsLaw/back-template/cmd/comment/internal/app"
)

func convertErr(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return app.ErrNotFound
	default:
		return err
	}
}
