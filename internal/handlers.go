package internal

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func ipHandler(w http.ResponseWriter, r *http.Request) {
	remoteSocket := r.RemoteAddr

	// This works for IPv6 and IPv4
	remoteSocketSplit := strings.Split(remoteSocket, ":")
	remoteIP := strings.Join(remoteSocketSplit[:len(remoteSocketSplit)-1], ":")
	// Remove brackets [] in case it is an IPv6 address with brackets
	remoteIP = strings.TrimSuffix(remoteIP, "]")
	remoteIP = strings.TrimPrefix(remoteIP, "[")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(remoteIP))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Success  bool   `json:"success"`
		Health   string `json:"health"`
		Datetime string `json:"datetime"`
	}{
		true,
		"healthy",
		time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	const status = http.StatusNotFound

	resp := struct {
		Success  bool   `json:"success"`
		Error    string `json:"error"`
		Datetime string `json:"datetime"`
	}{
		false,
		"404 - Not Found",
		time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
