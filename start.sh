#!/bin/bash

export SONOS_URL=http://localhost:5015
export SONOS_LOGIN=admin
export SONOS_PASSWORD=password

go run main.go api.go
