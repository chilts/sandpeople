#!/bin/bash
## ----------------------------------------------------------------------------

# / - anyone can hit this
curl                                          localhost:8080/

# /my/ - must be logged in
curl                                          localhost:8080/my/
curl -H 'X-Sandstorm-User-Id: ABCDEF123456'   localhost:8080/my/

# /admin/ - must have admin permission
curl                                          localhost:8080/admin/
curl -H 'X-Sandstorm-Permissions: edit'       localhost:8080/admin/
curl -H 'X-Sandstorm-Permissions: admin'      localhost:8080/admin/
curl -H 'X-Sandstorm-Permissions: edit,admin' localhost:8080/admin/
curl -H 'X-Sandstorm-Permissions: admin,edit' localhost:8080/admin/

# /user/ - dumps the user
curl                                                              \
    localhost:8080/user/

# curl                                                              \
#     -H 'X-Sandstorm-User-Id: ABCDEF123456'                        \
#     localhost:8080/user/

# curl                                                              \
#     -H 'X-Sandstorm-User-Id: ABCDEF123456'                        \
#     -H 'X-Sandstorm-Permissions: edit'                            \
#     localhost:8080/user/

curl                                                              \
    -H 'X-Sandstorm-User-Id: ABCDEF123456'                        \
    -H 'X-Sandstorm-Permissions: admin,edit'                      \
    -H 'X-Sandstorm-Username: Bob Jones'                          \
    -H 'X-Sandstorm-User-Pronouns: it'                            \
    -H 'X-Sandstorm-Preferred-Handle: bobby'                      \
    -H 'X-Sandstorm-User-Picture: https://example.com/avatar.jpg' \
    localhost:8080/user/

## ----------------------------------------------------------------------------
