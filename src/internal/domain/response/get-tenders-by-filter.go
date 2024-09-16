package response

type GetTendersByFilter struct {
	Tenders []Tender `json:"tenders"`
}

type Tender struct {
	Id          string `json:"id"`          // "id": "550e8400-e29b-41d4-a716-446655440000",
	Name        string `json:"name"`        // "name": "Доставка товары Казань - Москва",
	Description string `json:"description"` // "description": "Нужно доставить оборудовоние для олимпиады по робототехники",
	Status      string `json:"status"`      // "status": "Created",
	ServiceType string `json:"serviceType"` // "serviceType": "Delivery",
	Verstion    int    `json:"verstion"`    // "verstion": 1,
	CreatedAt   string `json:"createdAt"`   // "createdAt": "2006-01-02T15:04:05Z07:00"
}
