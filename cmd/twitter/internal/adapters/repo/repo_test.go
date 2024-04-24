//go:build integration

package repo_test

import (
	"github.com/ZergsLaw/back-template/cmd/twitter/internal/app"
	"github.com/gofrs/uuid"
	"testing"
	"time"
)

func TestRepo_Smoke(t *testing.T) {
	t.Parallel()

	ctx, r, assert := start(t)

	twit := app.Twit{
		AuthorID:  uuid.Must(uuid.NewV4()),
		ID:        uuid.Nil,
		Text:      "Hello Jopa Edgara",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	twitResCreate, err := r.Create(ctx, twit)
	assert.NoError(err)
	assert.NotEmpty(twitResCreate.ID)
	assert.NotEmpty(twitResCreate.CreatedAt)
	assert.NotEmpty(twitResCreate.UpdatedAt)

	twit.ID = twitResCreate.ID
	twit.CreatedAt = twitResCreate.CreatedAt
	twit.UpdatedAt = twitResCreate.UpdatedAt
	assert.Equal(twit, *twitResCreate)

	var twits []app.Twit
	twits = append(twits, twit)
	twitby, total, err := r.ByAuthor(ctx, twit.AuthorID, 10, 0)
	assert.NoError(err)
	assert.Equal(twits, twitby)
	assert.Len(twitby, 1)
	assert.Equal(1, total)

	twitResUpd, err := r.Update(ctx, twit)
	assert.NoError(err)
	twit.UpdatedAt = twitResUpd.UpdatedAt
	assert.Equal(twit, *twitResUpd)
	assert.NotEmpty(twitResUpd.ID)

	var twitts []app.Twit
	twitts = append(twitts, twit)
	twitby, total, err = r.ByAuthor(ctx, twit.AuthorID, 10, 0)
	assert.NoError(err)
	assert.Equal(twitts, twitby)
	assert.Len(twitby, 1)
	assert.Equal(1, total)

	err = r.Delete(ctx, twit.ID)
	assert.NoError(err)
}
