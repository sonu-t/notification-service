package db

import (
	"context"
	"github.com/ezrx/notification-service/graph/model"
)

type NotificationRepo interface {
	Notifications(ctx context.Context, offset int, count int) ([]*model.Notification, error)
	CreateNotification(ctx context.Context, input *model.NewNotification) (*model.Notification, error)
}

type notificationRepo struct {
	notifications []*model.Notification
}

func (n *notificationRepo) Notifications(ctx context.Context, offset int, count int) ([]*model.Notification, error) {
	return n.notifications[offset : offset+count], nil
}

func (n *notificationRepo) CreateNotification(ctx context.Context, input *model.NewNotification) (*model.Notification, error) {
	notification := &model.Notification{
		ID:     input.ID,
		Title:  input.Title,
		Body:   input.Body,
		UserID: input.UserID,
		Read:   false,
	}
	n.notifications = append(n.notifications, notification)
	return notification, nil
}

func NewNotificationRepo() NotificationRepo {
	return &notificationRepo{}
}
