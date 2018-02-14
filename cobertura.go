package main

import (
	"encoding/xml"
)

type Coverage struct {
	XMLName         xml.Name   `xml:"coverage"`
	LineRate        float32    `xml:"line-rate,attr,omitempty"`
	BranchRate      float32    `xml:"branch-rate,attr,omitempty"`
	Version         string     `xml:"version,attr"`
	Timestamp       int64      `xml:"timestamp,attr"`
	LinesCovered    int64      `xml:"lines-covered,attr,omitempty"`
	LinesValid      int64      `xml:"lines-valid,attr,omitempty"`
	BranchesCovered int64      `xml:"branches-covered,attr,omitempty"`
	BranchesValid   int64      `xml:"branches-valid,attr,omitempty"`
	Complexity      float32    `xml:"complexity,attr,omitempty"`
	Sources         []*Source  `xml:"sources>source"`
	Packages        []*Package `xml:"packages>package"`
}

type Source struct {
	Path string `xml:",chardata"`
}

type Package struct {
	Name       string   `xml:"name,attr"`
	LineRate   float32  `xml:"line-rate,attr,omitempty"`
	BranchRate float32  `xml:"branch-rate,attr,omitempty"`
	Complexity float32  `xml:"complexity,attr,omitempty"`
	Classes    []*Class `xml:"classes>class"`
}

type Class struct {
	Name       string    `xml:"name,attr"`
	Filename   string    `xml:"filename,attr"`
	LineRate   float32   `xml:"line-rate,attr,omitempty"`
	BranchRate float32   `xml:"branch-rate,attr,omitempty"`
	Complexity float32   `xml:"complexity,attr,omitempty"`
	Methods    []*Method `xml:"methods>method"`
	Lines      []*Line   `xml:"lines>line"`
}

type Method struct {
	Name       string  `xml:"name,attr"`
	Signature  string  `xml:"signature,attr"`
	LineRate   float32 `xml:"line-rate,attr,omitempty"`
	BranchRate float32 `xml:"branch-rate,attr,omitempty"`
	Complexity float32 `xml:"complexity,attr,omitempty"`
	Lines      []*Line `xml:"lines>line"`
}

type Line struct {
	Number int   `xml:"number,attr"`
	Hits   int64 `xml:"hits,attr"`
}
