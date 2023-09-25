package snipcart

type OrderStatus string
type NotificationType string

const (
	Processed  OrderStatus = "Processed"
	Disputed   OrderStatus = "Disputed"
	Shipped    OrderStatus = "Shipped"
	Delivered  OrderStatus = "Delivered"
	Pending    OrderStatus = "Pending"
	Cancelled  OrderStatus = "Cancelled"
	Dispatched OrderStatus = "Dispatched"

	Comment            NotificationType = "Comment"
	OrderStatusChanged NotificationType = "OrderStatusChanged"
	OrderShipped       NotificationType = "OrderShipped"
	TrackingNumber     NotificationType = "TrackingNumber"
	Invoice            NotificationType = "Invice"
)
