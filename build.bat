@echo off

echo ---=== Cleaning up old files ===---
del /Q .\bin\*

echo ---=== Building all GearTooth components ===---

cd geartoothc
go build -buildmode=exe -o gtc.exe
copy gtc.exe ..\bin\gtc.exe
cd ..

cd geartoothpre
go build -buildmode=exe -o gtpp.exe
copy gtpp.exe ..\bin\gtpp.exe
cd ..
