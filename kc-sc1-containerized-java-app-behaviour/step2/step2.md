# Initial Test Phase

The first phase of our testing involves launching the Java application directly from the JAR file. To facilitate this, we've provided a Makefile command. 

# Java Application Management with Makefile

### start-j

This command starts the Java application directly from the JAR file. It's the primary way to get the Java application up and running without using a container.

`make start-j`
### status-j

After starting the application with start-j, you can use this command to check the status of the Java application. It verifies if the process is running and also checks the health endpoint at http://127.0.0.1:8080/health.

`make status-j`
### metrics-j

This command fetches and displays metrics from the running Java application. It's useful for monitoring and understanding the performance and behavior of the application.

`make metrics-j`
### stop-j

When you're done testing or if you need to halt the Java application for any reason, use this command. It will find the process associated with the Java application and terminate it.

`make stop-j`
### clean-j

This command is used to clean up any resources or files related to the Java application. It's particularly useful to ensure a fresh environment before starting a new test or after completing tests.

`make clean-j`


