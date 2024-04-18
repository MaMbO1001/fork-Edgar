package repo

import (
	"context"
	"fmt"
	"github.com/ZergsLaw/back-template/cmd/twitter/internal/app"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database"
	"github.com/sipki-tech/database/connectors"
	"github.com/sipki-tech/database/migrations"
)

var _ app.Repo = &Repo{}

type (
	Config struct {
		Cockroach  connectors.CockroachDB
		MigrateDir string
		Driver     string
	}

	Repo struct {
		sql *database.SQL
	}
)

func New(ctx context.Context, reg *prometheus.Registry, namespace string, cfg Config) (*Repo, error) {
	const subsystem = "repo"
	m := database.NewMetrics(reg, namespace, subsystem, new(app.Repo))

	returnErrs := []error{
		app.ErrNotFound,
	}

	migrates, err := migrations.Parse(cfg.MigrateDir)
	if err != nil {
		return nil, fmt.Errorf("migrations.Parse: %w", err)
	}

	err = migrations.Run(ctx, cfg.Driver, &cfg.Cockroach, migrations.Up, migrates)
	if err != nil {
		return nil, fmt.Errorf("migrations.Run: %w", err)
	}

	conn, err := database.NewSQL(ctx, cfg.Driver, database.SQLConfig{
		Metrics:    m,
		ReturnErrs: returnErrs,
	}, &cfg.Cockroach)
	if err != nil {
		return nil, fmt.Errorf("librepo.NewCockroach: %w", err)
	}

	return &Repo{
		sql: conn,
	}, nil
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) Create(ctx context.Context, t app.Twit) (twits *app.Twit, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		newTwit := convert(t)
		const query = `
		insert into
		twit
			(text)
		values 
		    ($1)
		    returning id
		`
		err := db.GetContext(ctx, &twits, query, newTwit.Text)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		return nil
	})
	if err != nil {
		return twits, err
	}

	return twits, nil
}

func (r *Repo) Update(ctx context.Context, t app.Twit) (upTwit *app.Twit, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		updateTwit := convert(t)
		const query = `
		update twits
		set
		text = $1,
		updatedAt = now()
		where id = $2
		returning *`

		var upd twit
		err := db.GetContext(ctx, &upd, query, updateTwit.Text)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		upTwit = upd.convert()

		return nil
	})
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (r *Repo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `
		delete
		from twits
		where id = $1 returning *`

		err := db.GetContext(ctx, &twit{}, query, id)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		return nil
	})
}

func (r *Repo) ByAuthor(ctx context.Context, AuthorId uuid.UUID) (total int, t []app.Twit, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `select * from twits where authorId = any($1)`

		const getTotal = `select count(*) over() as total from twits where authorId = $1`

		res := twit{}
		err := db.GetContext(ctx, res, query, getTotal, AuthorId)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		return nil
	})
	if err != nil {
		return 0, t, err
	}

	return total, nil, err
}
