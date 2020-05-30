package handler

import (
	"encoding/json"
	"fmt"
	"github.com/emersonlugo/users-api/internal/users"
	"log"
	"net/http"
	"strings"
)

type usersHandler struct {
	userService users.Service
}

func NewHandler(service users.Service) http.Handler {
	usersHandler := &usersHandler{userService: service}
	httpHandler := http.DefaultServeMux
	httpHandler.HandleFunc("/ping", usersHandler.pingServerHTTP)
	httpHandler.HandleFunc("/users/", usersHandler.usersServeHTTP)
	return httpHandler
}

func (handler *usersHandler) pingServerHTTP(response http.ResponseWriter, r *http.Request) {
	responseBody := map[string]string{"message": "pong"}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(responseBody)
}

func (handler *usersHandler) usersServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case http.MethodGet:
		handler.getUsers(response, request)
	case http.MethodPost:
		handler.saveUser(response, request)
	case http.MethodPut:
		handler.updateUser(response, request)
	case http.MethodDelete:
		handler.deleteUser(response, request)
	default:
		http.Error(response, fmt.Sprintf("%s not allowed.", request.Method), http.StatusMethodNotAllowed)
	}
}

func (handler *usersHandler) getUsers(response http.ResponseWriter, request *http.Request) {
	if id := getUserID(request); id != "" {
		handler.getUserByID(response, id)
		return
	}

	usersResponse, err := handler.userService.GetAll()
	if err != nil {
		handlerError(response, err)
		return
	}

	json.NewEncoder(response).Encode(usersResponse)
}

func (handler *usersHandler) getUserByID(response http.ResponseWriter, userID string) {
	user, err := handler.userService.GetByID(userID)
	if err != nil {
		handlerError(response, err)
		return
	}

	json.NewEncoder(response).Encode(user)
}

func (handler *usersHandler) saveUser(response http.ResponseWriter, request *http.Request) {
	userBody, err := decodeUserFromRequestBody(response, request)
	if err != nil {
		return
	}
	user := handler.userService.Save(userBody)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(user)
}

func (handler *usersHandler) updateUser(response http.ResponseWriter, request *http.Request) {
	user, err := decodeUserFromRequestBody(response, request)
	if err != nil {
		return
	}

	id := getUserID(request)

	err = handler.userService.Update(id, user)
	if err != nil {
		handlerError(response, err)
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (handler *usersHandler) deleteUser(response http.ResponseWriter, request *http.Request) {
	id := getUserID(request)
	if id == "" {
		http.Error(response, "Resource Not Found", http.StatusNotFound)
		return
	}
	err := handler.userService.Delete(id)
	if err != nil {
		handlerError(response, err)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}

func getUserID(request *http.Request) string {
	if parameters := strings.Split(request.URL.Path, "/"); len(parameters) > 0 {
		return parameters[2]
	}
	return ""
}

func decodeUserFromRequestBody(response http.ResponseWriter, request *http.Request) (*users.User, error) {
	user := &users.User{}
	if err := json.NewDecoder(request.Body).Decode(user); err != nil {
		log.Print("Error to decode user request body. Reason: " + err.Error())
		http.Error(response, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return user, nil
}

func handlerError(response http.ResponseWriter, err error) {
	if err == users.ErrorUserNotFound || err == users.ErrorUsersNotFound {
		http.Error(response, err.Error(), http.StatusNotFound)
		return
	}
	http.Error(response, "unexpected error", http.StatusInternalServerError)
}
