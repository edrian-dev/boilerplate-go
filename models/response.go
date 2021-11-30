package models

import (
	"encoding/json"
	"errors"
)

// Error ...
type Error struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type Response struct {
	Errors []Error     `json:"errors"`
	Data   interface{} `json:"data"`
}

type STPResponse struct {
	Message string `json:"message"`
}

func (doc *Response) Error() error {
	jsonMarshal, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	return errors.New(string(jsonMarshal))
}
