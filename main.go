package main

import "github.com/dadongc/aggregation/handler"

func main() {
	//train.KNN()
	handler.NewKMeansHandler().Handle()
}
