package task

import (
	"encoding/json"

	"cloud.google.com/go/pubsub"
)

type Task struct {
	ID           int    `json:"id"`
	OriginalText string `json:"original_text"`
}

type ProvideSERPTask struct {
	Task
	Keyword string `json:"keyword"`
}

func TaskToGooglePubSubMessage(t any) (*pubsub.Message, error) {
	taskBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return &pubsub.Message{
		Data: taskBytes,
	}, nil
}