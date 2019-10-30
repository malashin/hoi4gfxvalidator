package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
	"strings"
	"time"
)

type imageJSON struct {
	Image Image
}

type Image struct {
	Name              string `json:"name"`
	BaseName          string `json:"baseName"`
	Format            string `json:"format"`
	FormatDescription string `json:"formatDescription"`
	Class             string `json:"class"`
	Geometry          struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		X      int `json:"x"`
		Y      int `json:"y"`
	} `json:"geometry"`
	Resolution struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"resolution"`
	PrintSize struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"printSize"`
	Units        string `json:"units"`
	Type         string `json:"type"`
	BaseType     string `json:"baseType"`
	Endianess    string `json:"endianess"`
	Colorspace   string `json:"colorspace"`
	Depth        int    `json:"depth"`
	BaseDepth    int    `json:"baseDepth"`
	ChannelDepth struct {
		Red   int `json:"red"`
		Green int `json:"green"`
		Blue  int `json:"blue"`
	} `json:"channelDepth"`
	Pixels          int `json:"pixels"`
	ImageStatistics struct {
		Overall struct {
			Min               int     `json:"min"`
			Max               int     `json:"max"`
			Mean              float64 `json:"mean"`
			StandardDeviation float64 `json:"standardDeviation"`
			Kurtosis          float64 `json:"kurtosis"`
			Skewness          float64 `json:"skewness"`
			Entropy           float64 `json:"entropy"`
		} `json:"Overall"`
	} `json:"imageStatistics"`
	ChannelStatistics struct {
		Red struct {
			Min               int     `json:"min"`
			Max               int     `json:"max"`
			Mean              float64 `json:"mean"`
			StandardDeviation float64 `json:"standardDeviation"`
			Kurtosis          float64 `json:"kurtosis"`
			Skewness          float64 `json:"skewness"`
			Entropy           float64 `json:"entropy"`
		} `json:"Red"`
		Green struct {
			Min               int     `json:"min"`
			Max               int     `json:"max"`
			Mean              float64 `json:"mean"`
			StandardDeviation float64 `json:"standardDeviation"`
			Kurtosis          float64 `json:"kurtosis"`
			Skewness          float64 `json:"skewness"`
			Entropy           float64 `json:"entropy"`
		} `json:"Green"`
		Blue struct {
			Min               int     `json:"min"`
			Max               int     `json:"max"`
			Mean              float64 `json:"mean"`
			StandardDeviation float64 `json:"standardDeviation"`
			Kurtosis          float64 `json:"kurtosis"`
			Skewness          float64 `json:"skewness"`
			Entropy           float64 `json:"entropy"`
		} `json:"Blue"`
	} `json:"channelStatistics"`
	Alpha           string   `json:"alpha"`
	ColormapEntries int      `json:"colormapEntries"`
	Colormap        []string `json:"colormap"`
	RenderingIntent string   `json:"renderingIntent"`
	Gamma           float64  `json:"gamma"`
	Chromaticity    struct {
		RedPrimary struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"redPrimary"`
		GreenPrimary struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"greenPrimary"`
		BluePrimary struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"bluePrimary"`
		WhitePrimary struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"whitePrimary"`
	} `json:"chromaticity"`
	MatteColor       string `json:"matteColor"`
	BackgroundColor  string `json:"backgroundColor"`
	BorderColor      string `json:"borderColor"`
	TransparentColor string `json:"transparentColor"`
	Interlace        string `json:"interlace"`
	Intensity        string `json:"intensity"`
	Compose          string `json:"compose"`
	PageGeometry     struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		X      int `json:"x"`
		Y      int `json:"y"`
	} `json:"pageGeometry"`
	Dispose     string `json:"dispose"`
	Iterations  int    `json:"iterations"`
	Scene       int    `json:"scene"`
	Scenes      int    `json:"scenes"`
	Compression string `json:"compression"`
	Orientation string `json:"orientation"`
	Properties  struct {
		DateCreate time.Time `json:"date:create"`
		DateModify time.Time `json:"date:modify"`
		Signature  string    `json:"signature"`
	} `json:"properties"`
	Tainted         bool   `json:"tainted"`
	Filesize        string `json:"filesize"`
	NumberPixels    string `json:"numberPixels"`
	PixelsPerSecond string `json:"pixelsPerSecond"`
	UserTime        string `json:"userTime"`
	ElapsedTime     string `json:"elapsedTime"`
	Version         string `json:"version"`
}

func (i Image) String() string {
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func Identify(filePath string) (i Image, err error) {
	c := exec.Command("magick",
		"convert",
		filePath,
		"json:",
	)
	var o bytes.Buffer
	var e bytes.Buffer
	c.Stdout = &o
	c.Stderr = &e
	err = c.Run()
	if err != nil {
		return i, errors.New(err.Error() + "\n" + string(e.Bytes()))
	}

	s := strings.TrimPrefix(o.String(), "[")
	s = strings.TrimSuffix(s, "]")
	s = strings.TrimSpace(s)

	var j *imageJSON
	json.Unmarshal([]byte(s), &j)
	i = j.Image

	return i, nil
}
