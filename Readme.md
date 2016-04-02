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
 - Add embedded database (Bolt or SQL)
 - Add flag: -member-file for access control (a list of emails.); on start import into database.
 - Import members from file.
 - List Members
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
