package app

import (
	"context"
	"fmt"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/gofrs/uuid"
)

func (a *App) TwitByAuthor(ctx context.Context, authorID uuid.UUID, limit int, offset int) (twit []Twit, total int, err error) {
	twit, total, err = a.repo.ByAuthor(ctx, authorID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("a.repo.ByAuthor: %w", err)
	}
	twit = append(twit, Twit{Text: "jopa"})
	return twit, total, nil
}

func (a *App) CreateTwit(ctx context.Context, session dom.Session, text string) (*Twit, error) {
	j, err := a.repo.Create(ctx, Twit{Text: text, AuthorID: session.UserID})
	if err != nil {
		return nil, fmt.Errorf("a.repo.Create: %w", err)
	}
	return j, nil
}

// ЗАМЕТКА
func (a *App) UpdateTwit(ctx context.Context, session dom.Session, text string, id uuid.UUID) (*Twit, error) {
	t, err := a.repo.GetTwitByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetTwitByID: %w", err)
	}
	if session.UserID != t.AuthorID || !session.Status.IsAdmin() {
		return nil, ErrAccessDenied
	}
	t.Text = text
	j, err := a.repo.Update(ctx, *t)
	if err != nil {
		return j, fmt.Errorf("a.repo.Update: %w", err)
	}
	return t, nil
}

func (a *App) DeleteTwit(ctx context.Context, session dom.Session, id uuid.UUID) error {
	t, err := a.repo.GetTwitByID(ctx, id)
	if err != nil {
		return fmt.Errorf("a.repo.GetTwitByID: %w", err)
	}
	if session.UserID != t.AuthorID || !session.Status.IsAdmin() {
		return ErrAccessDenied
	}
	err = a.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("a.repo.Delete: %w", err)
	}
	return nil
}

func (a *App) AdminCreateTwit(ctx context.Context, session dom.Session, text string, authorID uuid.UUID) (*Twit, error) {
	if !session.Status.IsAdmin() {
		return nil, ErrAccessDenied
	}
	j, err := a.repo.Create(ctx, Twit{Text: text, AuthorID: authorID})
	if err != nil {
		return nil, fmt.Errorf("a.repo.Create: %w", err)
	}
	return j, nil
}

func (a *App) AdminBannedTwit(ctx context.Context, session dom.Session, id uuid.UUID) (*Twit, error) {
	if !session.Status.IsAdmin() || !session.Status.IsSupport() {
		return nil, ErrAccessDenied
	}
	t, err := a.repo.GetTwitByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetTwitByID: %w", err)
	}
	j, err := a.repo.Update(ctx, *t)
	if err != nil {
		return nil, fmt.Errorf("a.repo.Update: %w", err)
	}
	return j, nil
}
