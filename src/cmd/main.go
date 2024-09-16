package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/app"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/ports"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/storage"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/storage/postgres"
)

func main() {
	pgString := os.Getenv("POSTGRES_CONN")
	serverAddr := os.Getenv("SERVER_ADDRESS")

	tenderStorage, err := postgres.NewTenderStorage(pgString)
	if err != nil {
		log.Fatal(err)
		return
	}

	employeeStorage, err := postgres.NewEmployeeStorage(pgString)
	if err != nil {
		log.Fatal(err)
		return
	}

	bidStorage, err := postgres.NewBidStorage(pgString)
	if err != nil {
		log.Fatal(err)
		return
	}

	storage := storage.NewStorage(tenderStorage, employeeStorage, bidStorage)

	tenderService := *app.NewTenderService(storage)
	bidService := *app.NewBidService(storage)

	ping := ports.NewPingHandler()
	tenderHandler := ports.NewTenderHandler(tenderService)
	bidHandler := ports.NewBidHandler(bidService)

	r := chi.NewRouter()

	r.Get("/api/ping", ping.Ping)

	r.Get("/api/tenders", tenderHandler.GetTenderByFilter)
	r.Post("/api/tenders/new", tenderHandler.CreateTender)
	r.Get("/api/tenders/my", tenderHandler.GetTendersByUsername)
	r.Get("/api/tenders/{tenderId}/status", tenderHandler.GetTenderStatusById)
	r.Put("/api/tenders/{tenderId}/status", tenderHandler.UpdateTenderStatusById)
	r.Patch("/api/tenders/{tenderId}/edit", tenderHandler.UpdateTenderParams)
	r.Put("/api/tenders/{tenderId}/rollback/{version}", tenderHandler.UpdateTenderVersionRollback)

	r.Post("/api/bids/new", bidHandler.CreateBid)
	r.Get("/api/bids/my", bidHandler.GetBidsByUsername)
	r.Get("/api/bids/{tenderId}/list", bidHandler.GetBidsByTenderId)
	r.Get("/api/bids/{bidId}/status", bidHandler.GetBidsStatusById)
	r.Put("/api/bids/{bidId}/status", bidHandler.UpdateBidStatusById)
	r.Patch("/api/bids/{bidId}/edit", bidHandler.UpdateBidParamsByTenderId)
	r.Put("/api/bids/{bidId}/submit_decision", bidHandler.UpdateBidDecisionByTenderId)
	r.Put("/api/bids/{bidId}/feedback", bidHandler.UpdateBidFeedBackById)
	r.Put("/api/bids/{bidId}/rollback/{version}", bidHandler.UpdateBidVersionRollback)
	r.Get("/api/bids/{tenderId}/reviews", bidHandler.GetBidReviewsById)

	s := &http.Server{
		Addr:           serverAddr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
