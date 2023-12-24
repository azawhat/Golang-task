package main

import (
	"encoding/json"
	"net/http"
)

const port = ":8080"

type reqData struct {
	Message string `json:"message"`
}
type ResponseData struct{ Status, Message string }

func main() { http.HandleFunc("/process", handleRequest); http.ListenAndServe(port, nil) }

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req reqData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseData{"success", "Data successfully received"})
}
