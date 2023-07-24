package train

import (
	"fmt"
	"github.com/dadongc/aggregation/convert"
	"github.com/dadongc/aggregation/dto"
	"github.com/dadongc/aggregation/tools"
	"time"
)

const (
	KNN_K    = 10
	TrainNum = 10000
	TestNum  = 1000
)

func KNN() {
	t := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(t).Seconds())
	}()
	trainVectors := convert.ConvertVector("./labels/train-labels-idx1-ubyte", "./labels/train-images-idx3-ubyte", TrainNum)
	testVectors := convert.ConvertVector("./labels/t10k-labels-idx1-ubyte", "./labels/t10k-images-idx3-ubyte", TestNum)
	correctNum := float64(0)
	for _, v1 := range testVectors {
		cosRes := make([]dto.CosineLabel, 0)
		for _, v2 := range trainVectors {
			cosine := tools.Cosine(v1.Image, v2.Image)
			cosRes = append(cosRes, dto.CosineLabel{
				Cosine: cosine,
				Label:  v2.Label,
			})
		}
		maxCosine := tools.GetMaxNumsOfCosine(cosRes, KNN_K)
		if isKNNCorrect(maxCosine, v1.Label) {
			correctNum++
		}
	}
	fmt.Println(correctNum / TestNum)
}

func isKNNCorrect(maxCosine []dto.CosineLabel, realLabel int) bool {
	labelCount := map[int]int{}
	for _, c := range maxCosine {
		labelCount[c.Label]++
	}
	maxLabel, maxCnt := -1, 0
	for label, cnt := range labelCount {
		if cnt > maxCnt {
			maxLabel = label
		}
	}
	return maxLabel == realLabel
}
