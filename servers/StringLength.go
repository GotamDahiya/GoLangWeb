package servers

// IsEmpty checks whether the string entered by user is empty or not
func IsEmpty(data string) bool {
	check := false
	if len(data) <= 0 {
		check = true
	}
	return check
}