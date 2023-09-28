
# Installing golang

* under your `$HOME` (`%USERPROFILE%` for Windows) directory :
  * create a `go` sub-directory
  * mkdir a `src` and `bin` under the `go` sub-directory
      - use this command for Unix like environments `mkdir -p $HOME/go/src $HOME/go/bin`
* Install binary:
  * Download a zip/tar.gz archive from [GOLANG](http://golang.org/dl)
  * archive file should include a sub-directory named go
  * extract your archive under `$HOME/go` (`%USERPROFILE%\go`) or any directory of your choice
    > will install go binaries under `$HOME/go/go`
* update your environment settings with the following changes:
  * define `GOPATH` variable to be `$HOME/go` (`%USERPROFILE%\go`)
  * define `GOROOT` variable to be `$HOME/go/go` (`%USERPROFILE%\go\go`)
  * add `$GOPATH/bin` (`%GOPATH%\bin` for Windows) to your `PATH`
  * add `$GOROOT/bin` (`%GOROOT%\bin` for Windows) to your `PATH`

Now open a terminal/console window and check your go installation `go version && go env`.
