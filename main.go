package main

import (
	"log"
	"math"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Title = "SysTerm"
	p.Text = "System Information for your taste =)"
	p.SetRect(0, 0, 50, 5)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan

	updateParagraph := func(count int) {
		if count%2 == 0 {
			p.TextStyle.Fg = ui.ColorRed
		} else {
			p.TextStyle.Fg = ui.ColorWhite
		}
	}

	listData := []string{
		"[0] Asus H97 Pro",
		"[1] Ram Corsair DDR4",
		"[2] Intel 7600 Box",
		"[3] NVidia geForce 970",
		"[4] Samsung SSD 500 evo",
		"[5] Western Digital T1",
		"[6] TSSTcorp CDDVDW",
		"[7] ATA Samsung SSD 850",
	}

	l := widgets.NewList()
	l.Title = "Hardware"
	l.Rows = listData
	l.SetRect(0, 5, 25, 12)
	l.TextStyle.Fg = ui.ColorYellow

	g := widgets.NewGauge()
	g.Title = "Network Overload"
	g.Percent = 50
	g.SetRect(0, 12, 50, 15)
	g.BarColor = ui.ColorRed
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorCyan

	sparklineData := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}

	sl := widgets.NewSparkline()
	sl.Title = "SSD:"
	sl.Data = sparklineData
	sl.LineColor = ui.ColorCyan
	sl.TitleStyle.Fg = ui.ColorWhite

	sl2 := widgets.NewSparkline()
	sl2.Title = "HDD:"
	sl2.Data = sparklineData
	sl2.TitleStyle.Fg = ui.ColorWhite
	sl2.LineColor = ui.ColorRed

	slg := widgets.NewSparklineGroup(sl, sl2)
	slg.Title = "Node Space"
	slg.SetRect(25, 5, 50, 12)

	sinData := (func() []float64 {
		n := 220
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	})()

	lc := widgets.NewPlot()
	lc.Title = "CPU Usage"
	lc.Data = make([][]float64, 1)
	lc.Data[0] = sinData
	lc.SetRect(0, 15, 50, 25)
	lc.AxesColor = ui.ColorWhite
	lc.LineColors[0] = ui.ColorRed
	lc.Marker = widgets.MarkerDot

	barchartData := []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}

	bc := widgets.NewBarChart()
	bc.Title = "Threads Breakdown"
	bc.SetRect(50, 0, 90, 10)
	bc.Labels = []string{"t0", "t1", "t2", "t3", "t4", "t5","t6","t7"}
	bc.BarColors[0] = ui.ColorGreen
	bc.NumStyles[0] = ui.NewStyle(ui.ColorBlack)

	lc2 := widgets.NewPlot()
	lc2.Title = "Memory Usage"
	lc2.Data = make([][]float64, 1)
	lc2.Data[0] = sinData
	lc2.SetRect(50, 15, 75, 25)
	lc2.AxesColor = ui.ColorWhite
	lc2.LineColors[0] = ui.ColorYellow

	draw := func(count int) {
		g.Percent = count % 101
		l.Rows = listData[count%9:]
		slg.Sparklines[0].Data = sparklineData[:30+count%50]
		slg.Sparklines[1].Data = sparklineData[:35+count%50]
		lc.Data[0] = sinData[count/2%220:]
		lc2.Data[0] = sinData[2*count%220:]
		bc.Data = barchartData[count/2%10:]

		ui.Render(p, l, g, slg, lc, bc, lc2)
	}

	tickerCount := 1
	draw(tickerCount)
	tickerCount++
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			updateParagraph(tickerCount)
			draw(tickerCount)
			tickerCount++
		}
	}
}