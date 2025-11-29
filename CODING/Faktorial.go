package main

import (
	"fmt"
	"image/color"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// --- 1. Fungsi Iteratif ---
func factorialIterative(n int64) int64 {
	var result int64 = 1
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return result
}

// --- 2. Fungsi Rekursif ---
func factorialRecursive(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func measureTime(n int64, f func(int64) int64) float64 {
	start := time.Now()
	f(n)
	return float64(time.Since(start).Nanoseconds()) / 1000000.0
}

func main() {
	inputs := []int64{3000, 8000, 31000, 49000, 56000, 131000, 211000}


	ptsIterative := make(plotter.XYs, len(inputs))
	ptsRecursive := make(plotter.XYs, len(inputs))

	fmt.Println("\n=== HASIL PENGUJIAN FAKTORIAL ===")
	fmt.Println("+--------+----------------------+----------------------+")
	fmt.Printf("| %-6s | %-20s | %-20s |\n", "   n", " Waktu Rekursif (ms)", " Waktu Iteratif (ms)")
	fmt.Println("+--------+----------------------+----------------------+")

	for i, n := range inputs {
		tIter := measureTime(n, factorialIterative)
		tRecur := measureTime(n, factorialRecursive)
		fmt.Printf("| %6d | %20.4f | %20.4f |\n", n, tRecur, tIter)
		fmt.Println("+--------+----------------------+----------------------+")
		ptsIterative[i].X = float64(n)
		ptsIterative[i].Y = tIter
		ptsRecursive[i].X = float64(n)
		ptsRecursive[i].Y = tRecur
	}

	p := plot.New()
	p.Title.Text = "Perbandingan Waktu: Iteratif vs Rekursif"
	p.X.Label.Text = "Jumlah Input (n)"
	p.Y.Label.Text = "Waktu Eksekusi (ms)"
	p.Add(plotter.NewGrid()) 

	
	lineIter, _ := plotter.NewLine(ptsIterative)
	lineIter.LineStyle.Width = vg.Points(2)
	lineIter.LineStyle.Color = color.RGBA{B: 255, A: 255}


	lineRecur, _ := plotter.NewLine(ptsRecursive)
	lineRecur.LineStyle.Width = vg.Points(2)
	lineRecur.LineStyle.Color = color.RGBA{R: 255, A: 255}
	lineRecur.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	p.Add(lineIter, lineRecur)
	p.Legend.Add("Iteratif (Cepat)", lineIter)
	p.Legend.Add("Rekursif (Lambat)", lineRecur)

	p.Legend.Top = true  
	p.Legend.Left = true 

	p.Legend.XOffs = vg.Points(15)
	p.Legend.YOffs = vg.Points(-15)
	
	outputFile := "Grafik_Tubes_AKA.png"
	if err := p.Save(8*vg.Inch, 5*vg.Inch, outputFile); err != nil {
		panic(err)
	}

	fmt.Printf(outputFile)
}
