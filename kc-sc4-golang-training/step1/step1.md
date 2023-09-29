
# Installing golang

Ensure previous installation is removed.

`sudo apt-get remove -y golang-g`{{exec}}

Choose the right installation version and corresponding checksum from [golang web site](https://golang.org/dl/).

```
GO_VERSION=1.21.0
GO_CHECKSUM=d0398903a16ba2232b389fb31032ddf57cac34efda306a0eebac34f0965a0742
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

Now open a terminal/console window and check your go installation `go version && go env`{{exec}}.

You may also install a few other binaries to run `gcc` commands for building Go apps:

```shell
sudo apt install build-essential
```{{exec}}

# Check your installation:

`go version`{{exec}}  
 It should display "go version go1.21.0 linux/amd64"  

`go env`{{exec}}  
It should display a list of environment variables, most of them named like. "GO*"  


Let's now Create your first Go program (click on Next)
