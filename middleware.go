package sandpeople

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getHeader(r *http.Request, name string) string {
	headerList := r.Header["X-Sandstorm-"+name]
	if len(headerList) == 0 {
		return ""
	}
	return headerList[0]
}

// RequireUser checks the `X-Sandstorm-User-Id` header and determines if the user is logged in or not. If they are, then the
// next middleware is called. If not, then the request is redirected to the path you specify.
func RequireUser(path string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log.Println("RequireUser(): entry")
			// check to see if there are any of these headers
			v := getHeader(r, "User-Id")
			if v == "" {
				http.Redirect(w, r, path, http.StatusFound)
				return
			}

			if v == "" {
				http.Redirect(w, r, path, http.StatusFound)
			} else {
				next.ServeHTTP(w, r)
			}
		}

		return http.HandlerFunc(fn)
	}
}

// HasPerm checks the `X-Sandstorm-Permissions` header and determines if the user has the permission asked for. If they
// do, then the next middleware is called. If not, then the request is redirected to the path you specify.
func HasPerm(perm string, path string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log.Println("HasPerm(): entry")
			v := getHeader(r, "Permissions")
			fmt.Printf("permissions=%s\n", v)
			if v == "" {
				http.Redirect(w, r, path, http.StatusFound)
				return
			}

			// loop over all permissions
			found := false
			permissions := strings.Split(v, ",")
			for _, permission := range permissions {
				if perm == permission {
					found = true
				}
			}

			if found {
				next.ServeHTTP(w, r)
			} else {
				http.Redirect(w, r, path, http.StatusFound)
			}
		}

		return http.HandlerFunc(fn)
	}
}

// MakeUser gets all of the `X-Sandstorm-*` headers and puts a *User into the request context, and calls the next
// middleware. If the user is not logged in, then the *User will be nil. You can retrieve the user by calling
// GetUser(r).
func MakeUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := getHeader(r, "User-Id")
		if id == "" {
			ctx := context.WithValue(r.Context(), userIdKey, nil)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		user := User{
			ID:          id,
			Name:        getHeader(r, "Username"),
			Permissions: strings.Split(getHeader(r, "Permissions"), ","),
			Pronoun:     getHeader(r, "User-Pronouns"),
			Handle:      getHeader(r, "Preferred-Handle"),
			Avatar:      getHeader(r, "User-Picture"),
		}

		// default the Pronoun
		if user.Pronoun == "" {
			user.Pronoun = "neutral"
		}

		ctx := context.WithValue(r.Context(), userIdKey, &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
