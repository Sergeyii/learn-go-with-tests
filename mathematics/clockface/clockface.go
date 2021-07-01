package clockface

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150

	svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg xmlns="http://www.w3.org/2000/svg"
width="100%"
height="100%"
viewBox="0 0 300 300"
version="1.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"></circle>`
	svgEnd = `</svg>`
)

type Point struct {
	X float64
	Y float64
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	//Text string `xml:",chardata"`
	Xmlns   string `xml:"xmlns,attr"`
	Width   string `xml:"width,attr"`
	Height  string `xml:"height,attr"`
	ViewBox string `xml:"viewBox,attr"`
	Version string `xml:"version,attr"`
	Circle  Circle `xml:"circle"`
	Line    []Line `xml:"line"`
}

type Circle struct {
	//Text string `xml:",chardata"`
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
	//Style string `xml:"style,attr"`
}

type Line struct {
	//Text string `xml:",chardata"`
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
	//Style string `xml:"style,attr"`
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} //scale
	p = Point{p.X, -p.Y}                                      //flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         //translate

	fmt.Fprintf(w, secondHandTag(p))
}

func secondHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;"></line>`, p.X, p.Y)
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} //scale
	p = Point{p.X, -p.Y}                                      //flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         //translate

	return p
}

func secondsInRadians(t time.Time) float64 {
	//TODO::здесь идёт потеря точности поэтому тест и не проходит
	//А в книге утверждается что всё норм.
	return (math.Pi / (30 / float64(t.Second())))
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func roughlyEqualFloat64(a, b float64) bool {
	//TODO::здесь идёт потеря точности поэтому тест и не проходит
	//А в книге утверждается что всё норм.
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func minuteHandPoint(t time.Time) Point {
	//TODO::здесь тест не проходит
	//А в книге утверждается что всё норм.
	return angleToPoint(minutesInRadians(t))
}
