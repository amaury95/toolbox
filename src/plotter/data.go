package plotter

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// GetLineData reads a CSV file whose first row is headers.
// It finds columns named xCol and yCol and returns one Point per data row (parsed as float64).
func GetLineData(csvFile, xCol, yCol string, compressX bool) []Point {
	f, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	if len(records) < 2 {
		return nil
	}

	header := records[0]
	xi := columnIndex(header, xCol)
	yi := columnIndex(header, yCol)
	if xi < 0 || yi < 0 {
		panic(fmt.Sprintf("plotter: csv columns %q and/or %q not found in header", xCol, yCol))
	}

	out := make([]Point, 0, len(records)-1)
	for _, row := range records[1:] {
		if len(row) <= xi || len(row) <= yi {
			continue
		}
		xStr := strings.TrimSpace(row[xi])
		yStr := strings.TrimSpace(row[yi])
		if xStr == "" && yStr == "" {
			continue
		}
		x, err := strconv.ParseFloat(xStr, 64)
		if err != nil {
			panic(fmt.Errorf("plotter: parse %q column %q: %w", xCol, xStr, err))
		}
		y, err := strconv.ParseFloat(yStr, 64)
		if err != nil {
			panic(fmt.Errorf("plotter: parse %q column %q: %w", yCol, yStr, err))
		}
		out = append(out, Point{X: x, Y: y})
	}
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].X < out[j].X
	})
	if compressX {
		for i := 1; i < len(out); i++ {
			out[i].X -= out[0].X
		}
		out[0].X = 0
	}
	return out
}

func columnIndex(header []string, name string) int {
	for i, h := range header {
		if strings.TrimSpace(h) == name {
			return i
		}
	}
	return -1
}
