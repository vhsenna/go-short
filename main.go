package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OrignalURL   string    `json:"orignal_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OrignalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OrignalURL))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:8]
}

func createURL(orignalURL string) string {
	shortURL := generateShortURL(orignalURL)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OrignalURL:   orignalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func rootPageURL(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/index.html")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	shortURL_ := createURL(data.URL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{
        ShortURL: shortURL_,
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLhandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/r/"):]
	if id == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	url, err := getURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OrignalURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/", rootPageURL)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/r/", redirectURLhandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Starting server on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error running the server:", err)
	}
}
