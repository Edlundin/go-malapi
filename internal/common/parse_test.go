package common

import (
	"testing"
)

func TestParseError(t *testing.T) {
	if _, err := ParseError([]byte("")); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if _, err := ParseError([]byte(" ")); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if _, err := ParseError([]byte("{}")); err == nil {
		t.Error("failed to pass empty JSON object: no error returned")
	}

	if _, err := ParseError([]byte("{message: error message, error: error type}")); err == nil {
		t.Error("failed to pass malformed JSON object: no error returned")
	}

	if _, err := ParseError([]byte("{\"message\": \"error message\", error: error type}")); err == nil {
		t.Error("failed to pass partially malformed JSON object: no error returned")
	}

	if _, err := ParseError([]byte("{\"message\": \"error message\", error: \"\"}")); err == nil {
		t.Error("failed to pass well formed JSON object with empty value properties: no error returned")
	}

	if malError, err := ParseError([]byte("{\"message\":\"error message\", \"pollution\":\"pollution string\", \"error\":\"error type\",\"pollution\":\"pollution string\"}")); err != nil {
		t.Errorf("failed to pass polluted JSON object: %s", err.Error())
	} else if malError.Message != "error message" {
		t.Errorf("failed to pass polluted formed JSON object: %s: %q != %q", "message", malError.Message, "message error")
	} else if malError.Error != "error type" {
		t.Errorf("failed to pass polluted formed JSON object: %s: %q != %q", "error", malError.Error, "error type")
	}

	if malError, err := ParseError([]byte("{\"pollution\":\"pollution string\", \"message\":\"error message\", \"error\":\"error type\"}")); err != nil {
		t.Errorf("failed to pass polluted JSON object: %s", err.Error())
	} else if malError.Message != "error message" {
		t.Errorf("failed to pass polluted formed JSON object: %s: %q != %q", "message", malError.Message, "message error")
	} else if malError.Error != "error type" {
		t.Errorf("failed to pass polluted formed JSON object: %s: %q != %q", "error", malError.Error, "error type")
	}

	if malError, err := ParseError([]byte("{\"message\":\"error message\", \"error\":\"error type\"}")); err != nil {
		t.Errorf("failed to pass well formed JSON object: %s", err.Error())
	} else if malError.Message != "error message" {
		t.Errorf("failed to pass well formed JSON object: %s: %q != %q", "message", malError.Message, "message error")
	} else if malError.Error != "error type" {
		t.Errorf("failed to pass well formed JSON object: %s: %q != %q", "error", malError.Error, "error type")
	}
}

func TestParsePaging(t *testing.T) {
	if _, err := ParsePaging([]byte("")); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if _, err := ParsePaging([]byte(" ")); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if _, err := ParsePaging([]byte("{}")); err != nil {
		t.Errorf("failed to pass empty JSON object: %s", err.Error())
	}

	if _, err := ParsePaging([]byte("{previous: https://google.com, next: https://myanimelist.net/}")); err == nil {
		t.Error("failed to pass malformed JSON object: no error returned")
	}

	if _, err := ParseError([]byte("{\"previous\": \"https://google.com\", next: https://myanimelist.net/}")); err == nil {
		t.Error("failed to pass partially malformed JSON object: no error returned")
	}

	if _, err := ParseError([]byte("{\"previous\": \"https://google.com\", next: \"\"}")); err == nil {
		t.Error("failed to pass well formed JSON object with empty value properties: no error returned")
	}

	if paging, err := ParsePaging([]byte("{\"previous\":\"https://google.com\", \"pollution\":\"pollution string\", \"next\":\"https://myanimelist.net/\",\"pollution\":\"pollution string\"}")); err != nil {
		t.Errorf("failed to pass polluted JSON object: %s", err.Error())
	} else if paging.PreviousURL.String() != "https://google.com" {
		t.Errorf("failed to pass polluted JSON object: %q: %q != %q", "previous", paging.PreviousURL.String(), "https://google.com")
	} else if paging.NextURL.String() != "https://myanimelist.net/" {
		t.Errorf("failed to pass polluted JSON object: %q: %q != %q", "next", paging.NextURL.String(), "https://myanimelist.net/")
	}

	if paging, err := ParsePaging([]byte("{\"pollution\":\"pollution string\", \"previous\":\"https://google.com\", \"next\":\"https://myanimelist.net/\"}")); err != nil {
		t.Errorf("failed to pass polluted JSON object: %s", err.Error())
	} else if paging.PreviousURL.String() != "https://google.com" {
		t.Errorf("failed to pass polluted JSON object: %q: %q != %q", "previous", paging.PreviousURL.String(), "https://google.com")
	} else if paging.NextURL.String() != "https://myanimelist.net/" {
		t.Errorf("failed to pass polluted JSON object: %q: %q != %q", "next", paging.NextURL.String(), "https://myanimelist.net/")
	}

	if paging, err := ParsePaging([]byte("{\"previous\":\"https://google.com\", \"next\":\"https://myanimelist.net/\"}")); err != nil {
		t.Errorf("failed to pass well JSON object: %s", err.Error())
	} else if paging.PreviousURL.String() != "https://google.com" {
		t.Errorf("failed to pass well JSON object: %q: %q != %q", "previous", paging.PreviousURL.String(), "https://google.com")
	} else if paging.NextURL.String() != "https://myanimelist.net/" {
		t.Errorf("failed to pass well JSON object: %q: %q != %q", "next", paging.NextURL.String(), "https://myanimelist.net/")
	}
}
