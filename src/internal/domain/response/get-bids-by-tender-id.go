package response

type GetBidsByTenderId struct {
	Bids []Bid `json:"bids"`
}
