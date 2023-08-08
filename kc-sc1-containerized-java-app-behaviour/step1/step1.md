# Material description

In your home directory, you can find a folder `killercoda-artifacts` which is a github repository clone.    
This folder contains a jar file that we will use for our test: `java-mem-block-reserver-j17-latest.jar`  

This jar contains a java application that launch a server on http://127.0.0.1:8080
You could find a description about this application in this [Readme](https://github.com/techlabfdj/killercoda/blob/main/java-mem-block-reserver/README.md)

We also need the containerized version of this java application.  
Let's get it :  

`docker pull ghcr.io/techlabfdj/go-mem-block-reserver:latest`{{exec}}  

After the image has been successfully pulled, you can proceed to the next step!