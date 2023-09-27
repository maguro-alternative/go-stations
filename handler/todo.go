package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var todo model.TODO

	if r.Method == http.MethodGet {
		query := r.URL.Query()
		readtodo := model.ReadTODORequest{
			PrevID: 0,
			Size:   5,
		}
		if query.Get("prev_id") != "" {
			prevID, err := strconv.Atoi(query.Get("prev_id"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			readtodo.PrevID = int64(prevID)
		}
		if query.Get("size") != "" {
			size, err := strconv.Atoi(query.Get("size"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			readtodo.Size = int64(size)
		}

		ctx := r.Context()
		res, err := h.svc.ReadTODO(ctx, readtodo.PrevID, readtodo.Size)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(&model.ReadTODOResponse{
			TODOs: res,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if todo.Subject == "" {
			http.Error(w, "subject is empty", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		res, err := h.svc.CreateTODO(ctx, todo.Subject, todo.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(&model.CreateTODOResponse{
			TODO: *res,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == http.MethodPut {
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if todo.ID == 0 || todo.Subject == "" {
			http.Error(w, "subject is empty", http.StatusBadRequest)
			return
		}
		res,err := h.svc.UpdateTODO(r.Context(), todo.ID, todo.Subject, todo.Description)
		if err != nil {
			if (err == &model.ErrNotFound{}) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(&model.UpdateTODOResponse{
			TODO: *res,
		})
	}

}
