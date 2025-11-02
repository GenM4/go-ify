package server

import (
	"log"
	"net/http"

	"github.com/GenM4/go-ify/internal/config"
)

type Server struct {
	http.Server
	reqCount    int
	webFileRoot string
	appPrefix   string
	apiPrefix   string
	adminPrefix string
}

func NewServer(cfg *config.GCFG) *Server {
	server := &Server{
		webFileRoot: cfg.Server.WebRoot,
		appPrefix:   cfg.Server.AppPrefix,
		apiPrefix:   cfg.Server.ApiPrefix,
		adminPrefix: cfg.Server.AdminPrefix,
	}

	server.Addr = cfg.Server.Address + ":" + cfg.Server.Port

	return server
}

func (srv *Server) Init() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(srv.webFileRoot+"/static"))))
	mux.HandleFunc("/app/", srv.RedirectToHome)
	mux.HandleFunc("/", srv.MainPageHandler)
	mux.HandleFunc("/processShareURL/", srv.ProcessShareURL)

	mux.HandleFunc("GET "+srv.apiPrefix+"/healthz", srv.ReadinessHandler)
	mux.HandleFunc("GET "+srv.adminPrefix+"/metrics", srv.MetricsHandlerReqCounter)
	mux.HandleFunc("POST "+srv.adminPrefix+"/reset", srv.MetricsHandlerReset)

	srv.Handler = mux
}

func (srv *Server) Serve() error {
	log.Print("Starting server on " + srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Print("Failed to start server")
		return err
	}

	return nil
}
