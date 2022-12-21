package utils

func MustString(s string) string {
	if len(s) == 0 {
		panic("empty string")
	}
	return s
}
