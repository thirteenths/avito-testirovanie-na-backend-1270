package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UpdateBidFeedbackById struct {
	BidId       string
	BidFeedback string
	Username    string
}

func NewUpdateBidFeedbackById() *UpdateBidFeedbackById {
	return &UpdateBidFeedbackById{}
}

func (u *UpdateBidFeedbackById) Bind(r *http.Request) error {
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

	if values.Has("bidFeedback") {
		u.BidFeedback = values.Get("bidFeedback")
	} else {
		return errors.New("bidFeedback is required")
	}

	return nil
}
