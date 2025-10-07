package handler

import (
	"core-service/internal/validations"
	"core-service/resources"
	"encoding/json"
	"net/http"
	"os"
)

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Service-Token")
	expected := os.Getenv("SERVICE_TOKEN")

	if token == "" || token != expected {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		return
	}

	var req validations.AuthorizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	approved := int64(req.Amount)%2 == 0
	msg := "Transaction declined"
	if approved {
		msg = "Transaction approved"
	}

	resp := resources.ToAuthorizationResponse(approved, msg)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
