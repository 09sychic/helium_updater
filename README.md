
## Helium (Windows) Updater

Simply grabs the latest release from github.com/imputnet/helium-windows and installs.

After its been run once, it will keep a .version file to avoid re-downloading the same version.

---
### Build yourself
Install Go: https://go.dev/doc/install

Download or clone the repository, then cd in and run `go build`.

---
There are prebuilt executables available in Releases (if you're a risk-taker), but they're unsigned so Windows Defender will
flag them. Defender might even flag your own executable. To add an exception:

> Windows Security > Virus & Threat Protection > Manage settings > Exclusions > Add > _path to helium_updater.exe_

---
### Run on login
To run the updater automagically, you can simply add the .exe to your Startup folder.

Open the Run dialog (Windows + R), type `shell:startup` and then drag and drop the exe or a shortcut into this folder.
