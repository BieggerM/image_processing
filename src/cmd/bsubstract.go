package cmd

import (
	"github.com/BieggerM/image_processing_golang/algorithms"
	"github.com/spf13/cobra"
)

var (
	reference         string
	bgReduceInputFile string
	threshold         float64
	hsv               bool
	multitheaded	  bool
	numberofthreads   int
)

var bgReduceCmd = &cobra.Command{
	Use:   "bgsubstract",
	Short: "Substract background in an image",
	Run: func(cmd *cobra.Command, args []string) {
		algorithms.Background_subtract(reference, bgReduceInputFile, threshold, hsv, multitheaded, numberofthreads)
	},
}

func init() {
	rootCmd.AddCommand(bgReduceCmd)
	bgReduceCmd.Flags().StringVarP(&reference, "reference", "r", "", "Reference image file (required)")
	bgReduceCmd.Flags().StringVarP(&bgReduceInputFile, "input", "i", "", "Input image file (required)")
	bgReduceCmd.Flags().Float64VarP(&threshold, "threshold", "t", 50.0, "Threshold value for background reduction")
	bgReduceCmd.Flags().BoolVarP(&hsv, "hsv", "s", false, "Use HSV color space for comparison")
	bgReduceCmd.Flags().BoolVarP(&multitheaded, "multithreaded", "m", false, "Use multithreaded version")
	bgReduceCmd.Flags().IntVarP(&numberofthreads, "numberofthreads", "n", 4, "Number of threads")
	bgReduceCmd.MarkFlagRequired("reference")
	bgReduceCmd.MarkFlagRequired("input")
}
