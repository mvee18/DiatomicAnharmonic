package main

import (
	"go-hep.org/x/hep/fit"
	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
)

var (
	a  = 1.0
	b  = 2.0
	ps = []float64{a, b}
)

var poly = func(x float64, ps []float64) float64 {
	return ps[0] + ps[1]*x*x
}

var xdata = []float64{0, 1, 2, 3, 4, 5}
var ydata = []float64{0, 1, 4, 9, 16, 25}

func anharmonicityConstant(bondE float64) {
}

func frequencyHarmonic() {}

func harmonicOscillator(bondE float64, v float64) float64 {
	We := bondE / (v + 1/2)
	return We
}

func plot() {
	res, _ := fit.Curve1D(
		fit.Func1D{
			F:  poly,
			X:  xdata,
			Y:  ydata,
			Ps: []float64{1, 1},
		},
		nil, &optimize.NelderMead{},
	)
	{
		p := hplot.New()
		p.X.Label.Text = "f(x) = a + b*x*x"
		p.Y.Label.Text = "y-data"
		p.X.Min = -10
		p.X.Max = +10
		p.Y.Min = 0
		p.Y.Max = 220

		s := hplot.NewS2D(hplot.ZipXY(xdata, ydata))
		s.Color = color.RGBA{0, 0, 255, 255}
		p.Add(s)

		f := plotter.NewFunction(func(x float64) float64 {
			return poly(x, res.X)
		})
		f.Color = color.RGBA{255, 0, 0, 255}
		f.Samples = 1000
		p.Add(f)

		p.Add(plotter.NewGrid())

		err := p.Save(20*vg.Centimeter, -1, "image.png")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	We := harmonicOscillator()
}
