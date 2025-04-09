package post

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
	store     types.PostStore
	userStore types.UserStore
}

func NewHandler(store types.PostStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:     store,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/post", auth.WithJWTAuth(h.handleCreatePost, h.userStore)).Methods("POST")

	router.HandleFunc("/feed", h.handleGetFeed).Methods("GET")
	router.HandleFunc("/post/{id}", h.handleGetPostByID).Methods("GET")
	// TODO router.HandleFunc("/post/{id}", h.handleDeletePostByID).Methods("DELETE")
}

func (h *Handler) handleGetFeed(w http.ResponseWriter, r *http.Request) {
	// var payload types.GetFeedPayload
	// if err := utils.ParseJSON(r, &payload); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// if err := utils.Validate.Struct(payload); err != nil {
	// 	errors := err.(validator.ValidationErrors)
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
	// 	return
	// }

	posts, err := h.store.GetPosts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}

func (h *Handler) handleGetPostByID(w http.ResponseWriter, r *http.Request) {
	// var payload types.GetPostByIDPayload
	// if err := utils.ParseJSON(r, &payload); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// if err := utils.Validate.Struct(payload); err != nil {
	// 	errors := err.(validator.ValidationErrors)
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
	// 	return
	// }

	rawID := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse url argument \"%s\" as integer", rawID))
		return
	}

	post, err := h.store.GetPostByID(int(id))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, post)
}

func (h *Handler) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	var payload types.CreatePostPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreatePost(types.Post{
		UserID: userID,
		Title:  payload.Title,
		Text:   payload.Text,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
