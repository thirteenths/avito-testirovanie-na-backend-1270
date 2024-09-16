package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UpdateBidDecisionFilter struct {
	BidId    string
	Username string
	Decision string
}

func NewUpdateBidDecisionFilter() *UpdateBidDecisionFilter {
	return &UpdateBidDecisionFilter{}
}

func (u *UpdateBidDecisionFilter) Bind(r *http.Request) error {
	values := r.URL.Query()

	u.BidId = chi.URLParam(r, "bidId")
	if u.BidId == "" {
		return errors.New("missing bidId")
	}

	if values.Has("username") {
		u.Username = values.Get("username")
	} else {
		return errors.New("username is required")
	}

	if values.Has("decision") {
		u.Username = values.Get("decision")
	} else {
		return errors.New("decision is required")
	}

	return nil
}
