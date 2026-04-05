# Removes the executable from Startup, and deletes the helium_updater app data folder.

$StartupExe = Join-Path ([System.Environment]::GetFolderPath("Startup")) "09sychic-HeliumSync.exe"
$CacheDir = Join-Path $env:LOCALAPPDATA "helium_updater"

Remove-Item -Force $StartupExe -ErrorAction SilentlyContinue
Remove-Item -Recurse -Force $CacheDir -ErrorAction SilentlyContinue

Write-Host "Uninstalled."
