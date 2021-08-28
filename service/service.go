package service

import (
	"net/http"
	"pricingengine/config"
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
		App: &app.App{Config:&config.Config{}},
	}
	// Initializing the config
	err := rpc.App.InitConfig()
	if err != nil {
		panic("Couldn't Initialize App. Please cross verify that the app configs are correct")
	}

	r.Post("/generate_pricing", rpc.GeneratePricing)
	http.ListenAndServe(":3000", r)
}
