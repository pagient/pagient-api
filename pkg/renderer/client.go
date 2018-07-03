package renderer

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/pagient/pagient-api/pkg/model"
)

// ClientResponse is the response payload for the client data model
type ClientResponse struct {
	*model.Client
}

// NewClientResponse creates a new client response from client model
func NewClientResponse(client *model.Client) *ClientResponse {
	resp := &ClientResponse{Client: client}

	return resp
}

// Render preprocesses the response before marshalling
func (cr *ClientResponse) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// ClientListResponse is the list response payload for the client data model
type ClientListResponse []*ClientResponse

// NewClientListResponse creates a new client list response from multiple client models
func NewClientListResponse(clients []*model.Client) []render.Renderer {
	list := []render.Renderer{}
	for _, client := range clients {
		list = append(list, NewClientResponse(client))
	}
	return list
}