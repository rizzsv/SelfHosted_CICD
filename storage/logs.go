package storage

import (
	"selfhosted-ci/model"
	"time"

	"github.com/google/uuid"
)

var logs []model.BuildLog

func AddLog(repo string, status string, log string) model.BuildLog {
	newLog := model.BuildLog{
		ID:        uuid.NewString(),
		Repo:      repo,
		Status:    status,
		Timestamp: time.Now().Format(time.RFC3339),
		Log:       log,
	}
	logs = append(logs, newLog)
	return newLog
}

func GetAllLogs() []model.BuildLog {
	return logs
}

func GetLogByID(id string) *model.BuildLog {
	for _, l := range logs {
		if l.ID == id {
			return &l
		}
	}
	return nil
}
