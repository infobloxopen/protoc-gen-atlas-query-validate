package example

import (
	"testing"

	"github.com/infobloxopen/atlas-app-toolkit/query"
	"github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
)

func TestFilteringPermissionsValidation(t *testing.T) {
	tests := []struct {
		Query string
		Err   bool
	}{
		{`first_name~"Sam"`, false},
		{`first_name=="Sam"`, false},
		{`weight==1`, false},
		{`comment=="comment1"`, false},
		{`speciality~"spec"`, false},
		{`last_name=="Smith"`, false},
		{`last_name~"Smith"`, false},
		{`id=="some_id"`, true},
		{`first_name<"Sam"`, true},
		{`weight<=1`, true},
		{`first_name<"Jan"`, true},
		{`speciality=="spec"`, true},
		{`unknown_field=="unk"`, true},
		{`id=="some_id"`, true},
		{`array=="array"`, true},
		{`custom_type=="custom_type"`, true},
		{`custom_type_string~"custom_type"`, false},
		{`home_address.city=="city"`, false},
		{`home_address.city~"city"`, true},
		{`home_address.country~"country"`, false},
		{`work_address.city=="city"`, true},
	}

	for _, test := range tests {
		f, err := query.ParseFiltering(test.Query)
		if err != nil {
			t.Fatalf("Invalid filtering data '%s'", test.Query)
		}
		err = options.ValidateFiltering(f, ExampleMessagesRequireQueryValidation["User"])
		if err != nil {
			if test.Err == false {
				t.Errorf("Unexpected error for %s query: %s", test.Query, err)
			}
		} else {
			if test.Err == true {
				t.Errorf("Expected error for %s query, but got no error", test.Query)
			}
		}
	}
}

func TestSortingPermissionsValidation(t *testing.T) {
	tests := []struct {
		Query string
		Err   bool
	}{
		{`first_name, weight`, false},
		{`comment`, false},
		{`on_vacation`, true},
		{`speciality`, true},
		{`unknown_field`, true},
		{`array`, true},
		{`custom_type`, true},
		{`custom_type_string`, false},
		{`home_address.city`, true},
		{`home_address.country`, false},
		{`work_address.city`, true},
	}

	for _, test := range tests {
		s, err := query.ParseSorting(test.Query)
		if err != nil {
			t.Fatalf("Invalid sorting data '%s'", test.Query)
		}
		err = options.ValidateSorting(s, ExampleMessagesRequireQueryValidation["User"])
		if err != nil {
			if test.Err == false {
				t.Errorf("Unexpected error for %s query: %s", test.Query, err)
			}
		} else {
			if test.Err == true {
				t.Errorf("Expected error for %s query, but got no error", test.Query)
			}
		}
	}
}
