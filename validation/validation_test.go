package validation

import (
	"bytes"
	"dix975.com/score/test"
	"errors"
	"net/http"
	"testing"
)

func TestLoadSchema(t *testing.T) {

	test.Setup()
	_, err := LoadSchema("/good.json")

	if err != nil {
		t.Errorf("Expected schema no error [%v]\n", err)
	}
}

func TestLoadMissingSchema(t *testing.T) {

	test.Setup()
	_, err := LoadSchema("/bad.json")

	if err == nil {
		t.Errorf("Expected an error")
	}
}

func TestError(t *testing.T) {

	issues := []Issues{
		Issues{
			Field:       "field.1",
			Value:       "value.1",
			Description: "description.1",
			Type:        "type.1",
		},
		Issues{
			Field:       "field.2",
			Value:       "value.2",
			Description: "description.2",
			Type:        "type.2",
		},
	}

	err := Error{Issues: issues}

	message := err.Error()

	test.AssertEqual(t, message, "- description.1\n- description.2\n", "")
}

func TestValidation(t *testing.T) {

	cases := []map[string]interface{}{
		{
			"json":          "{}",
			"schema":        "good.json",
			"expectedError": true,
			"result":        "- name is required\n",
		},
		{
			"json":          "{\"name\":\"name.1\"}",
			"schema":        "good.json",
			"expectedError": false,
			"result":        "",
		},
		{
			"json":          "{}",
			"schema":        "missing.json",
			"expectedError": true,
			"result":        "no such file or directory",
		},
		{
			"json":          "{",
			"schema":        "good.json",
			"expectedError": true,
			"result":        "unexpected EOF",
		},
		{
			"json":          "{}",
			"schema":        "bad.json",
			"expectedError": true,
			"result":        "invalid character '\\n' in string literal",
		},
	}

	for i, c := range cases {

		err := Validate(c["json"].(string), c["schema"].(string))

		if err != nil {
			test.AssertHasSuffix(t, err.Error(), c["result"].(string))
		} else {

			if c["expectedError"].(bool) {

				t.Error("expected error case : ", i)
			}

		}
	}
}

func TestFromPost(t *testing.T) {

	cases := []map[string]interface{}{
		{
			"json":          "{}",
			"schema":        "good.json",
			"expectedError": true,
			"result":        "- name is required\n",
		},
		{
			"json":          "{\"name\":\"name.1\"}",
			"schema":        "good.json",
			"expectedError": false,
			"result":        "",
		},
		{
			"json":          "{}",
			"schema":        "missing.json",
			"expectedError": true,
			"result":        "no such file or directory",
		},
		{
			"json":          "{",
			"schema":        "good.json",
			"expectedError": true,
			"result":        "unexpected EOF",
		},
		{
			"json":          "{}",
			"schema":        "bad.json",
			"expectedError": true,
			"result":        "invalid character '\\n' in string literal",
		},
	}

	for i, c := range cases {

		body := test.NopCloser{
			Reader: bytes.NewBufferString(c["json"].(string)),
		}

		request := http.Request{
			Body: body,
		}

		var model map[string]interface{}

		err := FromPost(&request, c["schema"].(string), &model)

		if err != nil {
			test.AssertHasSuffix(t, err.Error(), c["result"].(string))
		} else {

			if c["expectedError"].(bool) {

				t.Error("expected error case : ", i)
			}
		}
	}
}


func TestFromPostBadReader(t *testing.T) {

	badReader := test.BadReader{ Err: errors.New("error.1")}
	body := test.NopCloser{
		Reader: &badReader,
	}

	request := http.Request{
		Body: body,
	}

	var model map[string]interface{}

	err := FromPost(&request, "good.json", &model)
	test.AssertEqual(t, err.Error(), "error.1", "")

}