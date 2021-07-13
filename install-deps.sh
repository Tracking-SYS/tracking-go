#!/bin/bash
export RDK_PREFIX=/usr/local
sudo wget https://github.com/edenhill/librdkafka/archive/v1.7.0.tar.gz  -O - | sudo tar -xz
cd librdkafka-1.7.0/
./configure --prefix=$RDK_PREFIX
sudo make 
sudo make install
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$RDK_PREFIX/lib