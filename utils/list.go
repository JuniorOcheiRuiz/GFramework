package utils

// PHP's list()
func List(arr []string, dest ...*string) {
	for i := range dest {
		if len(arr) > i {
			*dest[i] = arr[i]
		}
	}
}
