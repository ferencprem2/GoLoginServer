@echo off

set ExeName=webserver
echo Building %ExeName%...


set GOOS=linux

if not exist .\build (
	mkdir .\build
)

if exist .\build\%ExeName% (
	del .\build\%ExeName%
)

@echo on
go build -v -o .\build\%ExeName%

echo Done.

