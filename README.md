# gh-auth

This is a `go` port of Chris Hunt's
[github-auth](https://github.com/chrishunt/github-auth).  I did this
mainly to learn the `go` programming language.

### Pairing with strangers has never been so good.

**gh-auth** allows you to quickly pair with anyone who has a GitHub account
by adding and removing their public ssh keys from your
[`authorized_keys`](http://en.wikipedia.org/wiki/Ssh-agent) file.

## Install
`$ go install github.com/marcusmyers/gh-auth`

## Ussage

### Command Line
```
$ gh-auth
NAME:
   gh-auth - allows you to quickly pair with anyone who has a GitHub
account

USAGE:
   gh-auth [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR(S):
   Mark Myers <marcusmyers@gmail>

COMMANDS:
     list     gh-auth list
     add      gh-auth add
     remove   gh-auth remove
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

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
