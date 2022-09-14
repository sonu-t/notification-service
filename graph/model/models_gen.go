// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ClearNotificationIn struct {
	UserID *int `json:"userId"`
}

type MarkRead struct {
	ID int `json:"id"`
}

type NewNotification struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type NewSimpleNotification struct {
	LangCode         string `json:"langCode"`
	UserID           int    `json:"userId"`
	OrderID          int    `json:"orderId"`
	OrderType        string `json:"orderType"`
	OrderDescription string `json:"orderDescription"`
	HyperLink        string `json:"hyperLink"`
}

type Notification struct {
	LangCode string `json:"langCode"`
	ID       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	UserID   int    `json:"userId"`
	Read     bool   `json:"read"`
}

type NotificationList struct {
	Count    *int    `json:"count"`
	Offset   *int    `json:"offset"`
	LangCode *string `json:"langCode"`
	UserID   *int    `json:"userId"`
}

type Notifications struct {
	Notifications []*Notification `json:"notifications"`
	NextOffset    int             `json:"nextOffset"`
}

type RegisterToken struct {
	UserID int    `json:"userId"`
	Token  string `json:"token"`
}

type SimpleNotification struct {
	ID               int    `json:"id"`
	OrderID          int    `json:"orderId"`
	OrderType        string `json:"orderType"`
	Status           bool   `json:"status"`
	OrderDescription string `json:"orderDescription"`
	HyperLink        string `json:"hyperLink"`
	CreatedTime      string `json:"createdTime"`
	UserID           int    `json:"userId"`
	LangCode         string `json:"langCode"`
}

type SimpleNotifications struct {
	Notifications []*SimpleNotification `json:"notifications"`
	NextOffset    int                   `json:"nextOffset"`
}
