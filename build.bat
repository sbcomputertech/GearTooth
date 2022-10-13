@echo off

echo ---=== Cleaning up old files ===---
del /Q .\bin\*

echo ---=== Building all GearTooth components ===---

cd geartoothc
go build -o gtc.exe
copy gtc.exe ..\bin\gtc.exe
cd ..

cd geartoothpre
go build -o gtpp.exe
copy gtpp.exe ..\bin\gtpp.exe
cd ..
