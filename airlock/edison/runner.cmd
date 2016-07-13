@echo off
if [%1] == [] goto usage
if [%2] == [] goto usage

rem If you're using putty, change these to pscp and plink
set SCP=scp
set SSH=ssh

:good
if [%3] == [runonly] goto run
echo "Compiling..."
set GOARCH=386
set GOOS=linux
go build %1.go
if errorlevel 1 goto exit
echo "Copying..."
%SCP% %1 root@%2:/home/root/%1
echo "Setting permissions..."
%SSH% -t root@%2 chmod +x ./%1
:run
echo "Running..."
%SSH% -t root@%2 ./%1
goto exit

:usage
echo Usage %0 [go file] [target host] [runonly]
echo     Do not add the .go extension for [go file]
echo     target host may be a host name or an ip address
echo     Specify runonly to skip compilation and copying

:exit