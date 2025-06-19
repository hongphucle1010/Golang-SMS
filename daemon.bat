@echo off
echo [DAEMON] Starting CompileDaemon...

CompileDaemon ^
 -build=".\scripts\build.bat" ^
 -command="./sms" ^
 -exclude-dir=docs ^
 -include=*.go
