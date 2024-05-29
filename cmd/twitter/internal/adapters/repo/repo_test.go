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
		IsBanned:  bool(false),
	}

	twitResCreate, err := r.Create(ctx, twit)
	assert.NoError(err)
	assert.NotEmpty(twitResCreate.ID)
	assert.NotEmpty(twitResCreate.Text)
	assert.NotEmpty(twitResCreate.CreatedAt)
	assert.NotEmpty(twitResCreate.UpdatedAt)
	assert.False(twitResCreate.IsBanned)

	twit.ID = twitResCreate.ID
	twit.IsBanned = twitResCreate.IsBanned
	twit.CreatedAt = twitResCreate.CreatedAt
	twit.UpdatedAt = twitResCreate.UpdatedAt
	assert.Equal(twit, *twitResCreate)

	twits := []app.Twit{twit}
	twitby, total, err := r.ByAuthor(ctx, twit.AuthorID, 10, 0)
	assert.NoError(err)
	assert.Equal(twits, twitby)
	assert.Len(twitby, 1)
	assert.Equal(1, total)
	twit.IsBanned = true

	twitResUpd, err := r.Update(ctx, twit)
	assert.NoError(err)
	twit.UpdatedAt = twitResUpd.UpdatedAt
	twit.IsBanned = twitResUpd.IsBanned
	assert.True(twitResUpd.IsBanned)
	assert.Equal(twit, *twitResUpd)
	assert.NotEmpty(twitResUpd.ID)

	twitts := []app.Twit{twit}
	twitby, total, err = r.ByAuthor(ctx, twit.AuthorID, 10, 0)
	assert.NoError(err)
	assert.Equal(twitts, twitby)
	assert.Len(twitby, 1)
	assert.Equal(1, total)

	getid, err := r.GetTwitByID(ctx, twitResCreate.ID)
	assert.NoError(err)
	assert.NotEmpty(getid.ID)

	err = r.Delete(ctx, twit.ID)
	assert.NoError(err)

	jopa, err := r.GetAllTwitWithJopa(ctx)
	assert.NoError(err)
	assert.Empty(jopa)
}
