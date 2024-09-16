package response

type UpdateTenderStatusById struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ServiceType string `json:"serviceType"`
	Verstion    int    `json:"verstion"`
	CreatedAt   string `json:"createdAt"`
}
