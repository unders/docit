# Docit
Browse markdown files in directories at a HTTP endpoint.

## Why

## Usage
Command `docit` shows how to use it.

**Examples**:
```
docit serve -index=Readme.md -root=doc -port=8080 -name=Coolio
docit serve -index=landing/index.md -root=home/Projects -port=8080
docit serve -index=landing/index.md -root=test/Projects -help=landing/help.md
```

## TODOs
 - [x] Projects link: Show all projects unders -root when that link is clicked.
 - [x] Import members from file.
 - [x] List Members
 - Add password protection.
 - Add embedded database (Bolt or SQL)
 - Add url list.
 - All Members should have the same password?

## Install

## For developer

Install go dependencies.
```
make deps
```

## Links
* https://github.com/russross/blackfriday
* http://blog.ralch.com/tutorial/golang-tools-comprehension/
* http://graphviz.org/
* https://www.atlassian.com/git/tutorials/git-log/filtering-the-commit-history

## Session
* https://gist.github.com/mschoebel/9398202
* https://golang.org/src/net/http/cookie.go
* http://www.gorillatoolkit.org/pkg/securecookie
* http://0xdabbad00.com/2015/04/23/password_authentication_for_go_web_servers/
* https://jacobmartins.com/2016/04/06/practical-golang-writing-a-simple-login-middleware/
* https://github.com/go-authboss/authboss

## Bolt
* http://npf.io/2014/07/intro-to-boltdb-painless-performant-persistence/
* https://www.progville.com/go/bolt-embedded-db-golang/

## Gobs
* https://blog.golang.org/gobs-of-data
* https://golang.org/pkg/encoding/gob/

## Protocol buffers
* https://developers.google.com/protocol-buffers/docs/gotutorial
