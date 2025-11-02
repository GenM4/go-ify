package server

import (
	"log"
	"net/http"

	"github.com/GenM4/go-ify/internal/api"
	"github.com/GenM4/go-ify/internal/config"
	"github.com/GenM4/go-ify/internal/services"
)

type Server struct {
	http.Server
	Config *config.ServerConfig
}

func NewServer(cfg *config.ServerConfig) *Server {
	srv := &Server{Config: cfg}
	srv.Addr = srv.Config.Address + ":" + srv.Config.Port
	return srv
}

func (srv *Server) Init() error {
	mux := http.NewServeMux()

	// For handling css sheet
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(srv.Config.WebFileRoot+"/static"))))

	mux.HandleFunc("/app/", srv.RedirectToHome)
	mux.HandleFunc("/", srv.MainPageHandler)
	mux.HandleFunc("GET "+srv.Config.ApiPrefix+"/healthz", srv.ReadinessHandler)
	//mux.HandleFunc("GET "+srv.Config.AdminPrefix+"/metrics", srv.MetricsHandlerReqCounter)
	//mux.HandleFunc("POST "+srv.Config.AdminPrefix+"/reset", srv.MetricsHandlerReset)

	r, err := api.NewSpotifyRepository()
	if err != nil {
		return err
	}
	s := services.NewSpotifyRetriever(r)
	h := NewSpotifyHandler(s)
	mux.HandleFunc("/processShareURL/", h.ServeTrackHTTP)

	srv.Handler = mux

	return nil
}

func (srv *Server) Serve() error {
	log.Print("Starting server on " + srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Print("Failed to start server")
		return err
	}

	return nil
}
