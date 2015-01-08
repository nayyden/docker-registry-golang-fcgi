#!/bin/bash

docker run -v "$PWD:/app" golang:1.3 bash -c "cd /app/src && go build -o /app/public/index.bin"
