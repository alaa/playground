FROM ruby:2.2.3-slim
USER root

# Dependencies
RUN apt-get update
RUN apt-get install -y uuid uuid-dev uuid-runtime
RUN apt-get install -y libzmq-dbg libzmq-dev libzmq1
RUN apt-get install -y build-essential gcc wget
RUN apt-get install -y pkg-config

# Sodium Library
WORKDIR /tmp
RUN wget https://download.libsodium.org/libsodium/releases/libsodium-1.0.3.tar.gz
RUN tar xvf libsodium-1.0.3.tar.gz

WORKDIR /tmp/libsodium-1.0.3
RUN ./configure
RUN make install

# ZeroMQ library
WORKDIR /tmp
RUN wget http://download.zeromq.org/zeromq-4.1.3.tar.gz
RUN tar xvf zeromq-4.1.3.tar.gz

WORKDIR zeromq-4.1.3
RUN ./configure
RUN make install

USER root
RUN make install
RUN ldconfig

# ZMQ ruby gem
RUN apt-get install libxml2-dev
RUN gem install ffi-rzmq

# Add the ZMQ APP
ADD ./app app

RUN ruby app/client.rb 127.0.0.1:1234 50 10 &
RUN ruby app/server.rb 127.0.0.1:1234 50 10
