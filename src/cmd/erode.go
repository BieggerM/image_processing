package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/BieggerM/image_processing_golang/algorithms"
)

var (
    erodeInputFile string
    erodeRadius    int
)

var erodeCmd = &cobra.Command{
    Use:   "erode",
    Short: "Erode an image",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Eroding an image")
        algorithms.Erode(erodeInputFile, erodeRadius)
    },
}

func init() {
    rootCmd.AddCommand(erodeCmd)
    erodeCmd.Flags().StringVarP(&erodeInputFile, "input", "i", "", "Input image file (required)")
    erodeCmd.Flags().IntVarP(&erodeRadius, "radius", "r", 2, "Pixel Radius (required)")
    erodeCmd.MarkFlagRequired("input")
}