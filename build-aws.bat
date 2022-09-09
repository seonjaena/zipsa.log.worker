@ECHO off

SET ZIP_DIR=C:\Users\sky11\OneDrive\Desktop
SET MAIN_FILE=main
SET MAIN_ZIP_FILE=main.zip
SET GOARCH=amd64
SET GOOS=linux

if exist %MAIN_FILE% (
    del %MAIN_FILE%
)
if exist %MAIN_ZIP_FILE% (
    del %MAIN_ZIP_FILE%
)
if exist %ZIP_DIR%\%MAIN_ZIP_FILE% (
    del %ZIP_DIR%\%MAIN_ZIP_FILE%
)

echo "GOARCH = %GOARCH%"
echo "GOOS = %GOOS%"
go build -o main main.go

zip %MAIN_ZIP_FILE% %MAIN_FILE%

move %MAIN_ZIP_FILE% %ZIP_DIR%