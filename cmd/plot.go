/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var _csv_file_path string

// plotCmd represents the plot command
var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "Visualize data from a CSV file in the browser",
	Long: `Commands for plotting data from a CSV. Use -p / --path to set the file
once; it is inherited by subcommands (for example, plot line). Run a
subcommand to open a chart; calling plot by itself only shows this help.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plot called")
	},
}

func init() {
	rootCmd.AddCommand(plotCmd)

	plotCmd.PersistentFlags().StringVarP(&_csv_file_path, "path", "p", "", "CSV file path to plot")
}
