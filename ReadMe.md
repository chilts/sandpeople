# sandpeople : middleware to check Sandstorm authorisation

## Overview [![GoDoc](https://godoc.org/github.com/chilts/sandpeople?status.svg)](https://godoc.org/github.com/chilts/sandpeople) [![Build Status](https://travis-ci.org/chilts/sandpeople.svg?branch=master)](https://travis-ci.org/chilts/sandpeople) [![Code Climate](https://codeclimate.com/github/chilts/sandpeople/badges/gpa.svg)](https://codeclimate.com/github/chilts/sandpeople)

GoLang middleware to check the relevant `X-Sandstorm-*` headers.

The following headers are used by Sandstorm and collected and provided to you by this middleware:

* `X-Sandstorm-User-Id` - The first 128 bits of a SHA-256. (e.g. '0ba26e59c64ec75dedbc11679f267a40') - **not sent for anonymous users**.
* `X-Sandstorm-Permissions` - Comma separated list of permissions as defined by your app (e.g. 'edit' or 'admin,edit').
* `X-Sandstorm-User-Pronouns` - Usually one of 'neutral', 'he', 'she'. or 'it'. If not specified, will default to 'neutral'.
* `X-Sandstorm-Username` - The full name (e.g. Kurt Friedrich Gödel).
* `X-Sandstorm-Preferred-Handle` - The user's preferred handle (e.g. 'chilts') - **not unique**.
* `X-Sandstorm-User-Picture` - URL of a profile picture (around 128x128).

(Information gleaned from [User Authentication and Permissions](https://docs.sandstorm.io/en/latest/developing/auth/) but you should read that for yourself.)

## Install

```
go get github.com/chilts/sandpeople
```

## Example

```
// middleware to gather up the Sandstorm headers into a sandpeople.User (or nil)
m.Get("/"protected-url", sandpeople.MakeUser("/"))

// check someone is logged in, if not redirect to "/"
m.Get("/"protected-url", sandpeople.RequireUser("/"), homeHandler)

// check someone is logged in, and they have the "admin" permission - if not, redirect to "/"
m.Get("/settings/", sandpeople.RequirePerm("admin", "/"), settingsHandler)

// in your handlers, will return a *sandpeople.User or nil
func handler(w http.ResponseWriter, r *http.Request) {
    user := sandpeople.GetUser(r)
    if user == nil {
        // no-one is logged in
    }

    // print out some info
    fmt.Printf("User = %#s\n", user)
}
```

## Author

By [Andrew Chilton](https://chilts.org), [@andychilton](https://twitter.com/andychilton).

For apps created by [AppsAttic](https://appsattic.com), [@AppsAttic](https://twitter.com/AppsAttic).

## License

[MIT](https://publish.li/mit-qLQqmVTO).
