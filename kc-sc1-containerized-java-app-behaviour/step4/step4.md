# Step 4 - Second Test Phase

In the next phase of our tests, we'll run the Java application inside a container.
We will use image that we previously download on step 1. 


For ease of execution, we've set up a command in the Makefile.

# Containerized Java Application Management with Makefile

### start-c

This command starts the Java application using container image downloaded. 

`make start-c`
### status-c

After starting the application with start-c, you can use this command to check the status of the Java application. It verifies if the process is running and also checks the health endpoint at http://127.0.0.1:8080/health.

`make status-c`
### metrics-c

This command fetches and displays metrics from the running Java application. It's useful for monitoring and understanding the performance and behavior of the application.

`make metrics-c`
### exec-c

The exec-c command in the Makefile launches a bash session inside the Docker container. It first checks if the container exists using docker container inspect before initiating the bash session with docker exec.

`make exec-c`
### stop-c

When you're done testing or if you need to halt the Java application for any reason, use this command. It will terminate the container.

`make stop-c`
### clean-c

This command is used to clean up any resources or files related to the Java application. It's particularly useful to ensure a fresh environment before starting a new test or after completing tests.

`make clean-c`

# first step: launch the containerized java application 

To initiate the test and proceed to the subsequent phase, start the containerized Java application using the command `make start-c`{{exec}}.


