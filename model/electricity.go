package model

type Electricity struct {
	UUIDPrimaryKey
	CustomerId      string  `gorm:"type:varchar(100)" json:"customer_id"`
	ProviderName    string  `gorm:"type:varchar(100)" json:"provider_name"`
	Type            string  `gorm:"type:varchar(100)" json:"product_type"`
	Name            string  `gorm:"type:varchar(100)" json:"name"`
	Period          string  `gorm:"type:varchar(100)" json:"period"`
	Token           string  `gorm:"type:varchar(100)" json:"token"`
	ElectricalPower int     `gorm:"type:int" json:"electrical_power"`
	DiscountId      string  `gorm:"type:varchar(100)" json:"discount_id"`
	Price           float64 `gorm:"type:decimal(12)" json:"price"`
}
