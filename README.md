# Helium Updater for Windows 🎈

A lightweight background utility that keeps your Helium installation up to date by checking the [official releases](https://github.com/imputnet/helium-windows) every time you log in.

---

### 🚀 Quick Install (One-Click)

Run **PowerShell** as **Administrator** and paste the command below. This will automatically install Go (if missing), build the updater, and set it to run at startup.

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex (Invoke-RestMethod 'https://raw.githubusercontent.com/ethantheb/helium_updater/main/install.ps1')
```

---

### 🛠️ How it works
1. **Automated Setup:** The script uses `winget` to fetch the Go compiler so you don't have to.
2. **Invisible Operation:** It builds the app with a `windowsgui` flag, meaning it runs silently in the background without a black command prompt popping up.
3. **Startup Integration:** The executable is placed in your Windows Startup folder (`AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`).
4. **Smart Updates:** It tracks your current version in `%LOCALAPPDATA%\helium_updater\.version` to avoid redundant downloads.

### 🗑️ Uninstall
To stop the updater and clean up all files, run:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('[https://raw.githubusercontent.com/ethantheb/helium_updater/main/uninstall.ps1](https://raw.githubusercontent.com/ethantheb/helium_updater/main/uninstall.ps1)'))
```
