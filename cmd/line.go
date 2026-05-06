/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/amaury95/toolbox/src/plotter"
	"github.com/spf13/cobra"
)

var _line_x_col string
var _line_y_col string
var _line_title string
var _line_compress_x bool

// lineCmd represents the line command
var lineCmd = &cobra.Command{
	Use:   "line",
	Short: "Plot a line chart from two CSV columns in the browser",
	Long: `Reads the CSV file given on the parent plot command (-p), takes the
header names for the X and Y series (-x / -y), and opens an interactive
line chart in your default browser. Use the chart toolbox to save the
figure as a PNG. The server keeps running until you stop the process (Ctrl+C).`,
	Run: func(cmd *cobra.Command, args []string) {
		points := plotter.GetLineData(_csv_file_path, _line_x_col, _line_y_col, _line_compress_x)
		plotter.PlotLine(_line_title, _line_x_col, _line_y_col, points...)
	},
}

func init() {
	plotCmd.AddCommand(lineCmd)

	lineCmd.Flags().StringVarP(&_line_x_col, "x", "x", "", "X column to plot")
	lineCmd.Flags().StringVarP(&_line_y_col, "y", "y", "", "Y column to plot")
	lineCmd.Flags().StringVarP(&_line_title, "title", "t", "Plotter", "Title of the plot")
	lineCmd.Flags().BoolVarP(&_line_compress_x, "compress-x", "c", false, "Compress the X axis")
}
