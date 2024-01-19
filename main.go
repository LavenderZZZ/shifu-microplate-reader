package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	var matrix [][]float64
	err = json.NewDecoder(resp.Body).Decode(&matrix)
	if err != nil {
		return nil, err
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
		fmt.Print("Average measurement: %f\n", average)

		time.Sleep(pollInterval * time.Second)
	}

}
