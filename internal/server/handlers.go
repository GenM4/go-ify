package server

import (
	"github.com/GenM4/go-ify/internal/services"
	"html/template"
	"net/http"
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
	tmpl := template.Must(template.ParseFiles(srv.Config.WebFileRoot + "templates/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *SpotifyHandler) ServeTrackHTTP(w http.ResponseWriter, req *http.Request) {
	url := req.PostFormValue("SpotifyURL")
	track, err := h.Service.GetTrack(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	r := track

	tmpl := template.Must(template.ParseFiles("./web/templates/index.html"))
	if err := tmpl.ExecuteTemplate(w, "assetListElement", Asset{Title: r.Name}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type Asset struct {
	Title string
}
