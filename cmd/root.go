/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/JunNishimura/okyo/internal/speaker"
	"github.com/spf13/cobra"
)

var (
	loopCount int
	playSpeed float64
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "okyo",
	Short: "A cli tool to chant a hannya shingyo",
	Long: "A cli tool to chant a hannya shingyo",
	Run: func(cmd *cobra.Command, args []string) {
		speaker := speaker.NewMp3Speaker(loopCount, playSpeed)
		speaker.Play()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&loopCount, "count", "c", 1, "The number of times to play the audio(-1 means infinite loop)")
	rootCmd.Flags().Float64VarP(&playSpeed, "speed", "s", 1, "The speed of the audio") 
}

