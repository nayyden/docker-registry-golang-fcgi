This is standalone version of the golang implementation of the docker registry server.
The original code can be found here https://github.com/docker/docker-registry/tree/0.9.0/contrib/golang_impl

The code is licensed under Apache License, Version 2.0, January 2004 which you can find here:
http://www.apache.org/licenses/

**All credits go to the authors of docker-registry -> contrib -> golang_impl.**

This is slightly modified version of the registry so we can run it under apache+fastcgi set-up.

### Compiling
The source code is located in src. It can be compiled with the following command
``` bash
cd src/ && go build -o ../public/index.bin
```
Alternatively, you can run the go compiler inside docker. There is already script for this in the main dir:
``` bash
./compile.sh
```

### Running
To start the registry inside docker, run:
``` bash
# Give permission to the apache user to write in the local storage (might require sudo on some systems)
chown www-data -R data/docker_index/

# Build and run
fig build
fig up -d example

# Dev container is also configured so rebuild is not needed on every compilation
fig up -d dev
```

### Deploying in non-docker environment
Most hosting companies already offer fastcgi support,
so it should be enough to copy .htaccess, conf.json and index.bin from public/ to your web directory.

### Known problems
None yet