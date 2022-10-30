package albumRestApi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

func RunAlbumStandardApi() {
	mux := http.NewServeMux()
	mux.HandleFunc("/albums/", albumsHandler)

	http.ListenAndServe(serverHostname, mux)
}

var albumsRouteRegex = regexp.MustCompile(`^\/albums[\/]*$`)
var albumByIdRouteRegex = regexp.MustCompile(`^\/albums\/(\d+)$`)

func albumsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("New request: [%s] %s\n", r.Method, r.URL)

	w.Header().Set("content-type", "application/json")

	switch r.Method {
	case http.MethodGet:
		if albumsRouteRegex.MatchString(r.URL.Path) {
			getAlbums_Standard(w, r)
			return
		}

		if albumByIdRouteRegex.MatchString(r.URL.Path) {
			getAlbumByID_Standard(w, r)
			return
		}

		return

	case http.MethodPost:
		if albumsRouteRegex.MatchString(r.URL.Path) {
			postAlbums_Standard(w, r)
			return
		}

		return
	}

	notFound(w, r)
}

func getAlbums_Standard(w http.ResponseWriter, r *http.Request) {
	writeSerialized(albums, w, r)
}

func getAlbumByID_Standard(w http.ResponseWriter, r *http.Request) {
	matches := albumByIdRouteRegex.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}

	album, ok := albums.GetAlbumByID(matches[1])
	if !ok {
		notFound(w, r)
		return
	}

	writeSerialized(album, w, r)
}

func postAlbums_Standard(w http.ResponseWriter, r *http.Request) {
	var album Album

	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		internalServerError(w, r)
		return
	}

	albums.AddAlbum(album)

	writeSerialized(album, w, r)
}

func writeSerialized(v any, w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("album not found"))
}
