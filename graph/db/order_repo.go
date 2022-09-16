package db

import (
	"context"
	"errors"
	"github.com/ezrx/notification-service/graph/model"
	"sync"
	"time"
)

type OrderRepo interface {
	Orders(ctx context.Context, offset, count, userId int) ([]*model.OrderDetail, int, error)
	NewOrder(ctx context.Context, input *model.NewOrder) (*model.OrderDetail, error)
	OrderDetail(ctx context.Context, orderId int) (*model.OrderDetail, error)
}

type orderRepo struct {
	sync.Mutex
	orderId int
	orders  []*model.OrderDetail
}

func (o *orderRepo) OrderDetail(ctx context.Context, orderId int) (*model.OrderDetail, error) {
	for _, order := range o.orders {
		if order.OrderID == orderId {
			return order, nil
		}
	}
	return nil, errors.New("order id not found")
}

func (o *orderRepo) Orders(ctx context.Context, offset, count, userId int) ([]*model.OrderDetail, int, error) {
	if offset < 1 {
		offset = 0
	}
	if offset >= len(o.orders) || len(o.orders) == 0 {
		return nil, -1, nil
	}
	var orders []*model.OrderDetail
	for offset < len(o.orders) {
		order := o.orders[offset]
		if order.UserID == userId {
			orders = append(orders, order)
		}
		offset += 1
		if len(orders) == count {
			break
		}
	}
	if offset == len(orders) {
		offset = -1
	}
	return orders, offset, nil
}

func (o *orderRepo) NewOrder(ctx context.Context, input *model.NewOrder) (*model.OrderDetail, error) {
	o.Lock()
	defer o.Unlock()
	o.orderId += 1
	order := &model.OrderDetail{
		OrderID:   o.orderId,
		OrderDate: time.Now().Format(time.RFC3339),
		TotalCost: input.TotalCost,
		UserID:    input.UserID,
		Status:    input.Status,
	}
	for _, product := range input.Products {
		order.Products = append(order.Products, &model.ProductDetail{
			Name:         product.Name,
			Count:        product.Count,
			PricePerUnit: product.PricePerUnit,
			ThumbnailURL: product.ThumbnailURL,
		})
	}
	orders := []*model.OrderDetail{order}
	o.orders = append(orders, o.orders...)
	return order, nil
}

func NewOrderRepo() OrderRepo {
	return &orderRepo{}
}
