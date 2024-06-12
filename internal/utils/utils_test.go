package utils

import "testing"

func TestValidateMethod(t *testing.T) {
	testTable := []struct {
		Expected bool
		Method   string
	}{
		{
			Method:   "PUT",
			Expected: true,
		},
		{
			Method:   "PATCH",
			Expected: true,
		},
		{
			Method:   "TEASDAS",
			Expected: false,
		},
		{
			Method:   "",
			Expected: false,
		},
	}

	for _, tC := range testTable {
		res := ValidateMethod(tC.Method)

		if res != tC.Expected {
			t.Errorf("%s method should be %t. Got: %t", tC.Method, tC.Expected, res)
		}
	}
}
