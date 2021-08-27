package models

import (
	"time"
)

type Job struct {
	Id         int64          `db:"id, primarykey, autoincrement" json:"id"`
	ObjectId   int64          `db:"object_id" json:"object_id"`
	Sleep      *time.Duration `db:"sleep" json:"sleep"`
	CreatedAt  time.Time      `db:"created_at" json:"created_at"`
	StartedAt  *time.Time     `db:"started_at" json:"started_at"`
	FinishedAt *time.Time     `db:"finished_at" json:"finished_at"`
}

func NewJob(ObjectId int64) Job {
	return Job{
		ObjectId:  ObjectId,
		CreatedAt: time.Now(),
	}
}
