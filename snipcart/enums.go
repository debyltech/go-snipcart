package snipcart

type OrderStatus string
type NotificationType string

const (
	Processed  OrderStatus = "Processed"
	Disputed               = "Disputed"
	Shipped                = "Shipped"
	Delivered              = "Delivered"
	Pending                = "Pending"
	Cancelled              = "Cancelled"
	Dispatched             = "Dispatched"

	Comment            NotificationType = "Comment"
	OrderStatusChanged                  = "OrderStatusChanged"
	OrderShipped                        = "OrderShipped"
	TrackingNumber                      = "TrackingNumber"
	Invoice                             = "Invice"
)
