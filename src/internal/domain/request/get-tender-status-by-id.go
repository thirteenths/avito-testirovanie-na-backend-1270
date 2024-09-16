package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type GetTenderStatusById struct {
	TenderId string `json:"tenderId"`
	Username string `json:"username"`
}

func NewGetTenderStatusById() *GetTenderStatusById {
	return &GetTenderStatusById{}
}

func (g *GetTenderStatusById) Bind(r *http.Request) error {
	g.TenderId = chi.URLParam(r, "tenderId")
	if g.TenderId == "" {
		return errors.New("missing tenderId")
	}

	if r.URL.Query().Has("username") {
		g.Username = r.URL.Query().Get("username")
	} else {
		return errors.New("missing username")
	}

	return nil
}
