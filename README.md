This is the ultimate, frictionless "Notepad-style" README. It gives the user the exact keyboard shortcuts they need so they don't have to search for PowerShell.

I've also condensed the one-liner to be as short as possible while still being robust.

***

### 📄 README.md (Copy-Paste Ready)

```text
# Helium Updater for Windows 🎈

A zero-config utility that keeps your Helium installation updated in the background.

---

### 🚀 FAST INSTALL (30 Seconds)

1. Press **Windows + X** on your keyboard and select **Terminal (Admin)** or **PowerShell (Admin)**.
2. Click **Yes** on the User Account Control (UAC) prompt.
3. **Copy and Paste** the following command and press **Enter**:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; $repo='09sychic/helium_updater'; $temp="$env:TEMP\h_upd"; iwr "[https://github.com/$repo/archive/refs/heads/main.zip](https://github.com/$repo/archive/refs/heads/main.zip)" -OutFile "$temp.zip"; Expand-Archive "$temp.zip" -Dest $temp -Force; cd "$temp\*"; .\install.ps1
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
```

***

### 💡 Final developer check for you:
Make sure your `install.ps1` on GitHub has these specific lines to ensure the **"No Pop-up"** behavior works:

1.  **`go mod tidy`** (Run this before building to fetch the GitHub API library).
2.  **`go build -ldflags="-H windowsgui" -o $StartupExePath .`** (The `-H windowsgui` is what makes it "invisible" at startup).

If those are in your script, your users are going to have a very smooth experience!
