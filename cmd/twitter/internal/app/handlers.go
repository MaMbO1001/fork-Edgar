package app

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
)

func (a *App) TwitByAuthor(ctx context.Context, authorID uuid.UUID, limit int, offset int) (twit []Twit, total int, err error) {
	twit, total, err = a.repo.ByAuthor(ctx, authorID, limit, offset)
	if err != nil {
		return twit, total, fmt.Errorf("a.repo.ByAuthor: %w", err)
	}
	return twit, total, nil
}

func (a *App) CreateTwit(ctx context.Context, twit Twit) (*Twit, error) {
	j, err := a.repo.Create(ctx, twit)
	if err != nil {
		return nil, fmt.Errorf("a.repo.Create: %w", err)
	}
	return j, nil
}

func (a *App) UpdateTwit(ctx context.Context, twit Twit) (*Twit, error) {
	j, err := a.repo.Update(ctx, twit)
	if err != nil {
		return nil, fmt.Errorf("a.repo.Update: %w", err)
	}
	return j, nil
}

func (a *App) DeleteTwit(ctx context.Context, ID uuid.UUID) error {
	err := a.repo.Delete(ctx, ID)
	if err != nil {
		return fmt.Errorf("a.repo.Delete: %w", err)
	}
	return nil
}
