#!/bin/bash

set -x # to test stderr output in /var/log/killercoda

echo starting... # to test stdout output in /var/log/killercoda

git clone https://github.com/techlabfdj/killercoda-artifacts.git

apt install openjdk-17-jre-headless -y
apt install jq -y

touch /tmp/finished