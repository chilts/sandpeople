package sandpeople

import "net/http"

// GetUser can be used to obtain the User from the request.
func GetUser(r *http.Request) *User {
	user := r.Context().Value(userIdKey)

	switch v := user.(type) {
	case *User:
		return v
	case nil:
		return nil
	}

	return nil
}
