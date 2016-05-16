#!/bin/sh

Run the server
```
sudo -u oren git clone https://github.com/leveros/leveros
sudo -u oren cd leveros
sudo -u oren make cli pull-docker-images
sudo -u oren make PROTOC_CMD="true" HAVE_GO="" cli pull-docker-images
make PROTOC_CMD="true" HAVE_GO="" install-cli
sudo -u oren make PROTOC_CMD="true" HAVE_GO="" make fastrun
```
