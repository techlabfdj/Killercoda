# Getting started

Your home folder contains a `docker-compose.yaml` file to start up our demo environment easily.
The environment is made of :
- a GitLab CE server
- a GitLab Runner

For more information about the environment, read the comments of the `docker-compose.yaml` file.  

## Set gitlab external url

First we will change the external URL to use for gitlab access.   
Killercoda use some specific URL to access a service with HTTP on a defined port.  
We can generate URLs in bash scripts by using file /etc/killercoda/host, for example :  

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

To starts our containers, you must be in the directory containing the `docker-compose.yaml` (in your HOME) and run the following command:  

`docker-compose up -d`{{exec}}

After this command has finished, you can move on to the next step by clicking the 'Check' button.  

The check button uses a clone of the `script verify_step1.sh` (that you could fin din your HOME) to check that you successfully start our test environment.
