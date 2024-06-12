package utils

func ValidateMethod(method string) (ok bool) {
	allowedMethods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

	for _, m := range allowedMethods {
		if m == method {
			return true
		}
	}

	return false
}
