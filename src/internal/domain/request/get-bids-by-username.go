package request

import (
	"errors"
	"net/http"
	"strconv"
)

type GetBidsByUsernameFilter struct {
	Limit    int
	Offset   int
	Username string
}

func NewGetBidsByUsernameFilter() *GetBidsByUsernameFilter {
	return &GetBidsByUsernameFilter{}
}

func (g *GetBidsByUsernameFilter) Bind(r *http.Request) (err error) {
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

	return nil
}
