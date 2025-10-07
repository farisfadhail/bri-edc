package handler

import "bri-edc/api/internal/services"

type AuthHandler struct {
	s *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{s: s}
}
