package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ezrx/notification-service/graph/generated"
	"github.com/ezrx/notification-service/graph/model"
)

// RegisterToken is the resolver for the registerToken field.
func (r *mutationResolver) RegisterToken(ctx context.Context, input *model.RegisterToken) (bool, error) {
	err := r.FirebaseRepo.StoreToken(ctx, input.UserID, input.Token)
	return err == nil, err
}

// CreateNotification is the resolver for the createNotification field.
func (r *mutationResolver) CreateNotification(ctx context.Context, input model.NewNotification) (*model.Notification, error) {
	token, err := r.FirebaseRepo.GetToken(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	data, err := r.NotificationRepo.CreateNotification(ctx, &input)
	if err != nil {
		return nil, err
	}
	err = r.NotificationSender.SendNotification(ctx, token, data)
	return data, err
}

// ClearNotification is the resolver for the clearNotification field.
func (r *mutationResolver) ClearNotification(ctx context.Context, input *model.ClearNotificationIn) (bool, error) {
	userId := 1
	if input != nil && input.UserID != nil {
		userId = *input.UserID
	}
	err := r.SimpleNotificationRepo.MarkAllRead(ctx, userId)
	return err != nil, err
}

// CreateSimpleNotification is the resolver for the createSimpleNotification field.
func (r *mutationResolver) CreateSimpleNotification(ctx context.Context, input *model.NewSimpleNotification) (*model.SimpleNotification, error) {
	return r.SimpleNotificationRepo.CreateNotification(ctx, input)
}

// MarkRead is the resolver for the markRead field.
func (r *mutationResolver) MarkRead(ctx context.Context, input *model.MarkRead) (*model.SimpleNotification, error) {
	return r.SimpleNotificationRepo.MarkRead(ctx, input)
}

// NewOrder is the resolver for the newOrder field.
func (r *mutationResolver) NewOrder(ctx context.Context, input *model.NewOrder) (*model.OrderDetail, error) {
	return r.OrderRepo.NewOrder(ctx, input)
}

// Notifications is the resolver for the notifications field.
func (r *queryResolver) Notifications(ctx context.Context, count *int, offset *int) (*model.Notifications, error) {
	off := 0
	cnt := 10
	if offset != nil {
		off = *offset
	}
	if count != nil {
		cnt = *count
	}
	notifications, err := r.NotificationRepo.Notifications(ctx, off, cnt)
	if err != nil {
		return nil, err
	}
	nextOffset := off + cnt
	if len(notifications) < cnt {
		nextOffset = -1
	}
	return &model.Notifications{
		Notifications: notifications,
		NextOffset:    nextOffset,
	}, err
}

// SimpleNotifications is the resolver for the simpleNotifications field.
func (r *queryResolver) SimpleNotifications(ctx context.Context, input *model.NotificationList) (*model.SimpleNotifications, error) {
	off := 0
	cnt := 10

	if input != nil && input.Offset != nil {
		off = *input.Offset
	}
	if input != nil && input.Count != nil {
		cnt = *input.Count
	}
	langCode := "en"
	if input != nil && input.LangCode != nil {
		langCode = *input.LangCode
	}
	userId := 1
	if input != nil && input.UserID != nil {
		userId = *input.UserID
	}
	notifications, nextOffset, unreadCnt, err := r.SimpleNotificationRepo.Notifications(ctx, off, cnt, langCode, userId)
	if err != nil {
		return nil, err
	}
	return &model.SimpleNotifications{
		Notifications:               notifications,
		NextOffset:                  nextOffset,
		NumberOfUnreadNotifications: &unreadCnt,
	}, err
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, input *model.OrderListingQuery) (*model.OrderListResponse, error) {
	off := 0
	cnt := 10

	if input != nil && input.Offset != nil {
		off = *input.Offset
	}
	if input != nil && input.Count != nil {
		cnt = *input.Count
	}
	userId := 1
	if input != nil && input.UserID != nil {
		userId = *input.UserID
	}
	orders, nextOffset, err := r.OrderRepo.Orders(ctx, off, cnt, userId)
	if err != nil {
		return nil, err
	}
	return &model.OrderListResponse{
		Orders:     orders,
		NextOffset: nextOffset,
	}, err
}

// OrderDetail is the resolver for the orderDetail field.
func (r *queryResolver) OrderDetail(ctx context.Context, orderID int) (*model.OrderDetail, error) {
	return r.OrderRepo.OrderDetail(ctx, orderID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
