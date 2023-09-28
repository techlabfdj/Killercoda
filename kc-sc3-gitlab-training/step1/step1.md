# Getting started

your home folder contains a `docker-compose.yaml` file to start up our demo environment easily.
The environment is made of :
- a GitLab CE server
- a GitLab Runner
For more information about the environment, read the comments of `docker-compose.yaml`.

To starts our containers, you must be in the directory containing the **docker-compose.yaml** (in your HOME) and run the following :  
`docker-compose up -d`{{exec}}

