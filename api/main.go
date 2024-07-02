package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


func getAlbumsmio(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	prettyJSON, err := json.MarshalIndent(albums, "", "        ")
	if err != nil {
		http.Error(w, "error in the json formatting!", http.StatusInternalServerError) //<----- 500
	} 

 	w.Write(prettyJSON)  /*this will print the string on the server side(obviously)*/
}
	
	

func main() {

	http.HandleFunc("/albums", getAlbumsmio)
	log.Fatal(http.ListenAndServe(":8081", nil))
}