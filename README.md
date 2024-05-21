# REST API in Go

*Disclaimer*: This is a practice project that follows the tutorial [How To Build a Complete API In Golang (Docker, JWT, MySQL)](https://www.youtube.com/watch?v=2JNUmzuBNV0&list=WL&index=10&t=432s&ab_channel=Tiago)

# Table of Contents
* [STEP 0: Initialize Project](#step-0-initialize-project)

## STEP 0: Initialize Project
1. Choose a directory on your local machine (e.g. `github`). Create a new directory (e.g. `rest-api-go-practice`). Go into this newly created directory.
```
cd /path/to/github/
mkdir rest-api-go-practice
cd rest-api-go-practice
```
2.  Initialize a Go module within this directory.
```
go mod init github.com/username/rest-api-go-practice // go mod init <module-path>
```
NOTE: `github.com/username/` is a part of the module path for Go to manage. It is NOT creating new folders (`github.com`, `username`) on your local machine. Rather, `github.com/username/rest-api-go-practice` is used within the Go ecosystem to help Go uniquely identify this project (aka module). Because domain names (e.g. `github.com`) are unique on the Internet, using this convention ensures global uniqueness of this module path. If you are linking the project to a remote repository hosted on GitHub, it is the common practice to use the full GitHub path `github.com/<username>/<project-name>` as Go module path to be the unique identifier. Two advantages of doing so: 1. using domain name ensures global uniqueness. 2. helps other developers easily find the source repository on GitHub. 
3. Create a new repository on Github. The repository name should be the same as `rest-api-go-practice`. Link the local project to the remote repository by following the instructions provided by GitHub.
    - add and edit `.gitignore`
    - add `README.md`


