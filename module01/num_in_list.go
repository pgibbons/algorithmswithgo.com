package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	// empty and or nil condition
	if len(list) == 0 {
		return false
	}
	// loop and find the number
	for _, value := range list {
		if value == num {
			return true
		}
	}
	// return default false
	return false
}
