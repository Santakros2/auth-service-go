package security

func PasswordCheck(password string, userPass string) bool {
	if password == userPass {
		return true
	} else {
		return false
	}
}
