package inputcontrol

func IsCommand(msg string) bool {
	if len(msg) == 0 {
		return false
	}
	return msg[0] == '/'
}
