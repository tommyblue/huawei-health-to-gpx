package ghht

import (
	"os"
	"strings"

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

type gpxFile struct {
	document *etree.Document
	points   []point
}

func (gpx *gpxFile) initXML() {
	gpx.document.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
}

func (gpx *gpxFile) instrument() *etree.Element {
	el := gpx.document.CreateElement("gpx")
	el.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	el.CreateAttr("xmlns:gpxdata", "http://www.cluetrust.com/XML/GPXDATA/1/0")
	el.CreateAttr("xmlns", "http://www.topografix.com/GPX/1/0")
	el.CreateAttr("xsi:schemaLocation", "http://www.topografix.com/GPX/1/0 http://www.topografix.com/GPX/1/0/gpx.xsd http://www.cluetrust.com/XML/GPXDATA/1/0 http://www.cluetrust.com/Schemas/gpxdata10.xsd")
	el.CreateAttr("version", "1.0")
	el.CreateAttr("creator", "GHHT")

	return el
}

func (gpx *gpxFile) fillData(root *etree.Element) {
	author := root.CreateElement("author")
	author.CreateText("GHHT")

	url := root.CreateElement("url")
	url.CreateText("GHHT")

	time := root.CreateElement("time")
	time.CreateText("GHHT")

	trk := root.CreateElement("trk")
	gpx.fillTrack(trk)
}

func (gpx *gpxFile) fillTrack(root *etree.Element) {
	name := root.CreateElement("name")
	name.CreateText("GHHT")

	trkseg := root.CreateElement("trkseg")
	gpx.fillSegment(trkseg)
}

func (gpx *gpxFile) fillSegment(root *etree.Element) {
	name := root.CreateElement("name")
	name.CreateText("GHHT")

	trkseg := root.CreateElement("trkseg")
	for _, p := range gpx.points {
		gpx.fillPoint(trkseg, p)
	}
}

func (gpx *gpxFile) fillPoint(root *etree.Element, p point) {
	trkpt := root.CreateElement("trkpt")
	trkpt.CreateAttr("lat", p.lat)
	trkpt.CreateAttr("lon", p.lon)

	ele := trkpt.CreateElement("ele")
	ele.CreateText(p.elev)

	time := trkpt.CreateElement("time")
	time.CreateText(p.time)

	extensions := trkpt.CreateElement("extensions")

	hr := extensions.CreateElement("gpxdata:hr")
	hr.CreateText(p.hr)

	cadence := extensions.CreateElement("gpxdata:cadence")
	cadence.CreateText(p.cadence)
}

func GPXFromDump(dump string) {
	gpx := &gpxFile{}
	// Fake data
	// gpx.points = []point{
	// 	{"asd", "das", "fsfd", "asdas", "werw", "asdas"},
	// }
	gpx.document = etree.NewDocument()
	gpx.initXML()
	root := gpx.instrument()
	gpx.fillData(root)
	gpx.document.Indent(2)
	gpx.document.WriteTo(os.Stdout)

	// r := map[string][]map[string]string{}
	// scanner := bufio.NewScanner(strings.NewReader(dump))
	// for scanner.Scan() {
	// 	tp, m := parseLine(scanner.Text())
	// 	_, ok := r[tp]
	// 	if ok {
	// 		r[tp] = append(r[tp], m)
	// 	} else {
	// 		r[tp] = []map[string]string{m}
	// 	}
	// }

	// keys := make([]string, len(r))
	// i := 0
	// for k := range r {
	// 	keys[i] = k
	// 	i++
	// }
	// fmt.Println(keys)
	// /* [
	// 	s-r (k, v sempre 0) cadence
	// 	rs (k, v)
	// 	lbs (alt, t, lon, lat, k) location
	// 	p-m (k, v)
	// 	b-p-m (k, v)
	// 	h-r (k, v) heart-rate
	// ] */
	// fmt.Println(r["h-r"])
}

func parseLine(line string) (string, map[string]string) {
	m := map[string]string{}
	var tp string
	for _, s := range strings.Split(line, ";") {
		if s != "" {
			r := strings.Split(s, "=")
			if r[0] == "tp" {
				tp = r[1]
			} else {
				m[r[0]] = r[1]
			}
		}
	}

	return tp, m
}
