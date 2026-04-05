
# Helium Updater for Windows 🎈

A zero-config utility that keeps your Helium installation updated in the background.

---

### 🚀 FAST INSTALL (30 Seconds)

1. Press **Windows + X** on your keyboard and select **Terminal (Admin)** or **PowerShell (Admin)**.
2. Click **Yes** on the User Account Control (UAC) prompt.
3. **Copy and Paste** the following command and press **Enter**:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; Invoke-WebRequest -Uri "https://github.com/09sychic/helium_updater/archive/refs/heads/main.zip" -OutFile "$env:TEMP\h_upd.zip"; Expand-Archive -Path "$env:TEMP\h_upd.zip" -DestinationPath "$env:TEMP\h_upd_folder" -Force; Set-Location "$env:TEMP\h_upd_folder\helium_updater-main"; .\install.ps1
```

---

### 🛠️ WHAT HAPPENS AUTOMATICALLY?
* **Downloads** the source code from this repository.
* **Installs Go** (the compiler) via Windows 'winget' if you don't have it.
* **Builds** the updater as a "Ghost" process (no window will ever pop up).
* **Startup**: Sets the app to run every time you log in to Windows.

---

### 🗑️ UNINSTALL
To stop the updater and delete all its files, run this in Admin PowerShell:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; iex (Invoke-RestMethod '[https://raw.githubusercontent.com/09sychic/helium_updater/main/uninstall.ps1](https://raw.githubusercontent.com/09sychic/helium_updater/main/uninstall.ps1)')
```
