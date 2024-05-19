package repo

import (
	"context"
	"fmt"
	"github.com/ZergsLaw/back-template/cmd/comment/internal/app"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database"
	"github.com/sipki-tech/database/connectors"
	"github.com/sipki-tech/database/migrations"
)

type (
	Config struct {
		Cockroach  connectors.CockroachDB
		MigrateDir string
		Driver     string
	}

	Jopa struct {
		sql *database.SQL
	}
)

func New(ctx context.Context, reg *prometheus.Registry, namespace string, cfg Config) (*Jopa, error) {
	const subsystem = "jopa"
	m := database.NewMetrics(reg, namespace, subsystem, new(app.Jopa))

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

	return &Jopa{
		sql: conn,
	}, nil
}

func (j *Jopa) Close() error {
	return j.sql.Close()
}

func (j *Jopa) CreateCom(ctx context.Context, c app.Comment) (com *app.Comment, err error) {
	err = j.sql.NoTx(func(db *sqlx.DB) error {
		newcom := convert(c)
		const query = `
insert into
com
(text, author_id)
values ($1, $2) returning *`
		crea := comment{}
		err = db.GetContext(ctx, &crea, query, newcom.Text, newcom.AuthorID)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		com = crea.convert()
		return nil
	})
	if err != nil {
		return nil, err
	}

	return com, nil
}

func (j *Jopa) UpdateCom(ctx context.Context, c app.Comment) (com *app.Comment, err error) {
	err = j.sql.NoTx(func(db *sqlx.DB) error {
		newcom := convert(c)
		const query = `
update com
set
text = $1,
updated_at = now()
where id = $2
returning *`
		upd := comment{}
		err = db.GetContext(ctx, &upd, query, newcom.Text, newcom.ID)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		com = upd.convert()
		return nil
	})
	if err != nil {
		return nil, err
	}
	return com, nil
}

func (j *Jopa) DeleteCom(ctx context.Context, id uuid.UUID) error {
	return j.sql.NoTx(func(db *sqlx.DB) error {
		const query = `
delete
from com
where id = $1
returning  *`
		err := db.GetContext(ctx, &comment{}, query, id)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		return nil
	})
}

func (j *Jopa) GetCom(ctx context.Context, twitid uuid.UUID, limit int, offset int) (com []app.Comment, total int, err error) {
	err = j.sql.NoTx(func(db *sqlx.DB) error {
		const query = `
select * from com where twit_id = $1
                  limit $2 offset $3`

		const tottal = `
select count(*) from com where twit_id = $1`
		ps := []comment{}
		err = db.SelectContext(ctx, &ps, query, twitid, limit, offset)
		if err != nil {
			return fmt.Errorf("db.SelectContext: %w", convertErr(err))
		}
		for _, p := range ps {
			com = append(com, *p.convert())
		}
		err = db.GetContext(ctx, &total, tottal, twitid)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return com, total, nil
}

func (j *Jopa) GetComByID(ctx context.Context, id uuid.UUID) (comm *app.Comment, err error) {
	com := comment{}
	err = j.sql.NoTx(func(db *sqlx.DB) error {
		const query = `SELECT * FROM com WHERE id = $1`
		err := db.GetContext(ctx, &com, query, id)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return com.convert(), nil
}
