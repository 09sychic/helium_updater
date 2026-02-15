
### Helium (Windows) Updater

Simply grabs the latest release from github.com/imputnet/helium-windows and installs.

After its been run once, it will keep a .version file to avoid re-downloading the same version.

---
If you download the executable Windows Defender will probably mark it as a virus. You can
either choose to trust me and add an exception, or you could just build it yourself.

Install Go: https://go.dev/doc/install

Download or clone the repository, then cd in and run:
```
go build
```

---
### Run on login
To run the updater automagically, you can simply add the .exe to your Startup folder.

Open the Run dialog (Windows + R), type `shell:startup` and then drag and drop the exe or a shortcut into this folder.
