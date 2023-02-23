package snipcart

type OrderStatus string

const (
	Processed  OrderStatus = "Processed"
	Disputed               = "Disputed"
	Shipped                = "Shipped"
	Delivered              = "Delivered"
	Pending                = "Pending"
	Cancelled              = "Cancelled"
	Dispatched             = "Dispatched"
)
