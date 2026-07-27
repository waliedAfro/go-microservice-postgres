package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-microservice-postgres/internal/dto"
	"go-microservice-postgres/internal/service"
	"go-microservice-postgres/internal/validation"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

// GET /books
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {

	filter := HandleSearchParam(r)

	books, err := h.service.GetAllBooks(filter)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, books)
}

// GET /books/{id}
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, book)
}

// POST /books
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if errs := validation.Validate(req); len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"errors": errs,
		})
		return
	}

	createdBook, err := h.service.CreateBook(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, createdBook)
}

// DELETE /books/{id}
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteBook(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Helper function
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

func HandleSearchParam(r *http.Request) *dto.BookFilter {

	query := r.URL.Query()

	// Handle pagination with fallback defaults
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	return &dto.BookFilter{
		Page:   page,
		Limit:  limit,
		Search: query.Get("search"),
		Title:  query.Get("title"),
		Author: query.Get("author"),
		Sort:   query.Get("sort"),
	}
}
