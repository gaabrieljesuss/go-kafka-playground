package validator

func IsPasswordValid(password string) bool {
	return len(password) >= 8 && len(password) <= 50
}
