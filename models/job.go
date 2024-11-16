package models

import (
	"context"
	"sync"
)

// *************************
type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

type Job struct {
	ID       int
	Results  []JobResult
	Status   string
	Error    []string
	mux      sync.Mutex
	Ctx      context.Context
	CancelFn context.CancelFunc
}

type JobResult struct {
	StoreID   string `json:"store_id"`
	ImageURL  string `json:"image_url"`
	Perimeter int    `json:"perimeter"`
	Error     string `json:"error,omitempty"`
}

func (j *Job) MarkFailed(err string) {
	j.mux.Lock()
	defer j.mux.Unlock()
	j.Status = "failed"
	j.Error = append(j.Error, err)
	j.CancelFn()
}

func (j *Job) MarkCompleted() {
	j.mux.Lock()
	defer j.mux.Unlock()
	if j.Status != "failed" {
		j.Status = "completed"
	}
}

func (j *Job) AddResult(result JobResult) {
	j.mux.Lock()
	defer j.mux.Unlock()
	j.Results = append(j.Results, result)
}

func (j *Job) GetStatus() map[string]interface{} {
	j.mux.Lock()
	defer j.mux.Unlock()

	response := map[string]interface{}{
		"status": j.Status,
		"job_id": j.ID,
	}
	if j.Status == "failed" {
		response["error"] = j.Error
	} else if j.Status == "completed" {
		response["results"] = j.Results
	}
	return response
}
