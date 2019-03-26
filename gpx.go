package ghht

import (
	"os"
	"sort"
	"time"

	"github.com/beevik/etree"
)

type point struct {
	lat     string
	lon     string
	elev    string
	time    string
	hr      string
	cadence string
}

type GpxFile struct {
	document *etree.Document
	points   []point
}

func (gpx *GpxFile) initXML() {
	gpx.document.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
}

func (gpx *GpxFile) instrument() *etree.Element {
	el := gpx.document.CreateElement("gpx")
	el.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	el.CreateAttr("xmlns:gpxdata", "http://www.cluetrust.com/XML/GPXDATA/1/0")
	el.CreateAttr("xmlns", "http://www.topografix.com/GPX/1/0")
	el.CreateAttr("xsi:schemaLocation", "http://www.topografix.com/GPX/1/0 http://www.topografix.com/GPX/1/0/gpx.xsd http://www.cluetrust.com/XML/GPXDATA/1/0 http://www.cluetrust.com/Schemas/gpxdata10.xsd")
	el.CreateAttr("version", "1.0")
	el.CreateAttr("creator", "GHHT")

	return el
}

func (gpx *GpxFile) fillData(root *etree.Element) {
	author := root.CreateElement("author")
	author.CreateText("GHHT")

	url := root.CreateElement("url")
	url.CreateText("GHHT")

	time := root.CreateElement("time")
	time.CreateText("GHHT")

	trk := root.CreateElement("trk")
	gpx.fillTrack(trk)
}

func (gpx *GpxFile) fillTrack(root *etree.Element) {
	name := root.CreateElement("name")
	name.CreateText("GHHT")

	trkseg := root.CreateElement("trkseg")
	gpx.fillSegment(trkseg)
}

func (gpx *GpxFile) fillSegment(root *etree.Element) {
	// name := root.CreateElement("name")
	// name.CreateText("GHHT")

	// trkseg := root.CreateElement("trkseg")
	for _, p := range gpx.points {
		gpx.fillPoint(root, p)
	}
}

var previousHr = "0.0"
var previousCadence = "0"

func (gpx *GpxFile) fillPoint(root *etree.Element, p point) {
	trkpt := root.CreateElement("trkpt")
	trkpt.CreateAttr("lat", p.lat)
	trkpt.CreateAttr("lon", p.lon)

	ele := trkpt.CreateElement("ele")
	ele.CreateText(p.elev)

	time := trkpt.CreateElement("time")
	time.CreateText(p.time)

	extensions := trkpt.CreateElement("extensions")

	if p.hr == "0.0" || p.hr == "" {
		p.hr = previousHr
	}
	previousHr = p.hr
	hr := extensions.CreateElement("gpxdata:hr")
	hr.CreateText(previousHr)

	if p.cadence == "" {
		p.cadence = previousCadence
	}
	previousCadence = p.cadence
	cadence := extensions.CreateElement("gpxdata:cadence")
	cadence.CreateText(previousCadence)
}

func GPXFromDump(dump *HuaweiTrack) *GpxFile {
	gpx := &GpxFile{}
	gpx.points = populatePoints(dump)
	gpx.document = etree.NewDocument()
	gpx.initXML()
	root := gpx.instrument()
	gpx.fillData(root)
	gpx.document.Indent(2)
	gpx.document.WriteTo(os.Stdout)

	return gpx
}

func populatePoints(dump *HuaweiTrack) []point {
	var pts []point
	for _, timestamp := range getSortedKeys(dump) {
		timedLine := (*dump)[timestamp]
		var pt point
		pt.time = getTimeStr(timestamp)
		for k, v := range timedLine {
			if k == "lbs" {
				pt.lat = v["lat"]
				pt.lon = v["lon"]
				pt.elev = v["alt"]
			} else if k == "h-r" {
				pt.hr = v["v"]
			}
		}
		if isValidPoint(pt) {
			pts = append(pts, pt)
		}
	}

	return pts
}

func isValidPoint(pt point) bool {
	// Clear data:
	/*
			try:
		        for line in data['hr']:
		            # Heart-rate is too low/high (type is xsd:unsignedbyte)
		            if line[5] < 1 or line[5] > 254:
		                data['hr'].remove(line)

		        for line in data['cad']:
		            # Cadence is too low/high (type is xsd:unsignedbyte)
		            if line[6] < 0 or line[6] > 254:
		                data['cad'].remove(line)

		        for line in data['alti']:
		            # Altitude is too low/high (dead sea/everest)
		            if line[4] < 1000 or line[4] > 10000:
		                data['alti'].remove(line)
	*/
	if pt.lat == "90.0" || pt.lat == "" || pt.lon == "-80.0" || pt.lon == "" {
		return false
	}
	// hr, err := strconv.Atoi(pt.hr)
	// if err != nil || hr < 0 || hr > 254 {
	// 	return false
	// }
	return true
}

func getTimeStr(timestamp int) string {
	ts := time.Unix(int64(timestamp), 0)
	return ts.Format(time.RFC3339)
}

func getSortedKeys(dump *HuaweiTrack) []int {
	keys := []int{}
	for k := range *dump {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
