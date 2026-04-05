# 1. Check for Go and install via Winget if missing
if (!(Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Go not found. Installing Go via Winget..." -ForegroundColor Cyan
    winget install OpenJS.Go --silent --accept-package-agreements --accept-source-agreements
    
    if ($LASTEXITCODE -ne 0) {
        Write-Error "Winget failed to install Go. Please install it manually from https://go.dev/dl/"
        exit 1
    }

    # Refresh Environment Variables for this session so 'go' command works immediately
    $env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")
}

# 2. Setup Paths
$ProjectDir = $PSScriptRoot
$ExeName = "09sychic-HeliumSync.exe"
$StartupDir = [System.Environment]::GetFolderPath("Startup")
$StartupExePath = Join-Path $StartupDir $ExeName

# 3. Build and Move
Write-Host "Building $ExeName..." -ForegroundColor Green
Push-Location $ProjectDir

# -ldflags="-H windowsgui" hides the console window when the updater runs at startup
go build -ldflags="-H windowsgui" -o $StartupExePath .

if ($LASTEXITCODE -ne 0) {
    Write-Error "Build failed. Make sure you are running as Administrator."
    Pop-Location
    exit 1
}
Pop-Location

Write-Host "Done! Helium Updater installed to Startup." -ForegroundColor Green
Write-Host "It will now run automatically every time you log in."
