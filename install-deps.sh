#!/bin/bash
sudo apt-key del 41468433
apt-get install -y --no-install-recommends apt-utils wget gnupg software-properties-common
apt-get install -y apt-transport-https ca-certificates
wget -qO - https://packages.confluent.io/deb/6.2/archive.key | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://packages.confluent.io/deb/6.2 stable main"
sudo apt-get update
sudo apt-get install -y librdkafka-dev