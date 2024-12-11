package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "image_processing",
    Short: "Image processing CLI application",
    Long:  `A CLI application for image processing tasks such as background reduction, dilation, and erosion.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}