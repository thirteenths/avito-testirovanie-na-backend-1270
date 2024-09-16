package request

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GetBidReviewByTenderId struct {
	TenderID          string
	AuthorUsername    string
	RequesterUsername string
	Limit             int
	Offset            int
}

func NewGetBidReviewByTenderId() *GetBidReviewByTenderId {
	return &GetBidReviewByTenderId{}
}

func (g *GetBidReviewByTenderId) Bind(r *http.Request) (err error) {
	values := r.URL.Query()

	g.TenderID = chi.URLParam(r, "tenderId")
	if g.TenderID == "" {
		return errors.New("missing bidId")
	}

	if values.Has("limit") {
		g.Limit, err = strconv.Atoi(values.Get("limit"))
		if err != nil {
			return err
		}
	} else {
		g.Limit = 5
	}

	if values.Has("offset") {
		g.Offset, err = strconv.Atoi(values.Get("offset"))
		if err != nil {
			return err
		}
	} else {
		g.Offset = 0
	}

	if values.Has("authorUsername") {
		g.AuthorUsername = values.Get("authorUsername")
	} else {
		return errors.New("authorUsername is required")
	}

	if values.Has("requesterUsername") {
		g.RequesterUsername = values.Get("requesterUsername")
	} else {
		return errors.New("requesterUsername is required")
	}

	return nil
}
