package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type GetBidsByTenderIDFilter struct {
	TenderId string
	Username string
	Limit    int
	Offset   int
}

func NewGetBidsByTenderIDFilter() *GetBidsByTenderIDFilter {
	return &GetBidsByTenderIDFilter{}
}

func (g *GetBidsByTenderIDFilter) Bind(r *http.Request) (err error) {
	values := r.URL.Query()

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

	if values.Has("username") {
		g.Username = values.Get("username")
	} else {
		return errors.New("username is required")
	}

	g.TenderId = chi.URLParam(r, "tenderId")
	if g.TenderId == "" {
		return errors.New("missing tenderId")
	}

	return nil
}
