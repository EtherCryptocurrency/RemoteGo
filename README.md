# RemoteGo
A remote control server written on Golang

# Install
git clone https://github.com/HaCk3Dq/RemoteGo.git

cd RemoteGo

go run remote.go

if u have *linux amd64*, u can use compiled one in RemoteGo/bin/remote

# Usage
It creates a server on a localhost:1337/linux
where u can post some commands, such as:

+ "+" - which makes volume louder by 5%
+ "-" - which makes volume quieter by 5%

**Music command compatible only with Foobnix, fix the code for ur music player**

+ "next" - turns the next music track
+ "prev" - turns the previous music track
+ "pause"
+ "play"