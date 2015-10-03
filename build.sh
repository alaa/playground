#!/bin/sh

docker run -it -v $(pwd)/go:/root --entrypoint="/bin/bash" alaa/go-zmq /root/run.sh
