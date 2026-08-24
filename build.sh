#!/usr/bin/env bash

mkdir -pv ./heartbeat

GOOS=linux GOARCH=amd64 go build -o ./heartbeat/heart_beat_amd64 main.go
GOOS=linux GOARCH=arm64 go build -o ./heartbeat/heart_beat_arm64 main.go
