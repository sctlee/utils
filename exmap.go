package utils

func IsExistInMap(m map[string]string, params ...string) bool {
	for _, param := range params {
		if _, ok := m[param]; !ok {
			return false
		}
	}
	return true
}
