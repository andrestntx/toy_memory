package handlers

import "flink/pgk/memory"

type History struct {
	OrderId string `json:"order_id"`
	Locations []memory.Location `json:"history"`
}
