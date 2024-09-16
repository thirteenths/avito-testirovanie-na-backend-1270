package request

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type GetTendersFilters struct {
	Limit       int
	Offset      int
	ServiceType []string
}

func NewGetTendersByFilters() *GetTendersFilters {
	return &GetTendersFilters{}
}

func (g *GetTendersFilters) Bind(r *http.Request) (err error) {
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

	if values.Has("service_type") {
		g.ServiceType = strings.Split(values.Get("service_type"), ",")
		for _, serviceType := range g.ServiceType {
			if !strings.EqualFold(serviceType, "Construction") &&
				!strings.EqualFold(serviceType, "Delivery") &&
				!strings.EqualFold(serviceType, "Manufacture") {
				return errors.New("invalid service type")
			}
		}
	} else {
		g.ServiceType = []string{"Construction", "Delivery", "Manufacture"}
	}

	return nil
}
