package handlers

import (
	"blablagram/logger"
	"blablagram/models"
	"blablagram/service"
	"encoding/json"
	"net/http"
)

type Handlers struct {
	log     logger.Logger
	service *service.Service
}

// New is a constructor for handlers.
func New(log logger.Logger, service *service.Service) *Handlers {
	return &Handlers{
		log:     log,
		service: service,
	}
}

// Status is a handler for status.
func (h *Handlers) Status(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

type ByID struct {
	ID string `json:"id"`
}

// RegisterUser is a handler for user registration, it makes email and password validation.
func (h *Handlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		newErrorResponse(w, h.log, http.StatusBadRequest, "invalid input body", err)
		return
	}

	if err := h.service.UserService.RegisterUser(r.Context(), &user); err != nil {
		newErrorResponse(w, h.log, http.StatusInternalServerError, "failed to register user", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// LoginUser is a handler for user login.
func (h *Handlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		newErrorResponse(w, h.log, http.StatusBadRequest, "invalid input body", err)
		return
	}

	_, err := h.service.UserService.LoginUser(r.Context(), &user)
	if err != nil {
		newErrorResponse(w, h.log, http.StatusInternalServerError, "failed to login user", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CreatePost is a handler for creating a post.
func (h *Handlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		newErrorResponse(w, h.log, http.StatusBadRequest, "invalid input body", err)
		return
	}

	if err := h.service.PostService.CreatePost(r.Context(), &post); err != nil {
		newErrorResponse(w, h.log, http.StatusInternalServerError, "failed to create post", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeletePost is a handler for deleting a post by id.
func (h *Handlers) DeletePost(w http.ResponseWriter, r *http.Request) {
	var req ByID
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		newErrorResponse(w, h.log, http.StatusBadRequest, "invalid input body", err)
		return
	}

	if err := h.service.PostService.DeletePost(r.Context(), req.ID); err != nil {
		newErrorResponse(w, h.log, http.StatusInternalServerError, "failed to delete post", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdatePost is a handler for updating a post by id.
func (h *Handlers) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		newErrorResponse(w, h.log, http.StatusBadRequest, "invalid input body", err)
		return
	}

	if err := h.service.PostService.UpdatePost(r.Context(), &post); err != nil {
		newErrorResponse(w, h.log, http.StatusInternalServerError, "failed to update post", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetPost is a handler for getting a post by id.
func (h *Handlers) GetPost(w http.ResponseWriter, r *http.Request) {
	var req ByID
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		newErrorResponse(w, h.log, http.StatusBadRequest, "invalid input body", err)
		return
	}

	post, err := h.service.PostService.GetPost(r.Context(), req.ID)
	if err != nil {
		newErrorResponse(w, h.log, http.StatusInternalServerError, "failed to get post", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

// GetPosts is a handler for getting all posts.
func (h *Handlers) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.PostService.GetPosts(r.Context())
	if err != nil {
		newErrorResponse(w, h.log, http.StatusInternalServerError, "failed to get posts", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
