.PHONY: all build start stop clean

CONTAINER_NAME?=java-mem-block-reserver
START_OPTS?=
STOP_OPTS?=
LOGS_OPTS?=
IMAGE_NAME?=java-mem-block-reserver
IMAGE_TAG?=latest
JAVA_OPTS?=
DOCKER_REPOSITORY?=ghcr.io/techlabfdj

all: | build start 
 
build:
	./mvnw package -DartifactID=$(IMAGE_NAME) -DartifactVersion=$(IMAGE_TAG) -DdockerRepository=${DOCKER_REPOSITORY} -Dspring-boot.build-image.publish=false

start:
	@if docker container inspect $(CONTAINER_NAME) > /dev/null 2>&1; then docker container rm --force --volumes $(CONTAINER_NAME) > /dev/null && printf -- "%s | container was stopped/removed\n" "$$(date --rfc-3339=ns)"; fi
	@docker run -d --pull never --name "$(CONTAINER_NAME)" -p 127.0.0.1:8080:8080 $(START_OPTS) -eJAVA_OPTS="$(JAVA_OPTS)" ${DOCKER_REPOSITORY}/$(IMAGE_NAME):$(IMAGE_TAG) > /dev/null && printf -- "%s | container was started\n" "$$(date --rfc-3339=ns)"
	@sleep .2; $(MAKE) status

start-oomexit-disabled:
	@JAVA_OPTS="$(JAVA_OPTS) -XX:-ExitOnOutOfMemoryError" $(MAKE) start

status:
	@inspection=$$(docker container inspect $(CONTAINER_NAME) 2> /dev/null) && docker container ps --all --filter "name=$(CONTAINER_NAME)" && { printf -- "$${inspection}" | jq -e '.[].State.Status == "running"' > /dev/null && printf -- "\n%s | container is running\n" "$$(date --rfc-3339=ns)" || { printf -- "\n%s | container isn't running\n" "$$(date --rfc-3339=ns)"; exit 1; } ; }
	@max=10; sleep=0.5; i=0; while [ $$i -le $$max ]; do health=$$(curl --silent --fail http://127.0.0.1:8080/health) && { printf -- "$${health}" | jq -e '.status == "UP"' > /dev/null; } && printf -- "%s | service is healthy : %s\n" "$$(date --rfc-3339=ns)" "$${health}" && exit 0 || printf -- "%s | service not yet healthy : %s\n" "$$(date --rfc-3339=ns)" "$${health}"; sleep $${sleep}; i=$$(( $$i + 1 )); done; exit 1

# @docker_stats=$(docker stats jdk8 --no-stream --no-trunc --format '{{json .}}'); for metric in "MemUsage" "MemPerc"; do printf -- "$${docker_stats} | xargs -I {} printf "%s | %s\n" "$$(date --rfc-3339=ns)" "{}" || break; done
metrics:
	@docker stats $(CONTAINER_NAME) --no-stream --no-trunc
	@for metric in system.cpu.count; do value=$$(curl --silent --fail "http://127.0.0.1:8080/metrics/$${metric}") && { printf -- "$${value}" | jq -re '(.name | . + (20 - length) * " ") +" | "+(.measurements[] | select( .statistic == "VALUE" ) | .value | tostring | (. + (20 - length)*""))+ " | " +.baseUnit'; } | xargs -I {} printf "%s | %s\n" "$$(date --rfc-3339=ns)" "{}" || break; done
	@for metric in "jvm.memory.used?tag=area:heap" "jvm.memory.max?tag=area:heap"; do value=$$(curl --silent --fail "http://127.0.0.1:8080/metrics/$${metric}") && { printf -- "$${value}" | jq -re '(.name | . + (20 - length) * " ") +" | "+(.measurements[] | select( .statistic == "VALUE" ) | (.value / (1024 * 1024) ) | round | tostring ) + " | " +"M"+.baseUnit'; } | xargs -I {} printf "%s | %s\n" "$$(date --rfc-3339=ns)" "{}" || break; done

exec:
	@if inspection=$$(docker container inspect $(CONTAINER_NAME) 2> /dev/null); then docker exec -it $(EXEC_OPTS) $(CONTAINER_NAME) bash; fi

stop:
	@if inspection=$$(docker container inspect $(CONTAINER_NAME) 2> /dev/null); then docker container stop $(STOP_OPTS) $(CONTAINER_NAME) > /dev/null && printf -- "%s | container was stopped\n" "$$(date --rfc-3339=ns)" || { printf -- "%s | container failed to stop\n" "$$(date --rfc-3339=ns)"; exit 1; }; else printf -- "%s | nothing to stop\n" "$$(date --rfc-3339=ns)"; fi

logs:
	if docker container inspect $(CONTAINER_NAME) > /dev/null 2>&1; then docker logs $(LOGS_OPTS) $(CONTAINER_NAME); fi

clean:
	./mvnw clean -DartifactID=$(IMAGE_NAME) -DartifactVersion=$(IMAGE_TAG) -DdockerRepository=${DOCKER_REPOSITORY}
	if docker container inspect $(CONTAINER_NAME) > /dev/null 2>&1; then docker container rm --force --volumes $(CONTAINER_NAME); fi
	for i in $$(docker image ls --quiet --filter "reference=${DOCKER_REPOSITORY}/$(IMAGE_NAME)*"); do docker image rm $$i || exit 1; done
