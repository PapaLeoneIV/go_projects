package main

import (
	"net/http"
	"regexp"
	"encoding/json"
	"sync"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
type albumHandler struct {
	store *datastore
}

type datastore struct {
	m map[string]album
	*sync.RWMutex
}
	

var (
    listAlbumsre = regexp.MustCompile(`^\/album[\/]*$`)
    getAlbumre  = regexp.MustCompile(`^\/album\/(\d+)$`)
    createAlbumre = regexp.MustCompile(`^\/album[\/]*$`)
)

func internalServerError(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("not found"))
}


func (h *albumHandler) List(w http.ResponseWriter, req *http.Request){

}
func (h *albumHandler) Get(w http.ResponseWriter, req *http.Request){
	matches := getAlbumre.FindStringSubmatch(req.URL.Path) 
	if len(matches) < 2{
		notFound(w, req)
		return
	}
	h.store.RLock()
	album, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("user not found"))
        return
    }
	prettyJSON, err := json.Marshal(album)
	if err != nil {
		internalServerError(w, req)
		return
	}
	w.WriteHeader(http.StatusOK)
    w.Write(prettyJSON)
}

func (h *albumHandler) Add(w http.ResponseWriter, req *http.Request){

}


func (h *albumHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch {
	case req.Method == "GET" && getAlbumre.MatchString(req.URL.Path):
		h.Get(w, req)
	case req.Method == "GET" && listAlbumsre.MatchString(req.URL.Path):
		h.List(w, req)
	case req.Method == "POST" && createAlbumre.MatchString(req.URL.Path):
		h.Add(w, req)
	default:
		notFound(w, req)
        return
	}
}

func main() {

	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	
	store := &datastore{
		m: make(map[string]album),
	}
	
	for _, album := range albums {
		store.m[album.ID] = album
	}
	
    mux := http.NewServeMux()
    mux.Handle("/album/", &albumHandler{ store:store })
    http.ListenAndServe(":8080", mux)
}
