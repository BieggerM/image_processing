package cmd

import (
    "github.com/spf13/cobra"
    "github.com/BieggerM/image_processing_golang/algorithms"
)

var dilateCmd = &cobra.Command{
    Use:   "dilate",
    Short: "Dilate an image",
    Run: func(cmd *cobra.Command, args []string) {
        algorithms.Dilate(dilateInputFile, dilateRadius)
    },
}

var (
    dilateInputFile string
    dilateRadius    int
)

func init() {
    rootCmd.AddCommand(dilateCmd)
    dilateCmd.Flags().StringVarP(&dilateInputFile, "input", "i", "", "Input image file (required)")
    dilateCmd.Flags().IntVarP(&dilateRadius, "radius", "r", 2, "Pixel Radius (required)")
    dilateCmd.MarkFlagRequired("input")
}