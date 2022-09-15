package graph

import (
	"github.com/ezrx/notification-service/graph/db"
	"github.com/ezrx/notification-service/graph/firebase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	NotificationRepo       db.NotificationRepo
	FirebaseRepo           db.FirebaseRepo
	SimpleNotificationRepo db.SimpleNotificationRepo
	NotificationSender     firebase.NotificationSender
	OrderRepo              db.OrderRepo
}
