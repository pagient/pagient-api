package database

import (
	"github.com/pagient/pagient-server/internal/model"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func activeScope(active bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("active = ?", active)
	}
}

func pagerScope(hasPager bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if hasPager {
			return db.Where("pager_id != ?", 0)
		}
		return db.Where("pager_id = ?", 0)
	}
}

// GetPatients lists all patients
func (t *tx) GetPatients() ([]*model.Patient, error) {
	var patients []*model.Patient
	err := t.Find(&patients).Error

	return patients, errors.Wrap(err, "select all patients failed")
}

// GetPatientsWithPagerByStatus returns all patients having pagers and one of the specified states
func (t *tx) GetPatientsWithPagerByStatus(statuses ...model.PatientStatus) ([]*model.Patient, error) {
	var patients []*model.Patient
	stmt := t.Where("pager_id != 0")

	if len(statuses) > 0 {
		stmt.Where("status IN (?)", statuses)
	}

	err := stmt.Find(&patients).Error

	return patients, errors.Wrap(err, "select patients with pagers failed")
}

// GetPatientsByClient returns all patients from client and activity status (first of slice) and assignment of a pager (second of slice)
func (t *tx) GetPatientsByClient(clientID uint, optionals ...bool) ([]*model.Patient, error) {
	var patients []*model.Patient
	stmt := t.Where("client_id = ?", clientID)

	if len(optionals) > 0 {
		stmt = stmt.Scopes(activeScope(optionals[0]))
	}

	if len(optionals) > 1 {
		stmt = stmt.Scopes(pagerScope(optionals[1]))
	}

	err := stmt.Find(&patients).Error

	return patients, errors.Wrap(err, "select patients by client, activity and pager assignment failed")
}

// GetPatient returns a patient by ID
func (t *tx) GetPatient(id uint) (*model.Patient, error) {
	patient := &model.Patient{}
	err := t.Find(patient, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}

	return patient, errors.Wrap(err, "select patient by id failed")
}

// AddPatient stores the values in the repository
func (t *tx) AddPatient(patient *model.Patient) error {
	// FIXME: handle sql constraint errors
	err := t.Create(patient).Error

	return errors.Wrap(err, "create patient failed")
}

// UpdatePatient updates the values in the repository
func (t *tx) UpdatePatient(patient *model.Patient) error {
	// FIXME: handle sql constraint errors
	err := t.Save(patient).Error
	if gorm.IsRecordNotFoundError(err) {
		return &entryNotExistErr{"patient not found"}
	}

	return errors.Wrap(err, "update patient failed")
}

// MarkPatientsInactiveByClient sets active to false for every patient by that client
func (t *tx) MarkPatientsInactiveByClient(clientID uint) error {
	err := t.Where(&model.Patient{
		ClientID: clientID,
		Active:   true,
	}).
		Model(model.Patient{}).
		Updates(map[string]interface{}{"active": false}).
		Error

	return errors.Wrap(err, "upate patients to inactive by client failed")
}

// RemovePatient deletes the values from the repository
func (t *tx) RemovePatient(patient *model.Patient) error {
	err := t.Delete(patient).Error
	if gorm.IsRecordNotFoundError(err) {
		return &entryNotExistErr{"patient not found"}
	}

	return errors.Wrap(err, "delete patient failed")
}

// RemovePatientsByClient removes all patients from client and activity status (first of slice) and assignment of a pager (second of slice)
func (t *tx) RemovePatientsByClient(clientID uint, optionals ...bool) error {
	stmt := t.Where("client_id = ?", clientID)

	if len(optionals) > 0 {
		stmt = stmt.Scopes(activeScope(optionals[0]))
	}

	if len(optionals) > 1 {
		stmt = stmt.Scopes(pagerScope(optionals[1]))
	}

	err := stmt.Delete(model.Patient{}).Error

	return errors.Wrap(err, "delete patients by client, activity and pager assignment failed")
}
