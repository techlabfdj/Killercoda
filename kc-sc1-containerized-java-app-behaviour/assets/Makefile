.PHONY: all build start stop clean

# Specific variables for JAR execution
JAR_PATH?=$$HOME/killercoda-artifacts/sc1/java-mem-block-reserver-j17-latest.jar
JAVA_CMD?=java

# Specific variables for container execution
CONTAINER_NAME?=java-mem-block-reserver
START_OPTS?=
STOP_OPTS?=
LOGS_OPTS?=
IMAGE_NAME?=java-mem-block-reserver
IMAGE_TAG?=latest
DOCKER_REPOSITORY?=ghcr.io/techlabfdj

#common variables
JAVA_OPTS?=
JAVA_VERSION?=17

#Makefile entries for container
start-c:
	@if docker container inspect $(CONTAINER_NAME) > /dev/null 2>&1; then docker container rm --force --volumes $(CONTAINER_NAME) > /dev/null && printf -- "%s | container was stopped/removed\n" "$$(date --rfc-3339=ns)"; fi
	@docker run -d --name "$(CONTAINER_NAME)" -p 127.0.0.1:8080:8080 $(START_OPTS) -eJAVA_OPTS="$(JAVA_OPTS)" ${DOCKER_REPOSITORY}/$(IMAGE_NAME):$(IMAGE_TAG) > /dev/null && printf -- "%s | container was started\n" "$$(date --rfc-3339=ns)"
	@sleep 2; $(MAKE) status-c

start-c-oomexit-disabled:
	@JAVA_OPTS="$(JAVA_OPTS) -XX:-ExitOnOutOfMemoryError" $(MAKE) start

status-c:
	@inspection=$$(docker container inspect $(CONTAINER_NAME) 2> /dev/null) && docker container ps --all --filter "name=$(CONTAINER_NAME)" && { printf -- "$${inspection}" | jq -e '.[].State.Status == "running"' > /dev/null && printf -- "\n%s | container is running\n" "$$(date --rfc-3339=ns)" || { printf -- "\n%s | container isn't running\n" "$$(date --rfc-3339=ns)"; exit 1; } ; }
	@max=10; sleep=0.5; i=0; while [ $$i -le $$max ]; do health=$$(curl --silent --fail http://127.0.0.1:8080/health) && { printf -- "$${health}" | jq -e '.status == "UP"' > /dev/null; } && printf -- "%s | service is healthy : %s\n" "$$(date --rfc-3339=ns)" "$${health}" && exit 0 || printf -- "%s | service not yet healthy : %s\n" "$$(date --rfc-3339=ns)" "$${health}"; sleep $${sleep}; i=$$(( $$i + 1 )); done; exit 1

metrics-c:
	@docker stats $(CONTAINER_NAME) --no-stream --no-trunc
	@for metric in system.cpu.count; do value=$$(curl --silent --fail "http://127.0.0.1:8080/metrics/$${metric}") && { printf -- "$${value}" | jq -re '(.name | . + (20 - length) * " ") +" | "+(.measurements[] | select( .statistic == "VALUE" ) | .value | tostring | (. + (20 - length)*""))+ " | " +.baseUnit'; } | xargs -I {} printf "%s | %s\n" "$$(date --rfc-3339=ns)" "{}" || break; done
	@for metric in "jvm.memory.used?tag=area:heap" "jvm.memory.max?tag=area:heap"; do value=$$(curl --silent --fail "http://127.0.0.1:8080/metrics/$${metric}") && { printf -- "$${value}" | jq -re '(.name | . + (20 - length) * " ") +" | "+(.measurements[] | select( .statistic == "VALUE" ) | (.value / (1024 * 1024) ) | round | tostring ) + " | " +"M"+.baseUnit'; } | xargs -I {} printf "%s | %s\n" "$$(date --rfc-3339=ns)" "{}" || break; done

exec-c:
	@if inspection=$$(docker container inspect $(CONTAINER_NAME) 2> /dev/null); then docker exec -it $(EXEC_OPTS) $(CONTAINER_NAME) bash; fi

stop-c:
	@if inspection=$$(docker container inspect $(CONTAINER_NAME) 2> /dev/null); then docker container stop $(STOP_OPTS) $(CONTAINER_NAME) > /dev/null && printf -- "%s | container was stopped\n" "$$(date --rfc-3339=ns)" || { printf -- "%s | container failed to stop\n" "$$(date --rfc-3339=ns)"; exit 1; }; else printf -- "%s | nothing to stop\n" "$$(date --rfc-3339=ns)"; fi

logs-c:
	if docker container inspect $(CONTAINER_NAME) > /dev/null 2>&1; then docker logs $(LOGS_OPTS) $(CONTAINER_NAME); fi

clean-c:
	./mvnw clean -DjavaVersion=$(JAVA_VERSION) -DartifactID=$(IMAGE_NAME) -DartifactVersion=$(IMAGE_TAG) -DdockerRepository=${DOCKER_REPOSITORY}
	if docker container inspect $(CONTAINER_NAME) > /dev/null 2>&1; then docker container rm --force --volumes $(CONTAINER_NAME); fi
	for i in $$(docker image ls --quiet --filter "reference=${DOCKER_REPOSITORY}/$(IMAGE_NAME)*"); do docker image rm $$i || exit 1; done

#Makefile entries for jar
start-j:
	@$(JAVA_CMD) -jar $(JAR_PATH) $(JAVA_OPTS) > $$HOME/logfile.log 2>&1 &
		@sleep 10; $(MAKE) status-j

status-j:
	@# Check if the Java application is running using the JAR name.
	@if pgrep -f $(JAR_PATH) > /dev/null ; then \
		echo "The Java application is running."; \
		health=$$(curl --silent --fail http://127.0.0.1:8080/health) && { \
			printf -- "Application health status: %s\n" "$${health}"; \
		} || { \
			echo "Unable to fetch the health status of the application."; \
		}; \
	else \
		echo "The Java application is not running."; \
	fi

metrics-j:
	@# Fetch metrics from the Java application via HTTP endpoints.
	@for metric in system.cpu.count; do \
		value=$$(curl --silent --fail "http://127.0.0.1:8080/metrics/$${metric}") && { \
			printf -- "$${value}" | jq -re '(.name | . + (20 - length) * " ") +" | "+(.measurements[] | select( .statistic == "VALUE" ) | .value | tostring | (. + (20 - length)*""))+ " | " +.baseUnit'; \
		} | xargs -I {} printf "%s | %s\n" "$$(date --rfc-3339=ns)" "{}" || break; \
	done
	@for metric in "jvm.memory.used?tag=area:heap" "jvm.memory.max?tag=area:heap"; do \
		value=$$(curl --silent --fail "http://127.0.0.1:8080/metrics/$${metric}") && { \
			printf -- "$${value}" | jq -re '(.name | . + (20 - length) * " ") +" | "+(.measurements[] | select( .statistic == "VALUE" ) | (.value / (1024 * 1024) ) | round | tostring ) + " | " +"M"+.baseUnit'; \
		} | xargs -I {} printf "%s | %s\n" "$$(date --rfc-3339=ns)" "{}" || break; \
	done

stop-j:
	@# Using `pkill` to kill the Java process based on the JAR path.
	@pkill -f $(JAR_PATH) && echo "The Java application has been stopped."

logs-j:
	@# Display logs for the Java application
	@ cat $$HOME/logfile.log

clean-j:
	@# clean java log file
	@rm -f $$HOME/logfile.log