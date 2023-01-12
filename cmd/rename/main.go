package main

import (
	"flag"
	"rename/internal/rename"
)

func main() {
	var (
		path   string
		inReg  string
		outReg string
		run    bool
	)

	flag.StringVar(&path, "p", ".", "Path to dir")
	flag.StringVar(&inReg, "in", "(.+)", "input regex")
	flag.StringVar(&outReg, "out", "$1", "output regex")
	flag.BoolVar(&run, "run", false, "run ")
	flag.Parse()
	rename.BatchRename(path, inReg, outReg, run)
	if !run {
		println("*************\nTo execute rename process, use -run")
	}

}
