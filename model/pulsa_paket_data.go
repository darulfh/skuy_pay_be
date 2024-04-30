package model

const PULSA_TYPE = "pulsa"
const PAKET_DATA_TYPE = "data"

type PulsaPaketData struct {
	UUIDPrimaryKey
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Code        string  `json:"code"`
	Provider    string  `json:"provider"`
	Price       float64 `json:"price"`
	IsActive    *bool   `json:"is_active" gorm:"default:true"`
	Description string  `json:"description"`
}

type PPDResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Code        string  `json:"code"`
	Provider    string  `json:"provider"`
	Price       float64 `json:"price"`
	IsActive    *bool   `json:"is_active"`
	Description string  `json:"description"`
}
