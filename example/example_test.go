package example

import (
	"testing"

	"github.com/infobloxopen/atlas-app-toolkit/query"
)

func TestFilteringPermissionsValidation(t *testing.T) {
	data := map[string]string{
		"first_name==\"Jan\"":      "Operation EQ is not allowed for 'first_name'",
		"User.first_name!=\"Sam\"": "Operation EQ is not allowed for 'User.first_name'",
		"first_name~\"Sam.*\"":     "",

		"middle_name==\"Jan\"":       "",
		"middle_name!=\"Sam\"":       "",
		"User.middle_name~\"Sam.*\"": "",

		"last_name~\"Jan\"":         "Operation MATCH is not allowed for 'last_name'",
		"User.last_name!~\"Sam\"":   "Operation MATCH is not allowed for 'User.last_name'",
		"last_name==\"Sam.*\"":      "",
		"User.last_name!=\"Sam.*\"": "",

		"age==18":     "Operation EQ is not allowed for 'age'",
		"age>=18":     "Operation GE is not allowed for 'age'",
		"User.age<18": "Operation LT is not allowed for 'User.age'",

		"height>=180":     "",
		"height>180":      "Operation GT is not allowed for 'height'",
		"User.height<180": "Operation LT is not allowed for 'User.height'",
	}

	for param, expected := range data {
		response := ""
		f, err := query.ParseFiltering(param)
		if err != nil {
			t.Fatalf("Invalid filtering data '%s'", param)
		}
		resp := Validate(f, nil, "/example.TestService/List")
		if resp != nil {
			response = resp.Error()
		}
		if expected != response {
			t.Errorf("Error, for filtering data '%s' expected '%s', got '%s'", param, expected, response)
		}

	}
}

func TestFilteringPermissionsValidationForRead(t *testing.T) {
	data := []string{
		"first_name==\"Jan\"",
		"last_name~\"Jan\"",
		"age==18",
		"height>=180",
		"height<180",
	}

	for _, param := range data {
		f, err := query.ParseFiltering(param)
		if err != nil {
			t.Fatalf("Invalid filtering data '%s'", param)
		}
		resp := Validate(f, nil, "/example.TestService/Read")
		if resp != nil {
			t.Errorf("Error, for filtering data '%s' got '%s'", param, resp.Error())
		}

	}
}

func TestSortingPermissionsValidation(t *testing.T) {
	data := map[string]string{
		"first_name, height, middle_name":                "pagination is not allowed for 'middle_name'",
		"User.first_name, User.height, User.middle_name": "pagination is not allowed for 'User.middle_name'",
	}

	for param, expected := range data {
		response := ""
		p, err := query.ParseSorting(param)
		if err != nil {
			t.Fatalf("Invalid paging data '%s'", param)
		}
		resp := Validate(nil, p, "/example.TestService/List")
		if resp != nil {
			response = resp.Error()
		}
		if expected != response {
			t.Errorf("Error, for paging data '%s' expected '%s', got '%s'", param, expected, response)
		}
	}
}
