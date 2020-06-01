package linearregression

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// RegressionData for collecting some data that needed in linear regression calculation
type RegressionData struct {
	XValues      []float64
	YValues      []float64
	Epoch        int
	LearningRate float64
	slope        float64
	intercept    float64
}

// New ...
func New() *RegressionData {
	rd := &RegressionData{}
	return rd
}

// PlotTheDataset will make the scatter plot based on x and y values
func (d *RegressionData) PlotTheDataset() {
	d.plotMaker(false)
}

// Calculate to calculate the linear regression
func (d *RegressionData) Calculate(isWriteResultToPlot bool) (float64, float64) {
	if len(d.XValues) == 0 || len(d.YValues) == 0 {
		panic("x values and y values cannot be empty")
	}
	for i := 0; i < d.Epoch; i++ {
		d.slope, d.intercept = d.gradientDescent(d.slope, d.intercept)
	}
	//round float to 2 decimal
	d.slope = math.Round(d.slope*100) / 100
	d.intercept = math.Round(d.intercept*100) / 100
	if isWriteResultToPlot {
		d.plotMaker(true)
	}
	return d.slope, d.intercept
}

func (d *RegressionData) gradientDescent(mCurrent float64, bCurrent float64) (finalM float64, finalB float64) {
	var mGradient float64
	var bGradient float64
	nData := len(d.XValues) //the x or y values length
	twoPerN := float64(2) / float64(nData)
	for i := 0; i < nData; i++ {
		mGradient += -twoPerN * d.XValues[i] * (d.YValues[i] - ((mCurrent * d.XValues[i]) + bCurrent))
		bGradient += -twoPerN * (d.YValues[i] - ((mCurrent * d.XValues[i]) + bCurrent))
	}
	finalM = mCurrent - (d.LearningRate * mGradient)
	finalB = bCurrent - (d.LearningRate * bGradient)
	return
}

func (d *RegressionData) plotMaker(isWriteResultToPlot bool) {
	file := ""
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Linear Regression using Go"
	p.X.Min = 0
	p.X.Max = 6
	p.X.Label.Text = "X"
	p.Y.Min = 0
	p.Y.Max = 6
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	dataset := make(plotter.XYs, len(d.XValues))
	for i := range dataset {
		dataset[i].X = d.XValues[i]
		dataset[i].Y = d.YValues[i]
	}
	s, err := plotter.NewScatter(dataset)
	if err != nil {
		log.Panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	if isWriteResultToPlot {
		lineData := make(plotter.XYs, 7)
		for i := 0; i < 7; i++ {
			x := float64(i)
			lineData[i].X = x
			lineData[i].Y = d.slope*x + d.intercept
		}
		l, err := plotter.NewLine(lineData)
		if err != nil {
			log.Panic(err)
		}
		l.LineStyle.Width = vg.Points(1)
		l.LineStyle.Color = color.RGBA{B: 255, A: 255}

		p.Add(s, l)
		p.Legend.Add("dataset", s)

		p.Legend.Add(fmt.Sprintf("y(x) = %.2fx + %.2f", d.slope, d.intercept), l)
		file = "plot/result.png"
	} else {
		p.Add(s)
		p.Legend.Add("dataset", s)
		file = "plot/scatter.png"
	}

	err = p.Save(300, 300, file)
	if err != nil {
		log.Panic(err)
	}
}
