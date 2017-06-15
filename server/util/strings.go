package util

// StringSmartDereference dereferences a string pointer; if nil, returns ""
func StringSmartDereference(sp *string) string {
	if sp == nil {
		return ""
	}
	return *sp
}
