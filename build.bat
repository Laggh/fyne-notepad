@echo off
go build -ldflags "-H=windowsgui" -o go-notepad.exe .
echo Build complete for go-notepad.exe
pause