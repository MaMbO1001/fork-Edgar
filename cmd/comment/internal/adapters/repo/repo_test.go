//go:build integration

package repo_test

import (
	"github.com/ZergsLaw/back-template/cmd/comment/internal/app"
	"github.com/gofrs/uuid"
	"testing"
	"time"
)

func TestRepo_Smoke(t *testing.T) {
	t.Parallel()

	ctx, j, assert := start(t)

	com := app.Comment{
		TwitID:    uuid.Must(uuid.NewV4()),
		AuthorID:  uuid.Must(uuid.NewV4()),
		ID:        uuid.Nil,
		Text:      "Hello Jopa Edgara",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	comCreate, err := j.CreateCom(ctx, com)
	assert.NoError(err)
	assert.NotEmpty(comCreate.ID)
	assert.NotEmpty(comCreate.TwitID)
	assert.NotEmpty(comCreate.Text)
	assert.NotEmpty(comCreate.CreatedAt)
	assert.NotEmpty(comCreate.UpdatedAt)

	com.TwitID = comCreate.TwitID
	com.ID = comCreate.ID
	com.CreatedAt = comCreate.CreatedAt
	com.UpdatedAt = comCreate.UpdatedAt
	assert.Equal(com, *comCreate)

	coms := []app.Comment{com}
	getcom, total, err := j.GetCom(ctx, com.TwitID, 5, 0)
	assert.NoError(err)
	assert.Equal(coms, getcom)
	assert.Equal(1, total)
	assert.Len(getcom, 1)

	comupd, err := j.UpdateCom(ctx, com)
	assert.NoError(err)
	com.UpdatedAt = comupd.UpdatedAt
	assert.Equal(com, *comupd)
	assert.NotEmpty(comupd.ID)
	assert.NotEmpty(comupd.TwitID)

	bycom, err := j.GetComByID(ctx, com.ID)
	assert.NoError(err)
	assert.Equal(com, *bycom)
	assert.NotEmpty(bycom.ID)
	assert.NotEmpty(bycom.TwitID)
	assert.NotEmpty(bycom.AuthorID)

	err = j.DeleteCom(ctx, com.ID)
	assert.NoError(err)
}
