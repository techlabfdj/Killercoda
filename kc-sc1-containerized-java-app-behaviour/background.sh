#!/bin/bash

set -x # to test stderr output in /var/log/killercoda

echo starting... # to test stdout output in /var/log/killercoda

#apt install maven -y
#curl -o springbootsample.jar "  https://github-registry-files.githubusercontent.com/655598047/4a33b680-1a46-11ee-8419-16d90300f7ac?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20230711%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230711T090731Z&X-Amz-Expires=300&X-Amz-Signature=ad90208b429d718540a743469f674759485655dfc16b4af44a9cc3fc42fb95aa&X-Amz-SignedHeaders=host&actor_id=0&key_id=0&repo_id=655598047&response-content-disposition=filename%3Dspringbootsample-0.0.1-20230704.083821-1.jar&response-content-type=application%2Foctet-stream"

git clone https://github.com/techlabfdj/killercoda.git

apt install default-jre -y
apt install openjdk-11-jre-headless -y
apt install openjdk-17-jre-headless -y
apt install openjdk-18-jre-headless -y
apt install openjdk-8-jre-headless -y

touch /tmp/finished