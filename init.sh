#!/bin/sh

TMPTMP=${TMPDIR:-/tmp}
echo $TMPTMP

# install c++ library
wget --user-agent=Mozilla -O $TMPTMP/apache-pulsar-client.deb "https://archive.apache.org/dist/pulsar/pulsar-2.4.1/DEB/apache-pulsar-client.deb"
wget --user-agent=Mozilla -O $TMPTMP/apache-pulsar-client-dev.deb "https://archive.apache.org/dist/pulsar/pulsar-2.4.1/DEB/apache-pulsar-client-dev.deb"

sudo apt install -y $TMPTMP/apache-pulsar-client.deb
sudo apt install -y $TMPTMP/apache-pulsar-client-dev.deb
