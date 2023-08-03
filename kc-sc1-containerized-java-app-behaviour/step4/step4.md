# Run our image 

    docker run -d --publish 8080:8080 java-docker

we could check the process with a curl

```bash
curl --request GET \
--url http://localhost:8080/actuator/health \
--header 'content-type: application/json'
```

# Access the app

[click here]({{TRAFFIC_HOST1_8080}})