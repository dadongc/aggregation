package handler

import (
	"fmt"
	"math"
	"time"

	"github.com/dadongc/aggregation/convert"

	"github.com/bytedance/sonic"
	"github.com/dadongc/aggregation/dto"
	"github.com/dadongc/aggregation/tools"
	"github.com/dadongc/aggregation/train"
)

const (
	TrainNum = 10000
	TestNum  = 1000

	centerTxt = "labels/kmeans_center.txt"
)

type KMeansHandler struct {
	centers []*dto.ImageVector
}

func NewKMeansHandler() *KMeansHandler {
	return &KMeansHandler{}
}

func (h *KMeansHandler) Handle() {
	t := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(t).Seconds())
	}()
	// 生成聚类中心，存入tx
	trainVectors := convert.ConvertVector("./labels/train-labels-idx1-ubyte", "./labels/train-images-idx3-ubyte", TrainNum)
	h.genMergeCenter(trainVectors)

	// 测试准确率
	h.getAccuracy()
}

func (h *KMeansHandler) genMergeCenter(trainVectors []*dto.ImageVector) {
	finalCenter := train.GenMergeCenter(30, 30, trainVectors)
	images := make([]string, 0)
	for _, center := range finalCenter {
		centerStr, _ := sonic.MarshalString(center)
		images = append(images, centerStr)
	}
	tools.WriteTxt("labels/kmeans_center.txt", images)
}

func (h *KMeansHandler) parseMergeCenter() {
	centerData := tools.ReadTxt(centerTxt)
	centers := make([]*dto.ImageVector, 0)
	for _, data := range centerData {
		val := &dto.ImageVector{}
		if err := sonic.UnmarshalString(data, &val); err != nil {
			fmt.Println(err)
			continue
		}
		centers = append(centers, val)
		fmt.Println(val.Label)
	}
	h.centers = centers
}

func (h *KMeansHandler) getAccuracy() {
	// 从txt解析出聚类中心
	h.parseMergeCenter()

	testVectors := convert.ConvertVector("./labels/t10k-labels-idx1-ubyte", "./labels/t10k-images-idx3-ubyte", TestNum)
	correctNum := 0
	for _, vector := range testVectors {
		minDis, minIdx := math.MaxFloat64, -1
		for idx, center := range h.centers {
			dis := tools.Euclidean(center.Image, vector.Image)
			if dis < minDis {
				minDis = dis
				minIdx = idx
			}
		}
		if h.centers[minIdx].Label == vector.Label {
			correctNum++
		}
	}

	fmt.Printf("kmeans accuracy=%f\n", float64(correctNum)/float64(len(testVectors)))
}
