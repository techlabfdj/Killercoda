#!/bin/bash

# check if the gitlab container exists
docker container inspect gitlab.local 2> /dev/null