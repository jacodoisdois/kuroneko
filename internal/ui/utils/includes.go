package utils

func Includes(arr []string, text string) bool {
	for _, v := range arr {
		if v == text {
			return true
		}
	}
	return false
}
