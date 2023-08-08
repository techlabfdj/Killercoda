# Material description

In your home directory, you can find a folder ```killercode-artifacts``` which is a github repository clone.  
This folder contains a jar file that we will use for our test:   ```java-mem-block-reserver-j17-latest.jar```  

This jar contains  java application that launch a server on http://127.0.0.1:8080  
You could find a description about this application in this [Readme](https://github.com/techlabfdj/killercoda/blob/main/java-mem-block-reserver/README.md)

We also need hte containerized version of this jaav application.  
Let's get it :   

```docker pull ghcr.io/techlabfdj/go-mem-block-reserver:latest```{{exec}}  

Let's now get to the next step !