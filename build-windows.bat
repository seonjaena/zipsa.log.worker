@ECHO off

IF "%1" == "" (
    ECHO "Please Input Parameter"
    GOTO end
) ELSE (
    SET PROFILE=%1
)

SET ZIP_DIR=C:\Users\sky11\OneDrive\Desktop
SET MAIN_GO_FILE=main.go
SET MAIN_FILE=main.exe
SET MAIN_ZIP_FILE=main.zip
SET GOARCH=amd64
SET GOOS=windows

if exist %MAIN_FILE% (
    del %MAIN_FILE%
)
if exist %MAIN_ZIP_FILE% (
    del %MAIN_ZIP_FILE%
)
if exist %ZIP_DIR%\%MAIN_ZIP_FILE% (
    del %ZIP_DIR%\%MAIN_ZIP_FILE%
)

echo "PROFILE = %PROFILE%"
echo "GOARCH = %GOARCH%"
echo "GOOS = %GOOS%"
go build -ldflags="-X zipsa.log.worker/properties.Profile=%PROFILE%" -o %MAIN_FILE% %MAIN_GO_FILE%

zip %MAIN_ZIP_FILE% %MAIN_FILE%

move %MAIN_ZIP_FILE% %ZIP_DIR%

:end