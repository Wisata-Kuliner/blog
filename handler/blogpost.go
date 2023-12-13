package handler

import (
	"blog/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) BlogPostList(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	posts := make([]*model.BlogPost, 0)

	err := h.db.Find(&posts).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = encoder.Encode(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
}

func (h *Handler) BlogPostCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	var payload model.BlogPost
	err := decoder.Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = encoder.Encode(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
}
