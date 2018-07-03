package renderer

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pagient/pagient-api/pkg/model"
)

// PatientRequest is the request payload for patient data model
type PatientRequest struct {
	*model.Patient
}

// Bind postprocesses the decoding of the request body
func (pr *PatientRequest) Bind(r *http.Request) error {
	var patient *model.Patient
	if r.Context().Value("patient") != nil {
		patient = r.Context().Value("patient").(*model.Patient)

		if pr.Patient.ID != patient.ID {
			return fmt.Errorf("id attribute is not allowed to be updated")
		}
	}

	if pr.Patient.ClientID != 0 {
		return fmt.Errorf("client_id not allowed")
	}
	return nil
}

// PatientResponse is the response payload for the patient data model
type PatientResponse struct {
	*model.Patient
}

// NewPatientResponse creates a new patient response from patient model
func NewPatientResponse(patient *model.Patient) *PatientResponse {
	resp := &PatientResponse{Patient: patient}

	return resp
}

// Render preprocesses the response before marshalling
func (pr *PatientResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// PatientListResponse is the list response payload for the patient data model
type PatientListResponse []*PatientResponse

// NewPatientListResponse creates a new patient list response from multiple patient models
func NewPatientListResponse(patients []*model.Patient) []render.Renderer {
	list := []render.Renderer{}
	for _, patient := range patients {
		list = append(list, NewPatientResponse(patient))
	}
	return list
}
