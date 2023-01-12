package main

import (
	"flag"
	"rename/internal/rename"
)

func main() {
	var (
		path      string
		inReg     string
		outReg    string
		run       bool
		doFolder  bool
		recursive bool
	)

	flag.StringVar(&path, "p", ".", "Path to directory")
	flag.StringVar(&inReg, "in", "(.+)", "input regex")
	flag.StringVar(&outReg, "out", "$1", "output regex")
	flag.BoolVar(&run, "run", false, "run ")
	flag.BoolVar(&recursive, "r", false, "recursive search all directory")
	flag.BoolVar(&doFolder, "d", false, "incloud remane directory ")
	flag.Parse()
	rename.BatchRename(path, inReg, outReg, run, doFolder, recursive)
	if !run {
		println("*************\nTo execute rename process, use -run")
	}

}
