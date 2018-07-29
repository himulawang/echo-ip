package internal

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

const service = "echo-ip"
const srcUrl = "https://github.com/greenstatic/echo-ip"

func ipHandler(w http.ResponseWriter, r *http.Request) {

	remoteSocket := r.RemoteAddr

	// This works for IPv6 and IPv4
	remoteSocketSplit := strings.Split(remoteSocket, ":")
	remoteIP := strings.Join(remoteSocketSplit[:len(remoteSocketSplit)-1], ":")

	ip := remoteIP

	forwardedForIP := r.Header.Get("X-Forwarded-For")

	if forwardedForIP != "" {
		ip = forwardedForIP
	}

	// Response
	type IPDetails struct {
		RemoteIP      string `json:"remoteIP"`
		XForwardedFor string `json:"forwardedForIP"`
	}

	resp := struct {
		Success   bool      `json:"success"`
		IP        string    `json:"ip"`
		Datetime  string    `json:"datetime"`
		IPDetails IPDetails `json:"ipDetails"`
		Service   string    `json:"service"`
		Version   string    `json:"version"`
		SrcUrl    string    `json:"srcUrl"`
	}{
		true,
		ip,
		time.Now().Format(time.RFC3339),
		IPDetails{
			remoteIP,
			forwardedForIP,
		},
		service,
		version,
		srcUrl,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
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
