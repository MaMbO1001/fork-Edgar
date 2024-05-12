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
			(text, author_id)
		values 
		    ($1, $2) returning *`

		crea := twit{}
		err := db.GetContext(ctx, &crea, query, newTwit.Text, newTwit.AuthorID)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		twits = crea.convert()

		return nil
	})
	if err != nil {
		return nil, err
	}

	return twits, nil
}

func (r *Repo) Update(ctx context.Context, t app.Twit) (upTwit *app.Twit, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		updateTwit := convert(t)
		const query = `
		update twit
		set
		text = $1,
		updated_at = now()
		is_banned = $2
		where id = $3
		returning *`

		var upd twit
		err := db.GetContext(ctx, &upd, query, updateTwit.Text, updateTwit.ID, updateTwit.Is_banned)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		upTwit = upd.convert()

		return nil
	})
	if err != nil {
		return nil, err
	}

	return upTwit, nil
}

func (r *Repo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `
		delete
		from twit
		where id = $1 returning *`

		err := db.GetContext(ctx, &twit{}, query, id)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		return nil
	})
}

func (r *Repo) ByAuthor(ctx context.Context, authorId uuid.UUID, limit int, offset int) (t []app.Twit, total int, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `select * from twit where author_id = $1
		limit $2 offset $3`

		const gettotal = `select count(*) from twit where author_id = $1`

		res := []twit{}
		err := db.SelectContext(ctx, &res, query, authorId, limit, offset)
		if err != nil {
			return fmt.Errorf("db.SelectContext: %w", convertErr(err))
		}
		for _, ff := range res {
			t = append(t, *ff.convert())
		}

		err = db.GetContext(ctx, &total, gettotal, authorId)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		return nil
	})
	if err != nil {
		return t, 0, err
	}

	return t, total, nil
}

func (r *Repo) GetTwitByID(ctx context.Context, Id uuid.UUID) (twit *app.Twit, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = ` select * from twit where id = $1`
		err := db.GetContext(ctx, &twit, query, Id)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return twit, nil
}

func (r *Repo) GetTotalAllTwit(ctx context.Context) (total int, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `select count(*) from twit`

		err = db.GetContext(ctx, &total, query)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		return nil
	})
	return total, nil
}

func (r *Repo) GetAllTwitWithJopa(ctx context.Context) (t []app.Twit, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `select * FROM twit
		WHERE text LIKE '%jopa%'`
		res := []twit{}
		err = db.GetContext(ctx, &res, query)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		for _, ff := range res {
			t = append(t, *ff.convert())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}
