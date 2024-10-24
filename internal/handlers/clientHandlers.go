package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"p1/internal/service" // Assuming this is the correct import for ClientService
	"p1/pkg/models"
)

// ClientHandler struct that contains the ClientService
type ClientHandler struct {
	clientService *service.ClientService
}

// NewClientHandler creates a new ClientHandler
func NewClientHandler(clientService *service.ClientService) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
	}
}

// HandleInitClient handles the API request to create a new client
func (h *ClientHandler) HandleInitClient(w http.ResponseWriter, r *http.Request) {
	var clientRequest models.ClientRequest
fmt.Printf("entered  handleInitClient\n")

	// Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&clientRequest); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest) // Return 400 for bad request
		return
	}
	name:= clientRequest.Name
	phone:=clientRequest.Phone
	email:= clientRequest.Email

	// Call the InitClientServices method to create a new client
	if err := h.clientService.InitClientServices(name,email,phone); err != nil {
		http.Error(w, "Failed to create client services: "+err.Error(), http.StatusInternalServerError) // Return 500 for internal server errors
		return
	}

	// Respond with a success message
	response := map[string]string{"message": "Client successfully added to the database!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Return 201 for resource creation
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError) // Handle response encoding failure
	}
}
