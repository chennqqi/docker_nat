#!/bin/bash

go build -v
sudo docker build -t 'sort/docker_nat' -f Dockerfile.local .
