package dataTransfer

import "github.com/IraIvanishak/test-task-contentquo/breeds"

// ApiResponse represents the structure of needed data of the response from the API
type ApiResponse struct {
	Data []breeds.Breed `json:"data"`
}
