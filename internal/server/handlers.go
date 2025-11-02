package server

import (
	"fmt"
	"github.com/GenM4/go-ify/internal/api"
	"html/template"
	"log"
	"net/http"
)

func (srv *Server) ReadinessHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (srv *Server) RequestMade(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.reqCount++
		next.ServeHTTP(w, r)
	})
}

func (srv *Server) MetricsHandlerReset(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	srv.reqCount = 0
}

func (srv *Server) MetricsHandlerReqCounter(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
        <html>

        <body>
            <h1>Welcome, Maxx</h1>
            <p>The dashboard has been visited %d times!</p>
        <body>

        <html>
    `, srv.reqCount)))
}

func (srv *Server) RedirectToHome(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/", http.StatusFound)
}

func (srv *Server) MainPageHandler(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles(srv.webFileRoot + "/templates/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (srv *Server) ProcessShareURL(w http.ResponseWriter, req *http.Request) {
	sAPI := api.SpotifyApiInit()

	url := req.PostFormValue("SpotifyURL")
	r := sAPI.GetArtist(url)

	tmpl := template.Must(template.ParseFiles(srv.webFileRoot + "/templates/index.html"))
	if err := tmpl.ExecuteTemplate(w, "assetListElement", Asset{Title: r.Name}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type Asset struct {
	Title string
}
