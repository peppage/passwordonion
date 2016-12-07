# PasswordOnion [![GoDoc](https://godoc.org/github.com/peppage/passwordonion?status.svg)](https://godoc.org/github.com/peppage/passwordonion)

PasswordOnion is a Go library for encryping and checking passwords.

This way of handling passwords is based off [this blog post](https://blogs.dropbox.com/tech/2016/09/how-dropbox-securely-stores-your-passwords)

## Install
```
    go get github.com/peppage/passwordonion
```

## Documentation

Read [GoDoc](https://godoc.org/github.com/peppage/passwordonion)

## Usage
```Go
// Encrypt password, store p.
p, err := Encrypt(pepper, password)

// Check a user entered password. Verify err is nil.
err := Compare(pepper, string(p), password)
```

License
-----------
[MIT License](LICENSE.md)
