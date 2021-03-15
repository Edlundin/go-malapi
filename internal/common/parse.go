package common

import (
	"encoding/json"
	"fmt"

	"github.com/Edlundin/go-malapi/pkg/mal/common"
)

//ParseError parses MAL's JSON "error" object.
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

//ParsePaging parses MAL's JSON "paging" object.
func ParsePaging(responseBody []byte) (common.Paging, error) {
	var paging common.Paging

	err := json.Unmarshal(responseBody, &paging)

	if err != nil {
		return paging, err
	}

	return paging, nil
}

//ParseCalendarDate parses MAL's JSON "end_date" and "start_date" (format YYYY-MM-DD) properties.
func ParseCalendarDate(responseBody []byte) (common.CalendarDate, error) {
	var calDate common.CalendarDate

	err := json.Unmarshal(responseBody, &calDate)

	if err != nil {
		return calDate, err
	}

	return calDate, nil
}
