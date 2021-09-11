package models

// FilterRequest represent one of registered filters. FilterName indicates the search
// will be use and values should satisfy search needs
type FilterRequest struct {
	FilterName string      `validate:"required,gt=0" json:"filter_name,omitempty"`
	Values     interface{} `validate:"required" json:"values,omitempty"`
}