package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/BieggerM/image_processing_golang/algorithms"
)

/*
erodeCmd is a command that erodes an image
It uses the Erode function from the algorithms package

The command has the following flags:
    --input (-i) : Input image file (required)
    --radius (-r) : Pixel Radius (required)
    --multithreaded (-m) : Use multithreaded version
    --numberofthreads (-n) : Number of threads
*/

var erodeCmd = &cobra.Command{
    Use:   "erode",
    Short: "Erode an image",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Eroding an image")
        algorithms.Erode(erodeInputFile, erodeRadius, multithreadedErode, numberofthreadsErode)
    },
}

var (
    erodeInputFile string
    erodeRadius    int
    multithreadedErode   bool
    numberofthreadsErode int
)

func init() {
    rootCmd.AddCommand(erodeCmd)
    erodeCmd.Flags().StringVarP(&erodeInputFile, "input", "i", "", "Input image file (required)")
    erodeCmd.Flags().IntVarP(&erodeRadius, "radius", "r", 2, "Pixel Radius (required)")
    erodeCmd.Flags().BoolVarP(&multithreadedErode, "multithreaded", "m", false, "Use multithreaded version")
    erodeCmd.Flags().IntVarP(&numberofthreadsErode, "numberofthreads", "n", 4, "Number of threads")
    erodeCmd.MarkFlagRequired("input")
}