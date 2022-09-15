package db

import (
	"context"
	"errors"
	"github.com/ezrx/notification-service/graph/model"
	"sync"
	"time"
)

type SimpleNotificationRepo interface {
	CreateNotification(ctx context.Context, input *model.NewSimpleNotification) (*model.SimpleNotification, error)
	MarkRead(ctx context.Context, read *model.MarkRead) (*model.SimpleNotification, error)
	Notifications(ctx context.Context, offset int, count int, langCode string, userId int) ([]*model.SimpleNotification, int, int, error)
	MarkAllRead(ctx context.Context, userId int) error
}

type simpleNotificationRepo struct {
	sync.Mutex
	counter       int
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

func (n *simpleNotificationRepo) Notifications(ctx context.Context, offset int, count int, langCode string, userId int) (notifications []*model.SimpleNotification, nextOffset int, unReadCount int, err error) {
	total := len(n.notifications)
	if total == 0 || offset < 0 || offset >= total {
		return nil, -1, 0, nil
	}
	for _, notification := range n.notifications {
		if !notification.Status {
			unReadCount += 1
		}
	}
	for offset < total {
		notification := n.notifications[offset]
		if notification.UserID == userId && notification.LangCode == langCode {
			notifications = append(notifications, notification)
		}
		offset += 1
		if len(notifications) == count {
			break
		}
	}
	return
}

func (n *simpleNotificationRepo) CreateNotification(ctx context.Context, input *model.NewSimpleNotification) (*model.SimpleNotification, error) {
	n.Lock()
	defer n.Unlock()
	n.counter += 1
	notification := &model.SimpleNotification{
		ID:               n.counter,
		UserID:           input.UserID,
		OrderID:          input.OrderID,
		OrderType:        input.OrderType,
		OrderDescription: input.OrderDescription,
		HyperLink:        input.HyperLink,
		CreatedTime:      time.Now().Format(time.RFC3339),
		LangCode:         input.LangCode,
	}
	notifications := []*model.SimpleNotification{notification}
	n.notifications = append(notifications, n.notifications...)
	return notification, nil
}

func NewSimpleNotificationRepo() SimpleNotificationRepo {
	return &simpleNotificationRepo{}
}
