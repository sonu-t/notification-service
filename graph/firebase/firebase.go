package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/ezrx/notification-service/graph/model"
	"google.golang.org/api/option"
	"log"
)

type NotificationSender interface {
	SendNotification(ctx context.Context, token string, notification *model.Notification) error
}

type notificationSender struct {
	app    *firebase.App
	client *messaging.Client
}

func (n *notificationSender) SendNotification(ctx context.Context, token string, notification *model.Notification) error {
	message := messaging.Message{
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Body,
		},
		Token: token,
	}
	_, err := n.client.Send(ctx, &message)
	return err
}

func NewNotificationSender() NotificationSender {
	n := &notificationSender{}
	opt := option.WithCredentialsFile("configs/cred.json")
	app, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: "test-project-115e6"}, opt)
	if err != nil {
		log.Fatal("error initializing firebase app", err)
	}
	n.app = app
	n.client, err = n.app.Messaging(context.Background())
	if err != nil {
		log.Fatal("error initializing messaging client", err)
	}
	return n
}
