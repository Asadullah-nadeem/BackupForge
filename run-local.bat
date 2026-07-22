@echo off
setlocal

:: Add local portable Go bin to PATH so sub-processes like goose can be found
set "PATH=%~dp0go-dist\go\bin;%PATH%"


:: Locate Go compiler (prefer portable local Go in go-dist)
set "LOCAL_GO=%~dp0go-dist\go\bin\go.exe"
if exist "%LOCAL_GO%" (
    set "GO_EXE=%LOCAL_GO%"
    echo [INFO] Using local portable Go version at %LOCAL_GO%
) else (
    where go >nul 2>nul
    if %errorlevel% equ 0 (
        set "GO_EXE=go"
    ) else (
        echo [ERROR] Go compiler not found.
        echo Please download and extract Go to %~dp0go-dist\go or install it in system PATH.
        pause
        exit /b 1
    )
)

:: Locate Redis Server (prefer portable local Redis in redis-dist)
set "LOCAL_REDIS=%~dp0redis-dist\redis-server.exe"
if exist "%LOCAL_REDIS%" (
    echo [INFO] Starting local Redis cache server...
    start "Redis Server" "%LOCAL_REDIS%"
) else (
    echo [WARN] Local portable Redis not found at %LOCAL_REDIS%. 
    echo        Ensure Valkey/Redis is already running on port 6379 on your machine.
)

echo ==========================================================
echo Building Frontend...
echo ==========================================================
echo.

cd frontend
if not exist "node_modules" (
    echo [INFO] Installing frontend dependencies...
    call pnpm install
)
echo [INFO] Compiling frontend assets...
call pnpm run build
echo [INFO] Copying compiled assets to Go backend...
if exist "..\backend\ui\build" rd /s /q "..\backend\ui\build"
xcopy /E /I /Y "out" "..\backend\ui\build\"
cd ..
    
echo.
echo ==========================================================
echo Starting Go Backend locally...
echo ==========================================================
echo.

cd backend
"%GO_EXE%" run ./cmd
if %errorlevel% neq 0 (
    echo.
    echo [ERROR] The Go backend exited with code %errorlevel%.
    pause
)
