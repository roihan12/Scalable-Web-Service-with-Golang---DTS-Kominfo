package model

type Item struct {
	ItemID      uint64   `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint64   `json:"order_id"`
}

