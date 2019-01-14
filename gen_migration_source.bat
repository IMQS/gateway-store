@echo off
del /f bindata.go
cd database/scripts
go-bindata -pkg database -o ../bindata.go .
cd..
cd..
