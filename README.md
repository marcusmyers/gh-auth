# gh-auth

This is a php port of Chris Hunt's
[github-auth](https://github.com/chrishunt/github-auth).  I did this
mainly to learn `go` programming language.

### Pairing with strangers has never been so good.

**gh-auth** allows you to quickly pair with anyone who has a GitHub account
by adding and removing their public ssh keys from your
[`authorized_keys`](http://en.wikipedia.org/wiki/Ssh-agent) file.

## Install
`$ go install github.com/marcusmyers/gh-auth`

## Ussage

### Adding a user
```bash
$ gh-auth add marcusmyers
Successfully added marcusmyers to your authorized keys file
```

### Removing a user
```bash
$ gh-auth remove marcusmyers
Removed 2 keys from your authorized_keys file
``` 

### Listing users
```bash
$ gh-auth list-users
marcusmyers
```
