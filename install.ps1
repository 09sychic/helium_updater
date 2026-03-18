# This script simply builds the executable and places it
# in the users startup directory. This forces it to run at login.

$ProjectDir = $PSScriptRoot
$ExeName = "helium_updater.exe"
$StartupDir = [System.Environment]::GetFolderPath("Startup")
$StartupExePath = Join-Path $StartupDir $ExeName

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Error "Go is not installed or not in PATH. Download it from https://go.dev/dl"
    exit 1
}

Write-Host "Building $ExeName..."
Push-Location $ProjectDir
go build -o $StartupExePath .
if ($LASTEXITCODE -ne 0) {
    Write-Error "Build failed."
    Pop-Location
    exit 1
}
Pop-Location

Write-Host "Installed to: $StartupExePath"
