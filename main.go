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

    "github.com/google/go-github/v57/github"
)

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

func readLastVersion() string {
    data, err := os.ReadFile(".version")
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

    // Find the correct executable
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
        return
    }

    fmt.Println("Downloading:", downloadName)
    if err := downloadFile(downloadName, downloadURL); err != nil {
        panic(err)
    }

    fmt.Printf("Installing...")

    filePath, err := filepath.Abs(downloadName)
    if err != nil {
        panic(err)
    }

    cmd := exec.Command(filePath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        panic(err)
    }

    if err := os.Remove(filePath); err != nil {
        fmt.Println("Warning: could not delete file:", err)
        os.Exit(1)
    }

    if err := os.WriteFile(".version", []byte(downloadName), 0644); err != nil {
        fmt.Println("Warning: could not write version file:", err)
        os.Exit(1)
    }

    fmt.Println(" done.")
}
