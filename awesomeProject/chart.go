package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"github.com/wcharczuk/go-chart"
	"os"
	"time"
)

func main() {

	numValues := 2
	numSeries := 5
	series := make([]chart.Series, numSeries)

	for i := 0; i < numSeries; i++ {
		v, _ := mem.VirtualMemory()
		fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
		//fmt.Println(v)
		xValues := make([]time.Time, numValues)
		yValues := make([]float64, numValues)

		for j := 0; j < numValues; j++ {
			xValues[j] = time.Now().AddDate(0, 0, (numValues-j)*-1)
			yValues[j] = v.UsedPercent
		}

		series[i] = chart.TimeSeries{
			Name:    fmt.Sprintf("name-%v", i),
			XValues: xValues,
			YValues: yValues,
		}

		time.Sleep(10*time.Second)
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Time",
		},
		YAxis: chart.YAxis{
			Name: "Value",
		},
		Series: series,
	}

	f, _ := os.Create("/tmp/a.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
