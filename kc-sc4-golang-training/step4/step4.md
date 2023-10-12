# Compile time

Go is a compiled language, so we can't run this source file directly: we need to **compile** it by creating an executable (binary) file.

For this, we need to run the **go install** command, along with our module name:

```bash
go install training/basics/hello-world
```{{exec}}

![Scan results](./images/go-install.png)

In case of error, the error message is shown in the terminal. But if this command is successful (which should be), it will print nothing.

> And now what ? Where is the binary ?

The Go compiler has kicked in, compiled your source file into a binary file, and stored it inside **$GOPATH/bin** directory:  

`ls $GOPAH/bin`{{exec}}  

![Scan results](./images/ls-gopath-bin.png)

# Run time !

If you have correctly set up your `$PATH` variable during installation steps, you should be able to run your amazing program using `hello-world`{{exec}} command.

If you don't trust me, you can control that the program you ran is the one located under `$GOPATH/bin` directory too:  

`which hello-world`{{exec}}  

![Scan results](./images/hello-world-run.png)

> Amazing ! But now, when I change the **hello.go** file, I have to re-run the freaking long **go install** command every time ?

Yes. And no. I mean... It depends !

If you want to try the code, you can use the `go run`{{exec}} command.

For example, edit the **hello.go** file and change the `"Hello World !"` (English) string to `"Hola Mundo !"` (Spanish).

Then, run `go run hello.go`{{exec}}, and see what happens:

![Scan results](./images/hola-mundo.png)

This will compile and execute the source file immediately. You don't need to compile it first, which is handy during the development process.

> Caution ! It does not means Go lang is both an interpreted language AND a compiled language. It's a compiled language. When executing **go run**, an intermediate binary file is created in a temporary directory, but you don't notice it.

But if you want the binary output creation, the **go install training/basics/hello-world** command from earlier can be shortened.

In our working directory, the following commands are all equivalent:

`go install training/basics/hello-world`{{exec}}  
or  
`go install .`{{exec}}  
or  
`go install`{{exec}}  


For convenience, go commands accept paths relative to the working directory, and default to the package in the current working directory if no other path is given.

Another possibility is running `go build`{{exec}} command. This is similar to `go run`{{exec}} as it will create a binary file, but this file will be placed in the current directory, instead of **$GOPATH/bin** directory for **go install**.

So a typical dev process would be:

1. **go run file.go** to try the source file in a quick way
1. control that the code behaves as expected
1. **go build** to compile the source file locally
1. run that binary file, and check everything is fine
1. **go install** to compile and make it available globally on the system

That's all folks ! 
