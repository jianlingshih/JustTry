package DataAnalysis

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"nncms/utils"
	"os"
)

// ingest is a function that ingests the file and outputs the header, data, and index.
func ingest(f io.Reader) (header []string, data [][]string, indices []map[string][]int, err error) {
	r := csv.NewReader(f)

	// handle header
	if header, err = r.Read(); err != nil {
		return
	}

	indices = make([]map[string][]int, len(header))
	var rowCount, colCount int = 0, len(header)
	for rec, err := r.Read(); err == nil; rec, err = r.Read() {
		if len(rec) != colCount {
			return nil, nil, nil, fmt.Errorf("expected Columns: %d. Got %d columns in row %d", colCount, len(rec), rowCount)
		}
		data = append(data, rec)
		for j, val := range rec {

			if indices[j] == nil {
				indices[j] = make(map[string][]int)
			}
			indices[j][val] = append(indices[j][val], rowCount)
		}
		rowCount++
	}
	return
}
func mHandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func handleCsv() {
	f, err := os.Open(utils.GetCurrentPath() + "/train.csv")
	mHandleErr(err)
	hdr, data, indices, err := ingest(f)

	mHandleErr(err)
	c := cardinality(indices)

	fmt.Printf("Original Data: \nRows: %d, Cols: %d\n========\n", len(data), len(hdr))
	//c := cardinality(indices)
	for i, h := range hdr {
		fmt.Printf("%v: %v\n", h, c[i])
	}
	fmt.Println("")

}
func cardinality(indices []map[string][]int) []int {
	retVal := make([]int, len(indices))
	for i, m := range indices {
		retVal[i] = len(m)
	}
	return retVal
}
