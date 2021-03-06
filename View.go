package main

import (
	"strconv"

	"github.com/wesleycoakley/ral-api"
)

func View(s ral.Site, flags CommandSet, args []string) {
	var continuity string
	var topic, year int
	var err error

	if len(args) > 0 {
		continuity = args[0] }

	if len(args) > 1 && args[1] != "" {
		year, err = strconv.Atoi(args[1])
		if err != nil { return } }

	if len(args) > 2 && args[2] != "" {
		topic, err = strconv.Atoi(args[2])
		if err != nil { return } }

	format := Formats[*flags["format"].(*string)]
	nowrap := *flags["nowrap"].(*bool)
	wrap := *flags["wrap"].(*int)

	if nowrap && format == ral.FormatSimple {
		wrap = 0 }

	if topic > 0 {
		r, err := s.Replies(continuity, year, topic)
		if err != nil { panic(err) }
		r.Print(format, wrap)
	} else if year > 0 {
		t, err := s.Topics(continuity, year)
		if err != nil { panic(err) }
		t.Print(format, wrap)
	} else if continuity != "" {
		y, err := s.Years(continuity)
		if err != nil { panic(err) }
		y.Print(format)
	} else {
		c, err := s.Continuities()
		if err != nil { panic(err) }
		c.Print(format)
	} }
