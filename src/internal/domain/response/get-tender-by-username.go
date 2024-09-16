package response

type GetTenderByUsername struct {
	Tenders []Tender `json:"tenders"`
}
