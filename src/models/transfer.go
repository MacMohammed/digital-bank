package models

import "time"

type Transfer struct {
	ID                   uint64    `json:"id,omitempty"`
	AccountOriginID      uint64    `json:"account_origin_id,string,omitempty"`
	AccountDestinationID uint64    `json:"account_destination_id,string,omitempty"`
	Name                 string    `json:"name,omitempty"`
	Amount               float64   `json:"amount,string"`
	Created_at           time.Time `json:"created_at,omitempty"`
}
