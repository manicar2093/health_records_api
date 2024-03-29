package models

// ErrorReponse is just to encode a json with error attribute.
//
// Example:
// {"error". "This is an error"}
type ErrorReponse struct {
	Error interface{} `json:"error,omitempty"`
}
