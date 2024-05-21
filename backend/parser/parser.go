package parser

import (
	"encoding/json"
	"io"
)

func ParseBodyParams(body io.ReadCloser, params interface{}) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &params)
}