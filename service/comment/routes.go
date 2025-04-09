package comment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/akamaaru/go-forum/service/auth"
	"github.com/akamaaru/go-forum/types"
	"github.com/akamaaru/go-forum/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store		types.CommentStore
	userStore 	types.UserStore
}

func NewHandler(store types.CommentStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:     store,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/post/{id}/comments", auth.WithJWTAuth(h.handleCreateComment, h.userStore)).Methods("POST")
	router.HandleFunc("/post/{id}/comments", h.handleGetCommentsByPostID).Methods("GET")
}

func (h *Handler) handleCreateComment(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	rawPostID := mux.Vars(r)["id"]
	postID, err := strconv.ParseInt(rawPostID, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse url argument \"%s\" as integer", rawPostID))
		return
	}

	var payload types.CreateCommentPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
 
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err = h.store.CreateComment(types.Comment{
		PostID: int(postID),
		UserID: userID,
		Text:   payload.Text,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleGetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse url argument \"%s\" as integer", rawID))
		return
	}

	comments, err := h.store.GetCommentsByPostID(int(id))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, comments)
}
