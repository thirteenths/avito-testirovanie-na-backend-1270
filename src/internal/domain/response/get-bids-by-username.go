package response

type GetBidsByUsername struct {
	Bids []Bid `json:"bids"`
}

type Bid struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	AuthorType string `json:"authorType"`
	AuthorId   string `json:"authorId"`
	Version    int    `json:"version"`
	CreatedAt  string `json:"createdAt"`
}
