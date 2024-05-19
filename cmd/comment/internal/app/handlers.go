package app

import (
	"context"
	"fmt"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/gofrs/uuid"
)

func (a *App) ByTwit(ctx context.Context, twitID uuid.UUID, limit int, offset int) (com []Comment, total int, err error) {
	com, total, err = a.jopa.ByTwit(ctx, twitID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("a.jopa.byTwit: %w", err)
	}
	return com, total, nil
}

func (a *App) CreateComment(ctx context.Context, session dom.Session, text string) (*Comment, error) {
	com, err := a.jopa.CreateComment(ctx, Comment{Text: text, AuthorID: session.UserID})
	if err != nil {
		return nil, fmt.Errorf("a.jopa.CreateComment: %w", err)
	}
	return com, nil
}

func (a *App) UpdateComment(ctx context.Context, session dom.Session, text string, comId uuid.UUID) (*Comment, error) {
	com, err := a.jopa.ByCommentID(ctx, comId)
	if err != nil {
		return nil, fmt.Errorf("a.jopa.ByCommentID: %w", err)
	}
	if session.UserID != com.AuthorID {
		return nil, ErrAccessDenied
	}
	com.Text = text
	cop, err := a.jopa.UpdateComment(ctx, Comment{Text: text, AuthorID: session.UserID})
	if err != nil {
		return cop, fmt.Errorf("a.jopa.UpdateComment: %w", err)
	}
	return com, nil
}

func (a *App) DeleteComment(ctx context.Context, session dom.Session, commentID uuid.UUID) error {
	com, err := a.jopa.ByCommentID(ctx, commentID)
	if err != nil {
		return fmt.Errorf("a.jopa.ByCommentID: %w", err)
	}
	if session.UserID != com.AuthorID {
		return ErrAccessDenied
	}
	err = a.jopa.DeleteComment(ctx, commentID)
	if err != nil {
		return fmt.Errorf("a.jopa.DeleteComment: %w", err)
	}
	return nil
}
