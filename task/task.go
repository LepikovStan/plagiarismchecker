package task

import (
	"encoding/json"

	"cloud.google.com/go/pubsub"
)

type Task struct {
	ID           int    `json:"id"`
	OriginalText string `json:"original_text"`
}

t

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

type ProvideSERPTask struct {
	Task
	Keyword string `json:"keyword"`
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

type ArticleScraperTask struct {
	Task
	ArticleURL string `json:"article_url"`
}

func ContentScraperTaskFromGooglePubSubMessage(msg pubsub.Message) (ArticleScraperTask, error) {
	var (
		t ArticleScraperTask
	)

	if err := json.Unmarshal(msg.Data, &t); err != nil {
		return t, err
	}

	return t, nil
}
