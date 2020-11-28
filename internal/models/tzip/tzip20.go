package tzip

import (
	"encoding/json"

	"github.com/baking-bad/bcdhub/internal/helpers"
)

// TZIP20 -
type TZIP20 struct {
	Events []Event `json:"events,omitempty"`
}

// Event -
type Event struct {
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	Pure            string                `json:"pure"`
	Implementations []EventImplementation `json:"implementations"`
}

// EventImplementation -
type EventImplementation struct {
	MichelsonParameterEvent       MichelsonParameterEvent       `json:"michelson-parameter-event"`
	MichelsonInitialStorageEvent  MichelsonInitialStorageEvent  `json:"michelson-initial-storage-event"`
	MichelsonExtendedStorageEvent MichelsonExtendedStorageEvent `json:"michelson-extended-storage-event"`
}

// MichelsonParameterEvent -
type MichelsonParameterEvent struct {
	Sections
	Entrypoints []string `json:"entrypoints"`
}

// InEntrypoints -
func (event MichelsonParameterEvent) InEntrypoints(entrypoint string) bool {
	return helpers.StringInArray(entrypoint, event.Entrypoints)
}

// Is -
func (event MichelsonParameterEvent) Is(entrypoint string) bool {
	return !event.Empty() && event.InEntrypoints(entrypoint)
}

// Sections -
type Sections struct {
	Parameter  json.RawMessage `json:"parameter"`
	ReturnType json.RawMessage `json:"return-type"`
	Code       json.RawMessage `json:"code"`
}

// Empty -
func (s Sections) Empty() bool {
	null := "null"
	return string(s.Code) == null && string(s.Parameter) == null && string(s.ReturnType) == null
}

// MichelsonInitialStorageEvent -
type MichelsonInitialStorageEvent struct {
	Sections
}

// MichelsonExtendedStorageEvent -
type MichelsonExtendedStorageEvent struct {
	Sections
	Entrypoints []string `json:"entrypoints"`
}

// InEntrypoints -
func (event MichelsonExtendedStorageEvent) InEntrypoints(entrypoint string) bool {
	return helpers.StringInArray(entrypoint, event.Entrypoints)
}

// Is -
func (event MichelsonExtendedStorageEvent) Is(entrypoint string) bool {
	return !event.Empty() && event.InEntrypoints(entrypoint)
}