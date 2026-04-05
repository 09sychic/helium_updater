
# 09sychic-HeliumSync [Updater] for Windows 🎈

A zero-config utility that keeps your Helium installation updated in the background. 
This is a fork of the original tool by ethantheb.

---

### Install

1. Press **Windows + X** and select **Terminal (Admin)** or **PowerShell (Admin)**.
2. Click **Yes** for the UAC prompt.
3. Paste the following command and press **Enter**:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; Invoke-WebRequest -Uri "[https://github.com/09sychic/09sychic-HeliumSync/archive/refs/heads/main.zip](https://github.com/09sychic/09sychic-HeliumSync/archive/refs/heads/main.zip)" -OutFile "$env:TEMP\h_sync.zip"; Expand-Archive -Path "$env:TEMP\h_sync.zip" -DestinationPath "$env:TEMP\h_sync_folder" -Force; Set-Location "$env:TEMP\h_sync_folder\09sychic-HeliumSync-main"; .\install.ps1
````

-----

### What happens when pasted?

  * **Downloads** the source code from this repository.
  * **Installs Go** (the compiler) via Windows 'winget' if you don't have it.
  * **Builds** the updater as a "Ghost" process (no window will ever pop up).
  * **Startup**: Sets the app to run every time you log in to Windows.

-----

### Credits

  * **Original Developer:** [ethantheb](https://github.com/ethantheb) (Original repo: `helium_updater`)
  * **Maintained by:** [09sychic](https://github.com/09sychic)

-----

### Uninstall

To stop the updater and delete all its files, run this in Admin PowerShell:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; iex (Invoke-RestMethod '[https://raw.githubusercontent.com/09sychic/helium_updater/main/uninstall.ps1](https://raw.githubusercontent.com/09sychic/helium_updater/main/uninstall.ps1)')
```


### 📜 Credits
Special thanks to **[ethantheb](https://github.com/ethantheb)** for the original implementation.


