package cmd

import (
    "github.com/spf13/cobra"
    "github.com/BieggerM/image_processing_golang/algorithms"
)

/*
dilateCmd is a command that dilates an image
It uses the Dilate function from the algorithms package

The command has the following flags:
    --input (-i) : Input image file (required)
    --radius (-r) : Pixel Radius (required)
    --multithreaded (-m) : Use multithreaded version
    --numberofthreads (-n) : Number of threads
*/
var dilateCmd = &cobra.Command{
    Use:   "dilate",
    Short: "Dilate an image",
    Run: func(cmd *cobra.Command, args []string) {
        algorithms.Dilate(dilateInputFile, dilateRadius, multithreadedDilate, numberofthreadsDilate)
    },
}

var (
    dilateInputFile string
    dilateRadius    int
    multithreadedDilate   bool
    numberofthreadsDilate int
)

func init() {
    rootCmd.AddCommand(dilateCmd)
    dilateCmd.Flags().StringVarP(&dilateInputFile, "input", "i", "", "Input image file (required)")
    dilateCmd.Flags().IntVarP(&dilateRadius, "radius", "r", 2, "Pixel Radius (required)")
    dilateCmd.Flags().BoolVarP(&multithreadedDilate, "multithreaded", "m", false, "Use multithreaded version")
    dilateCmd.Flags().IntVarP(&numberofthreadsDilate, "numberofthreads", "n", 4, "Number of threads")
    dilateCmd.MarkFlagRequired("input")
}