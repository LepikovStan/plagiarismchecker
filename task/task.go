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

type r struct{}

func TaskToGooglePubSubMessage(t interface{}) (*pubsub.Message, error) {
	taskBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return &pubsub.Message{
		Data: taskBytes,
	}, nil
}

func ProvideSERPTaskFromGooglePubSubMessage(msg pubsub.Message) (ProvideSERPTask, error) {
	var (
		t ProvideSERPTask
	)

	if err := json.Unmarshal(msg.Data, &t); err != nil {
		return t, err
	}

	return t, nil
}
