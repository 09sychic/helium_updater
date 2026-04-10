# 1. Check for Go and install via Winget if missing
if (!(Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Go not found. Installing Go via Winget..." -ForegroundColor Cyan
    winget install OpenJS.Go --silent --accept-package-agreements --accept-source-agreements
    
    # Refresh Path immediately
    $env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")
}

# 2. Setup Paths
$ProjectDir = $PSScriptRoot
$ExeName = "09sychic-HeliumSync.exe"
$StartupDir = [System.Environment]::GetFolderPath("Startup")
$StartupExePath = Join-Path $StartupDir $ExeName
$AppDataFolder = Join-Path $env:LOCALAPPDATA "09sychic-HeliumSync"

# 3. Security Exclusions (So Antivirus doesn't kill your background sync)
Write-Host "Configuring security exclusions..." -ForegroundColor Gray
Add-MpPreference -ExclusionPath $StartupExePath -ErrorAction SilentlyContinue
Add-MpPreference -ExclusionPath $AppDataFolder -ErrorAction SilentlyContinue

# 4. Stop existing instance to allow overwrite
Stop-Process -Name "09sychic-HeliumSync" -ErrorAction SilentlyContinue
Start-Sleep -Seconds 1

# 5. Build and Move
Write-Host "Building $ExeName..." -ForegroundColor Green
Push-Location $ProjectDir

# Ensure dependencies are fetched
go mod tidy

# -ldflags="-H windowsgui" hides the console window completely
go build -ldflags="-H windowsgui" -o $StartupExePath .

if ($LASTEXITCODE -ne 0) {
    Write-Error "Build failed. Ensure you are running PowerShell as Administrator."
    Pop-Location
    exit 1
}
Pop-Location

# 6. Start the app immediately in the background
Start-Process $StartupExePath

Write-Host "Done! 09sychic-HeliumSync is now active and in Startup." -ForegroundColor Green
