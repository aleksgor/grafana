package models

import (
	"errors"
)

// Typed errors
var (
	ErrPreferencesNotFound = errors.New("Preferences not found")
)

type Preferences struct {
	Id       int64
	PrefId   int64
	PrefType string
	PrefData map[string]interface{}
}

// ---------------------
// QUERIES

type GetPreferencesQuery struct {
	PrefId   int64
	PrefType string

	Result *Preferences
}

// ---------------------
// COMMANDS

type SavePreferencesCommand struct {
	PrefData map[string]interface{} `json:"prefData" binding:"Required"`
	PrefId   int64                  `json:"-"`
	PrefType string                 `json:"-"`
}

// ----------------------
// DTO & Projections

type PreferencesDTO struct {
	PrefId   int64                  `json:"prefId"`
	PrefType string                 `json:"prefType"`
	PrefData map[string]interface{} `json:"prefData"`
}