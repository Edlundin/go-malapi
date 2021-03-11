package common

import "testing"

func Test_jsonURL_UnmarshalJSON(t *testing.T) {
	jURL := jsonURL{}

	//Testing malformed URL (empty string)
	if err := jURL.UnmarshalJSON([]byte("\"\"")); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	//Testing malformed URL (string composed of one space)
	if err := jURL.UnmarshalJSON([]byte("\" \"")); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	//Testing malformed URL (missing scheme)
	if err := jURL.UnmarshalJSON([]byte("\"myanimelist.net\"")); err == nil {
		t.Errorf("failed to pass malformed URL: no error returned")
	}

	//Testing well formed URL without arguments
	if err := jURL.UnmarshalJSON([]byte("\"https://myanimelist.net\"")); err != nil {
		t.Errorf("failed to pass well formed URL: %q", err.Error())
	} else if jURL.String() != "https://myanimelist.net" {
		t.Errorf("failed to pass well formed URL: URL: %q != %q", jURL.String(), "https://myanimelist.net")
	}

	//Testing well formed URL with one argument
	if err := jURL.UnmarshalJSON([]byte("\"https://myanimelist.net?arg=test\"")); err != nil {
		t.Errorf("failed to pass well formed URL: %q", err.Error())
	} else {
		for k := range jURL.Query() {
			switch k {
			case "arg":
				if v := jURL.Query().Get(k); v != "test" {
					t.Errorf("failed to pass well formed URL: argument %q value %q != %q", k, v, "test")
				}
			default:
				t.Errorf("failed to pass well formed URL: argument %q should not exist", k)
			}
		}
	}

	//Testing well formed URL with two arguments
	if err := jURL.UnmarshalJSON([]byte("\"https://myanimelist.net?arg=test&arg2=0\"")); err != nil {
		t.Errorf("failed to pass well formed URL: %q", err.Error())
	} else {
		for k := range jURL.Query() {
			switch k {
			case "arg":
				if v := jURL.Query().Get(k); v != "test" {
					t.Errorf("failed to pass well formed URL: argument %q value %q != %q", k, v, "test")
				}
			case "arg2":
				if v := jURL.Query().Get(k); v != "0" {
					t.Errorf("failed to pass well formed URL: argument %q value %q != %q", k, v, "test")
				}
			default:
				t.Errorf("failed to pass well formed URL: argument %q should not exist", k)
			}
		}
	}

	//Testing well formed URL with more than two arguments with mixed separator
	if err := jURL.UnmarshalJSON([]byte("\"https://myanimelist.net?arg=test&arg2=0&arg3=test1;test2&arg4=1970-01-01\"")); err != nil {
		t.Errorf("failed to pass well formed URL: %q", err.Error())
	} else {
		for k := range jURL.Query() {
			switch k {
			case "arg":
				if v := jURL.Query().Get(k); v != "test" {
					t.Errorf("failed to pass well formed URL: argument %q value %q != %q", k, v, "test")
				}
			case "arg2":
				if v := jURL.Query().Get(k); v != "0" {
					t.Errorf("failed to pass well formed URL: argument %q value %q != %q", k, v, "test")
				}
			case "arg3":
				if v := jURL.Query().Get(k); v != "test1" {
					t.Errorf("failed to pass well formed URL: argument %q value %q != %q", k, v, "test")
				}
			case "arg4":
				if v := jURL.Query().Get(k); v != "1970-01-01" {
					t.Errorf("failed to pass well formed URL: argument %q value %q != %q", k, v, "test")
				}
			case "test2": //url.URL parser allow semicolon as separator for query arguments
				if v := jURL.Query().Get(k); len(v) > 0 {
					t.Errorf("failed to pass well formed URL: argument %q value %q != %q", k, v, "test")
				}
			default:
				t.Errorf("failed to pass well formed URL: argument %q should not exist", k)
			}
		}
	}
}
