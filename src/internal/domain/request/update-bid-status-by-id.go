package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UpdateBidStatusById struct {
	BidId    string `json:"bidId"`
	Username string `json:"username"`
	Status   string `json:"status"`
}

func NewUpdateBidStatusById() *UpdateBidStatusById {
	return &UpdateBidStatusById{}
}

func (g *UpdateBidStatusById) Bind(r *http.Request) error {
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

	if values.Has("status") {
		g.Status = values.Get("status")
	} else {
		return errors.New("status is required")
	}

	return nil
}
