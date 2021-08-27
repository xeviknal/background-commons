package database

import (
	"testing"

	"github.com/xeviknal/background-commons/models"

	"github.com/stretchr/testify/assert"
)

func TestJobInsert(t *testing.T) {
	asserter := assert.New(t)

	db := GetDb()
	job := models.NewJob(10)
	err := db.Insert(&job)
	asserter.NoError(err)

	// Asserting default values
	// Id is filled automatically
	asserter.NotEmpty(job.Id)
	// Object Id is the same as passed
	asserter.Equal(job.ObjectId, int64(10))
	// CreatedAt is filled
	asserter.NotNil(job.CreatedAt)

	// Background job timestamps are nil
	asserter.Nil(job.StartedAt)
	asserter.Nil(job.FinishedAt)
	asserter.Nil(job.QueuedAt)

	// Status field might be null
	asserter.Nil(job.Status)
}
