package ui

import (
    "fmt"
    "time"
    "github.com/briandowns/spinner"
    "github.com/fatih/color"
    "github.com/mmycin/npmdc/pkg/api"
)

func DisplayNpmPackageInfo(info *api.NpmPackageInfo) {
    // Example loading spinner for demonstration
    s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
    s.Start()
    time.Sleep(2 * time.Second) // simulate some waiting
    s.Stop()

    fmt.Println("\nPackage Information:")
    fmt.Println("---------------------")
    fmt.Printf("%-20s: %s\n", color.CyanString("Name"), info.Name)
    fmt.Printf("%-20s: %s\n", color.YellowString("Version"), info.Version)
    fmt.Printf("%-20s: %s\n", color.GreenString("Size"), info.UnpackedSizeKB)
    fmt.Printf("%-20s: %s\n", color.BlueString("Downloads"), info.WeeklyDownloads)
}
