#!/bin/bash

sudo wget https://github.com/edenhill/librdkafka/archive/v1.7.0.tar.gz  -O - | sudo tar -xz
cd librdkafka-1.7.0/
sudo ./configure --install-deps
sudo make
sudo make install