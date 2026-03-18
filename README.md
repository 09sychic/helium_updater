
## Helium Updater for Windows

Simply grabs the latest release from github.com/imputnet/helium-windows
and installs when necessary.

---
### Install

First, install Go: https://go.dev/doc/install

Then run the install script: `./install.ps1`

This will build and place the executable in your users Startup folder, 
so it should run automatically at login.

Keeps a .version file in AppData to track current install.
