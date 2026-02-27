package types

import (
	"context"

	"github.com/dibyochakraborty/kitchen/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
