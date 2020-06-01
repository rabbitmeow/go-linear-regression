package main

import (
	"fmt"
	"go-linear-regression/linearregression"
)

func main() {
	xValues := []float64{1, 2, 3, 4, 5}
	yValues := []float64{1, 3, 2, 3, 5}
	lr := linearregression.New()
	lr.XValues = xValues
	lr.YValues = yValues
	lr.Epoch = 200
	lr.LearningRate = 0.05
	lr.PlotTheDataset()
	m, b := lr.Calculate(true)
	fmt.Println("m (slope) = ", m)
	fmt.Println("b (intercept) = ", b)
	fmt.Printf("so, the function is y(x) = %.2f x + %.2f", m, b)
}
