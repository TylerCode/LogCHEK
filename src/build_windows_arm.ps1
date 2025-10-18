param(
    [string]$LlvmMingwPath = "C:\llvm-mingw",
    [switch]$SkipDependencyCheck,
    [switch]$SkipRCEdit,
    [string]$Version = "0.3.1.0"
)

Write-Host "LogCHEK Windows ARM64 Build Script" -ForegroundColor Cyan
Write-Host "=====================================" -ForegroundColor Cyan
Write-Host ""

# Check if we're on ARM64
$arch = $env:PROCESSOR_ARCHITECTURE
if ($arch -ne "ARM64") {
    Write-Host "Warning: This script is designed for ARM64 Windows. Current architecture: $arch" -ForegroundColor Yellow
    $continue = Read-Host "Continue anyway? (y/n)"
    if ($continue -ne "y") {
        exit 1
    }
}

# Check if llvm-mingw exists
$gccPath = Join-Path $LlvmMingwPath "bin\aarch64-w64-mingw32-gcc.exe"
if (-not (Test-Path $gccPath)) {
    Write-Host "Error: llvm-mingw not found at $LlvmMingwPath" -ForegroundColor Red
    Write-Host ""
    Write-Host "Please install llvm-mingw:" -ForegroundColor Yellow
    Write-Host "1. Download from: https://github.com/mstorsjo/llvm-mingw/releases" -ForegroundColor Yellow
    Write-Host "   (Get the llvm-mingw-*-ucrt-aarch64.zip file)" -ForegroundColor Yellow
    Write-Host "2. Extract to $LlvmMingwPath" -ForegroundColor Yellow
    Write-Host "3. Run this script again" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "Or specify a different path:" -ForegroundColor Yellow
    Write-Host "  .\build_windows_arm.ps1 -LlvmMingwPath ""C:\your\path""" -ForegroundColor Yellow
    exit 1
}

Write-Host "Found llvm-mingw at: $LlvmMingwPath" -ForegroundColor Green

# Set up environment
$env:PATH = "$LlvmMingwPath\bin;$env:PATH"
$env:CC = "aarch64-w64-mingw32-gcc"
$env:CXX = "aarch64-w64-mingw32-g++"
$env:CGO_ENABLED = "1"

Write-Host "Compiler set to: $env:CC" -ForegroundColor Green

# Verify Go is installed
try {
    $goVersion = go version 2>&1
    Write-Host "Go version: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "Error: Go is not installed or not in PATH" -ForegroundColor Red
    exit 1
}

# Check dependencies (unless skipped)
if (-not $SkipDependencyCheck) {
    Write-Host ""
    Write-Host "Checking dependencies..." -ForegroundColor Cyan
    go mod download
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Error: Failed to download dependencies" -ForegroundColor Red
        exit 1
    }
}

# Build GUI version
Write-Host ""
Write-Host "Building LogCHEK-GUI.exe..." -ForegroundColor Cyan
go build -ldflags -H=windowsgui -o LogCHEK-GUI.exe main.go
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to build GUI version" -ForegroundColor Red
    exit 1
}
Write-Host "✓ Built LogCHEK-GUI.exe" -ForegroundColor Green

# Build CLI version
Write-Host ""
Write-Host "Building LogCHEK.exe..." -ForegroundColor Cyan
go build -o LogCHEK.exe cli.go
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: Failed to build CLI version" -ForegroundColor Red
    exit 1
}
Write-Host "✓ Built LogCHEK.exe" -ForegroundColor Green

# RCEdit (if available and not skipped)
if (-not $SkipRCEdit) {
    $rceditPath = "..\rcedit.exe"
    if (Test-Path $rceditPath) {
        Write-Host ""
        Write-Host "Applying resource edits..." -ForegroundColor Cyan
        
        & $rceditPath LogCHEK-GUI.exe --set-version-string "FileDescription" "Log file checker."
        & $rceditPath LogCHEK-GUI.exe --set-version-string "ProductName" "LogCHEK"
        & $rceditPath LogCHEK-GUI.exe --set-file-version $Version
        & $rceditPath LogCHEK-GUI.exe --set-product-version $Version
        & $rceditPath LogCHEK-GUI.exe --set-icon "logo.ico"
        
        & $rceditPath LogCHEK.exe --set-version-string "FileDescription" "Log file checker CLI."
        & $rceditPath LogCHEK.exe --set-version-string "ProductName" "LogCHEK"
        & $rceditPath LogCHEK.exe --set-file-version $Version
        & $rceditPath LogCHEK.exe --set-product-version $Version
        & $rceditPath LogCHEK.exe --set-icon "logo.ico"
        
        Write-Host "✓ Applied resource edits" -ForegroundColor Green
    } else {
        Write-Host "Note: rcedit.exe not found, skipping resource edits" -ForegroundColor Yellow
    }
}

Write-Host ""
Write-Host "Build completed successfully!" -ForegroundColor Green
Write-Host ""
Write-Host "Output files:" -ForegroundColor Cyan
Write-Host "  - LogCHEK-GUI.exe (GUI version)" -ForegroundColor White
Write-Host "  - LogCHEK.exe (CLI version)" -ForegroundColor White
