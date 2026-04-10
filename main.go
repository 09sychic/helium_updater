package main

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
    "syscall"
    "time"

    "github.com/google/go-github/v57/github"
)

// killHelium targets the specific chrome.exe within the Helium application path
func killHelium() {
    fmt.Println("Closing Helium to prepare for update...")
    
    // Using WMIC to target only the chrome.exe inside the Helium directory
    // This prevents closing the user's actual Google Chrome browser
    targetPath := `C:\Users\ADMIN\AppData\Local\imput\Helium\Application\chrome.exe`
    
    // Build the command: wmic process where ExecutablePath='...' call terminate
    cmd := exec.Command("wmic", "process", "where", fmt.Sprintf("ExecutablePath='%s'", strings.ReplaceAll(targetPath, `\`, `\\`)), "call", "terminate")
    
    // Hide the black console window when running the kill command
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    _ = cmd.Run()

    // Wait 2 seconds for Windows to release file locks
    time.Sleep(2 * time.Second)
}

func downloadFile(filepath string, url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}

func userCachePath() (string, error) {
    cacheDir, err := os.UserCacheDir()
    if err != nil {
        return "", err
    }
    dir := filepath.Join(cacheDir, "09sychic-HeliumSync")
    if err := os.MkdirAll(dir, 0755); err != nil {
        return "", err
    }
    return filepath.Join(dir, ".version"), nil
}

func readLastVersion() string {
    path, err := userCachePath()
    if err != nil {
        return ""
    }
    data, err := os.ReadFile(path)
    if err != nil {
        return ""
    }
    return strings.TrimSpace(string(data))
}

func main() {
    if runtime.GOOS != "windows" {
        fmt.Println("This program is intended to run on Windows.")
        os.Exit(1)
    }

    arch := runtime.GOARCH
    ctx := context.Background()
    client := github.NewClient(nil)

    release, _, err := client.Repositories.GetLatestRelease(ctx, "imputnet", "helium-windows")
    if err != nil {
        panic(err)
    }

    var downloadURL string
    var downloadName string

    for _, asset := range release.Assets {
        name := strings.ToLower(asset.GetName())
        match := false
        switch arch {
            case "amd64":
                match = strings.Contains(name, "x64-installer")
            case "arm64":
                match = strings.Contains(name, "arm64-installer")
        }

        if match {
            downloadURL = asset.GetBrowserDownloadURL()
            downloadName = asset.GetName()
            break
        }
    }

    if downloadURL == "" {
        fmt.Println("No download found for architecture:", arch)
        os.Exit(1)
    }

    lastVersion := readLastVersion()
    if lastVersion == downloadName {
        fmt.Println("No new version to install.")
        time.Sleep(2 * time.Second)
        return
    }

    vPath, err := userCachePath()
    if err != nil {
        fmt.Println("Error: could not determine AppData path:", err)
        os.Exit(1)
    }
    downloadPath := filepath.Join(filepath.Dir(vPath), downloadName)

    fmt.Println("Downloading:", downloadName)
    if err := downloadFile(downloadPath, downloadURL); err != nil {
        panic(err)
    }

    // --- CRITICAL STEP: KILL PROCESS BEFORE INSTALLING ---
    killHelium()
    // -----------------------------------------------------

    fmt.Printf("Installing...")

    filePath, err := filepath.Abs(downloadPath)
    if err != nil {
        panic(err)
    }

    // Run installer silently
    cmd := exec.Command(filePath, "/S") // Added /S for silent install if supported
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        panic(err)
    }

    if err := os.Remove(filePath); err != nil {
        fmt.Println("Warning: could not delete installer:", err)
    }

    if err := os.WriteFile(vPath, []byte(downloadName), 0644); err != nil {
        fmt.Println("Warning: could not write version file:", err)
    }

    fmt.Println(" done.")
}
