# Build a Docker File 

with our Docker file 
```
# syntax=docker/dockerfile:1

FROM eclipse-temurin:17-jdk-jammy

WORKDIR /app

COPY .mvn/ .mvn
COPY mvnw pom.xml ./
RUN ./mvnw dependency:resolve

COPY src ./src

CMD ["./mvnw", "spring-boot:run"]
```

we could run the following command

```
docker build --tag java-docker .
```

then, we could check our image 

    docker images 