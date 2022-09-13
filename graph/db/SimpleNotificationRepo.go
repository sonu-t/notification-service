package db

import (
	"context"
	"errors"
	"github.com/ezrx/notification-service/graph/model"
	"sync/atomic"
	"time"
)

type SimpleNotificationRepo interface {
	CreateNotification(ctx context.Context, input *model.NewSimpleNotification) (*model.SimpleNotification, error)
	MarkRead(ctx context.Context, read *model.MarkRead) (*model.SimpleNotification, error)
	Notifications(ctx context.Context, offset int, count int, langCode string, userId int) ([]*model.SimpleNotification, int, error)
	MarkAllRead(ctx context.Context, userId int) error
}

type simpleNotificationRepo struct {
	counter       atomic.Int32
	notifications []*model.SimpleNotification
}

func (n *simpleNotificationRepo) MarkAllRead(ctx context.Context, userId int) error {
	for _, notification := range n.notifications {
		if notification.UserID == userId {
			notification.Status = true
		}
	}
	return nil
}

func (n *simpleNotificationRepo) MarkRead(ctx context.Context, read *model.MarkRead) (*model.SimpleNotification, error) {
	for _, notification := range n.notifications {
		if notification.ID == read.ID {
			notification.Status = true
			return notification, nil
		}
	}
	return nil, errors.New("notification id not found")
}

func (n *simpleNotificationRepo) Notifications(ctx context.Context, offset int, count int, langCode string, userId int) ([]*model.SimpleNotification, int, error) {
	total := len(n.notifications)
	if total == 0 {
		return nil, -1, nil
	}
	start := offset
	if start > total {
		return nil, -1, nil
	}

	var notifications []*model.SimpleNotification
	for _, notification := range n.notifications[start:] {
		if notification.UserID == userId && notification.LangCode == langCode {
			notifications = append(notifications, notification)
		}
		offset += 1
		if len(notifications) == count {
			break
		}
	}

	return notifications, offset, nil
}

func (n *simpleNotificationRepo) CreateNotification(ctx context.Context, input *model.NewSimpleNotification) (*model.SimpleNotification, error) {
	id := n.counter.Add(1)
	notification := &model.SimpleNotification{
		ID:               int(id),
		UserID:           input.UserID,
		OrderID:          input.OrderID,
		OrderType:        input.OrderType,
		OrderDescription: input.OrderDescription,
		HyperLink:        input.HyperLink,
		CreatedTime:      time.Now().Format(time.RFC3339),
		LangCode:         input.LangCode,
	}
	n.notifications = append(n.notifications, notification)
	return notification, nil
}

func NewSimpleNotificationRepo() SimpleNotificationRepo {
	return &simpleNotificationRepo{}
}
