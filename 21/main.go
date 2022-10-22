// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

type Dvier interface {
	InsertDvi()
}

type HdmiMonitor struct{}
type DviMonitor struct{}
type DviAdapter struct {
	hdmiMonitor *HdmiMonitor
}

func (hdmi *HdmiMonitor) InsertHdmi() { fmt.Println("HDMI Inserted") }
func (dvi *DviMonitor) InsertDvi()    { fmt.Println("DVI Inserted") }
func (dvi *DviAdapter) InsertDvi()    { dvi.hdmiMonitor.InsertHdmi() }

func main() {
	var dvi DviMonitor
	var hdmi HdmiMonitor
	adp := DviAdapter{&hdmi}

	dvier := []Dvier{&dvi, &adp}

	for _, dvi := range dvier {
		dvi.InsertDvi()
	}
}
