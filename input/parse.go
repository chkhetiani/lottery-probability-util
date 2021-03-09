package input

import "flag"

func Parse() *Data {
	var data Data

	flag.IntVar(&data.Total, "total", 42, "total number of possible balls")
	flag.IntVar(&data.Chosen, "chosen", 6, "number of chosen balls")
	flag.IntVar(&data.Generated, "generated", 6, "number of generated balls")

	flag.Parse()
	return &data
}