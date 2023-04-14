package task

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/LepikovStan/plagiarismchecker/page"

	"cloud.google.com/go/pubsub"
	uuid "github.com/satori/go.uuid"
)

type Task struct {
	ID              string    `json:"id"`
	State           string    `json:"state"`
	OriginalArticle string    `json:"original_article"`
	ErrorMessage    string    `json:"error_message"`
	UserID          string    `json:"-"`
	Title           string    `json:"title"`
	CreatedAt       time.Time `json:created_at`
}

func New(article, userID string) Task {
	return Task{
		ID:              uuid.NewV4().String(),
		State:           "created",
		UserID:          userID,
		OriginalArticle: article,
		Title:           strings.Split(strings.Split(strings.Split(article, ".")[0], "?")[0], "!")[0],
	}
}

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
	Page       page.Page
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

type CheckPlagiarismTask struct {
	Task
	Article    string
	ArticleURL string `json:"article_url"`
	Page       page.Page
}

func CheckPlagiarismTaskFromGooglePubSubMessage(msg pubsub.Message) (CheckPlagiarismTask, error) {
	var (
		t CheckPlagiarismTask
	)

	if err := json.Unmarshal(msg.Data, &t); err != nil {
		return t, err
	}

	return t, nil
}
