package service

import (
	"encoding/json"
	"go-kafka/entity"
	"net/http"
)

func (h *Handler) Inquiry(w http.ResponseWriter, r *http.Request) {
	rs := h.db.Find(&[]entity.UserInfo{})
	if rs.Error != nil {
		return
	}
	u := rs.Value.(*[]entity.UserInfo)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&u)
}
