package utils

import (
	pagination "github.com/gemcook/pagination-go"
	"github.com/go-xorm/xorm"
)

// ApplyOrders はソート条件を設定する
func ApplyOrders(s *xorm.Session, orders []*pagination.Order) {
	for _, order := range orders {
		if order.Direction == pagination.DirectionAsc {
			s.Asc(order.ColumnName)
		} else if order.Direction == pagination.DirectionDesc {
			s.Desc(order.ColumnName)
		}
	}
}
