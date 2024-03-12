package model

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"item_id"`
	ItemCode    int    `json:"item_code"`
	Description string `gorm:"not null;type:varchar(190)" json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
