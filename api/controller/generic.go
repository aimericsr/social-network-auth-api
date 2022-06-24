package controller

import (
	"encoding/json"
	"net/http"

	"github.com/aimericsr/social-network-auth-api/api/custom_error"
)

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	//w.Write([]byte("dfsdf"))
	json.NewEncoder(w).Encode(custom_error.ErrMethodNotAllowed)
	//render.Render(w, r, custom_error.ErrMethodNotAllowed)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	//render.Render(w, r, custom_error.ErrNotFound)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is alive"))
	//fmt.Fprintf(w, "API is alive")
}
