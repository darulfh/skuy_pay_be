package model

type Insurance struct {
	UUIDPrimaryKey
	CustomerID     string  `gorm:"type:varchar(100)" json:"customer_id"`
	ProviderName   string  `gorm:"type:varchar(100)" json:"provider_name"`
	Type           string  `gorm:"type:varchar(100)" json:"product_type"`
	Name           string  `gorm:"type:varchar(100)" json:"name"`
	Period         string  `gorm:"type:varchar(100)" json:"period"`
	Class          string  `gorm:"type:varchar(100)" json:"class"`
	NumberOffamily int     `gorm:"type:int" json:"number_of_family"`
	DiscountId     string  `gorm:"type:varchar(100)" json:"discount_id"`
	Price          float64 `gorm:"type:decimal(12)" json:"price"`
}
