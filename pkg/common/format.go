package common

import (
	"encoding/json"
)

func JSON(r interface{}) ([]byte, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return b, nil
}
