package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type xy struct {
	x, y float64
}

func main() {
	var xys []xy
	f, err := os.Open("./data.txt")
	if err != nil {
		fmt.Errorf("unable to opem file%v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		xyline := scanner.Text()
		xys = append(xys, fmtxys(xyline))
		fmt.Println(xys)
	}

}

func fmtxys(xyline string) xy {
	var txy xy
	xyArray:= strings.Split(xyline, ",")
	txy.x= xyarr[0]
	txy.= 
}

func rxys(cyCompl []xy) plotter.XYs {
	xy1 := make(plotter.XYs, len(xyComppl))
	var err error
	xyArray := strings.Split(xyline, ",")
	for xy := range xyArray {
		xy1.x, err = strconv.ParseFloat(xyArray[0], 64)
		if err != nil {
			fmt.Errorf("unable to convert xy.x%v", err)
		}
		xy1.y, err = strconv.ParseFloat(xyArray[1], 64)
		if err != nil {
			fmt.Errorf("unable to convert xy.y%v", err)
		}
	}
	return xy
}
