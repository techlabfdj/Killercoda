
# Installing golang

Commencing the Golang Installation Journey!

Welcome to the initial phase of our Golang adventure! In this section, we  guide you through the process of installing Golang on your system. We'll guide you through each step, from removing old installations to setting up necessary directories and installing your selected Go version. By the end of this segment, you’ll have a functional Golang setup ready for crafting your very first Go program. Let's dive in!

## Remove previous version
Ensure previous installation is removed.

`rm -rf /usr/local/go`{{exec}}
 
## Initialization
First, let's create a go/src and a go/bin directory:   
`mkdir -p $HOME/go/src $HOME/go/bin`{{exec}}

## Installation
Choose the right installation version and corresponding checksum from [golang web site](https://golang.org/dl/).

```
GO_VERSION=1.21.0
GO_CHECKSUM=d0398903a16ba2232b389fb31032ddf57cac34efda306a0eebac34f0965a0742
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
```{{exec}}

Then run the installation

```
wget --quiet https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz -O /tmp/go${GO_VERSION}.linux-amd64.tar.gz\
&& sha256sum /tmp/go${GO_VERSION}.linux-amd64.tar.gz | egrep "^${GO_CHECKSUM}\s+"\
&& echo "Download Ok"\
&& sudo rm -rf /usr/local/go\
&& cd /usr/local && sudo tar zxf /tmp/go${GO_VERSION}.linux-amd64.tar.gz && cd - > /dev/null\
&& cd /usr/local/bin && for file in /usr/local/go/bin/*; do sudo ln -s -f $file; done && cd - > /dev/null\
&& go version >&2\
&& echo "Installation Ok"\
&& rm -f /tmp/go${GO_VERSION}.linux-amd64.tar.gz
```{{exec}}

# Check your installation:

`go version`{{exec}}  
 It should display "go version go1.21.0 linux/amd64"  

`go env`{{exec}}  
It should display a list of environment variables, most of them named like. "GO*"  


Let's now Create your first Go program (click on Next)
