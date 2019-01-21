#!/usr/bin/env bash

docker build -t sonatype .

docker run -p 8081:8081 -p 5005:5005 -v /srv/docker/nexus:/nexus-data sonatype