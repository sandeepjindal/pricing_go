package service

import (
	"net/http"
	"time"

	"pricingengine/service/app"
	"pricingengine/service/rpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Start begins a chi-Mux'd net/http server on port 3000
func Start() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(5 * time.Second))

	rpc := rpc.RPC{
		App: &app.App{},
	}

	r.Post("/generate_pricing", rpc.GeneratePricing)
	http.ListenAndServe(":3000", r)
}



//https://dev.to/digi0ps/making-concurrent-api-requests-in-go-4jig
//https://medium.com/dev-bits/making-concurrent-http-requests-in-go-programming-language-823b51bb1dc2
//https://docs.google.com/spreadsheets/d/10z5KCDmTXJQiCGIXDgPH7pmqRffbObUCdORXZLWmFVU/edit#gid=34097807