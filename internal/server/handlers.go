package server

import (
	"context"
	"net/http"

	"github.com/GenM4/go-ify/internal/services"
	"github.com/GenM4/go-ify/web/templates"
)

type SpotifyHandler struct {
	Service *services.SpotifyRetriever
}

func NewSpotifyHandler(s *services.SpotifyRetriever) *SpotifyHandler {
	return &SpotifyHandler{Service: s}
}

func (srv *Server) ReadinessHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

/* middleware
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
*/

func (srv *Server) RedirectToHome(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/", http.StatusFound)
}

func (srv *Server) MainPageHandler(w http.ResponseWriter, req *http.Request) {
	/*
		tmpl := template.Must(template.ParseFiles(srv.Config.WebFileRoot + "templates/index.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	*/

	t := templates.PublicHome("Hello")
	t.Render(context.Background(), w)
}

func (h *SpotifyHandler) ServeTrackHTTP(w http.ResponseWriter, req *http.Request) {
	url := req.PostFormValue("SpotifyURL")
	track, err := h.Service.GetTrack(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	t := templates.AssetList(track.Album.Images[0].URL, track.Name)
	t.Render(req.Context(), w)
}

func (h *SpotifyHandler) ServePlaylistHTTP(w http.ResponseWriter, req *http.Request) {
	url := req.PostFormValue("SpotifyURL")
	pl, err := h.Service.GetPlaylist(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	t := templates.PlaylistConfirm(pl.Images[0].URL, pl.Name)
	t.Render(req.Context(), w)
}
