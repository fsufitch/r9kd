package db

// NilIfEmpty return a pointer to the string if it has data, or nil if the
// string is empty. Use this to insert NULLs into the database as required.
func NilIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// NilIfZero return a pointer to the int if it is not zero, or nil otherwise.
// Use this to insert NULLs into the database as required.
func NilIfZero(x int) *int {
	if x == 0 {
		return nil
	}
	return &x
}
