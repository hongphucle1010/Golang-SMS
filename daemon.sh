#!/bin/bash

echo "[DAEMON] Starting CompileDaemon..."

CompileDaemon \
 -build="./scripts/build.sh" \
 -command="./sms" \
 -exclude-dir=docs \
 -include="*.go"
