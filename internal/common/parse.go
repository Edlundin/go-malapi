package common

import (
	"encoding/json"
	"fmt"

	"github.com/Edlundin/go-malapi/pkg/mal/common"
)

//ParseError parses MAL's JSON error message.
func ParseError(responseBody []byte) (common.Error, error) {
	var malError common.Error

	err := json.Unmarshal(responseBody, &malError)

	if err != nil {
		return malError, fmt.Errorf("demarshalling response body: %s", err.Error())
	}

	if len(malError.Error) == 0 {
		return malError, fmt.Errorf("value for key %q is empty", "error")
	}

	if len(malError.Message) == 0 {
		return malError, fmt.Errorf("value for key %q is empty", "message")
	}

	return malError, nil
}
