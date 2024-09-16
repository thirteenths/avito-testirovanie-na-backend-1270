package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type GetBidStatusById struct {
	BidId    string `json:"bidId"`
	Username string `json:"username"`
}

func NewGetBidStatusById() *GetBidStatusById {
	return &GetBidStatusById{}
}

func (g *GetBidStatusById) Bind(r *http.Request) error {
	values := r.URL.Query()

	g.BidId = chi.URLParam(r, "bidId")
	if g.BidId == "" {
		return errors.New("missing bidId")
	}

	if values.Has("username") {
		g.Username = values.Get("username")
	} else {
		return errors.New("username is required")
	}

	return nil
}
