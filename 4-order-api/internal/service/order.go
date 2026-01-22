package service

type CreateOrder struct {
	UserId     uint
	ProductIds []uint
}

type CreatedOrder struct {
	Id uint
	CreateOrder
}
