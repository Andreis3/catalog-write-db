package entities

import (
	"errors"
	"slices"

	err "github.com/andreis3/catalog-write-api/internal/domain/errors"
)

type status uint8

const (
	ACTIVE status = iota
	INACTIVE
)

func (s status) String() string {
	switch s {
	case ACTIVE:
		return "active"
	case INACTIVE:
		return "inactive"
	default:
		return "unknown"
	}
}

var APIKeyStatus = [...]string{ACTIVE.String(), INACTIVE.String()}

type APIKey struct {
	ID     int64
	Name   string
	Status string
	err.EntityErrors
}

func NewAPIKey(name, status string) *APIKey {
	return &APIKey{
		Name:   name,
		Status: status,
	}
}

func (a *APIKey) Validate() *err.EntityErrors {
	if a.Name == "" {
		a.Add(errors.New("name: is required"))
	}
	if a.Status == "" {
		a.Add(errors.New("status: is required"))
	}
	if a.Status != "" && !slices.Contains(APIKeyStatus[:], a.Status) {
		a.Add(errors.New("status: is invalid, valid values are active or inactive"))
	}
	return &a.EntityErrors
}
