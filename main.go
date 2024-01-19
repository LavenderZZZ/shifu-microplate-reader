package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	endPointUrl  = "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"
	pollInterval = 10
)

func getMeasurement() ([][]float64, error) {
	resp, err := http.Get(endPointUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(body), "\n")

	matrix := make([][]float64, len(lines))
	for i, line := range lines {
		values := strings.Fields(line)
		matrix[i] = make([]float64, len(values))
		for j, value := range values {
			if matrix[i][j], err = strconv.ParseFloat(value, 64); err != nil {
				return nil, fmt.Errorf("parsing value %q: %v", value, err)
			}
		}
	}

	return matrix, nil
}

func calculateAverage(matrix [][]float64) float64 {
	var sum float64
	for _, row := range matrix {
		for _, val := range row {
			sum += val
		}
	}
	return sum / float64(len(matrix)*len(matrix[0]))
}

func main() {
	for {
		matrix, err := getMeasurement()
		if err != nil {
			fmt.Println("Error fetching measurements:", err)
			continue
		}
		average := calculateAverage(matrix)
		fmt.Printf("Average measurement: %f\n", average)

		time.Sleep(pollInterval * time.Second)
	}

}
