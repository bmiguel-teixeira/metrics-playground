#!/bin/bash

docker run --rm -p 9090:9090 -v ./prometheus.yml:/prometheus/prometheus.yml prom/prometheus --enable-feature=remote-write-receiver
