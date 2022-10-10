@echo off
go build
geartooth about
echo ---
geartooth compile test.gt
echo --- running ---
python gtre\run_gtil.py hw.gtil