#!/bin/bash

echo "[CLEAN] Removing existing docs folder..."
rm -rf docs

echo "[SWAG] Generating Swagger docs..."
swag init

echo "[BUILD] Building Go application..."
go build -o sms cmd/main.go
