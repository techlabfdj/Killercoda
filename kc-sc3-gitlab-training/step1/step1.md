# Getting started

Your home folder contains a `docker-compose.yaml` file to start up our demo environment easily.
The environment is made of :
- a GitLab CE server
- a GitLab Runner

For more information about the environment, read the comments of the `docker-compose.yaml` file.  

## Set gitlab external url

First, we need to modify the external URL for accessing GitLab.  
Killercoda utilizes specific URLs to access a service via HTTP on a designated port.  
URLs can be generated within bash scripts using the file located at /etc/killercoda/host, as shown in the following example:

`sed 's/PORT/80/g' /etc/killercoda/host`{{exec}} will generate a link for URL on port 80 on the host/VM where this is executed.  

Ok so let's replace gitlab external url with the killercoda specific one in the docker compose file.  
In `docker-compose.yaml` we define '#EXTERNAL_URL' as a placeholder for the killercode URL.  
`grep EXTERNAL_URL ./docker-compose.yaml`{{exec}}  

Let's update it :  

``` 
EXTERNAL_URL=$(sed 's/PORT/80/g' /etc/killercoda/host)
sed "s|#EXTERNAL_URL|$EXTERNAL_URL|" -i ./docker-compose.yaml
```{{exec}}

You could check with a grep the the variable has been updated:  
`grep external ./docker-compose.yaml`{{exec}}.  

## start the gitlab container

To start our containers, we must be in the directory containing the `docker-compose.yaml` (in your HOME) and run the following command:  

`docker-compose up -d`{{exec}}

After the command completes, proceed by clicking the 'Check' button.  
This button triggers a clone of the script verify_step1.sh (located in your HOME directory) to verify the successful initialization of our test environment.