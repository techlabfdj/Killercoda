 
# Introduction

The focus of this scenario is to test the behavior of a Java application on the same host but under different deployment strategies. 

Firstly, we will deploy Java directly on the host, specifying certain parameters to observe the application's performance and behavior. 
Secondly, we will deploy the application within a container, leveraging the isolation that containerization provides. During this phase, we will also manipulate certain container parameters to understand their impact on the application. 

This approach allows us to compare and contrast the behavior of the application in different deployment contexts. The aim is to test and understand java application behaviour reaching memory limit.

Additionnaly, we will also discuss about cpu limit during these different deployment contexts.

# Material description

On this host, you can find a jar file, a github repositor clone: killercode-artifacts.  
This repository contains a jar file that we will use for test:   java-mem-block-reserver-j17-latest.jar  

This jar contains  java application that launch a server on http://127.0.0.1:8080  
You could find a description about this application in this [Readme](https://github.com/techlabfdj/killercoda/blob/main/java-mem-block-reserver/README.md)

Let's now get to the next step !