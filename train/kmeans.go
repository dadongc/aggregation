package train

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/dadongc/aggregation/tools"

	"github.com/dadongc/aggregation/dto"
)

func genInitCenter(size int, data []*dto.ImageVector) []*dto.ImageVector {
	rand.Seed(time.Now().Unix())
	exist := map[int]bool{}
	centers := make([]*dto.ImageVector, size)
	for i := 0; i < size; i++ {
		for {
			idx := rand.Intn(len(data))
			if !exist[idx] {
				centers[i] = data[idx]
				exist[idx] = true
				break
			}
		}
	}
	return centers
}

func isSame(oldCenters, newCenters []*dto.ImageVector, dimension int) bool {
	sort.Slice(oldCenters, func(i, j int) bool {
		for k := 0; k < dimension; k++ {
			if oldCenters[i].Image[k] == oldCenters[j].Image[k] {
				continue
			}
			return oldCenters[i].Image[k] < oldCenters[j].Image[k]
		}
		return true
	})
	sort.Slice(newCenters, func(i, j int) bool {
		for k := 0; k < dimension; k++ {
			if newCenters[i].Image[k] == newCenters[j].Image[k] {
				continue
			}
			return newCenters[i].Image[k] < newCenters[j].Image[k]
		}
		return true
	})

	for idx, oldCenter := range oldCenters {
		newCenter := newCenters[idx]
		for i := 0; i < dimension; i++ {
			if oldCenter.Image[i] != newCenter.Image[i] {
				return false
			}
		}
	}
	return true
}

func GenMergeCenter(size, iterNum int, data []*dto.ImageVector) []*dto.ImageVector {
	if len(data) == 0 {
		return nil
	}
	dimension := len(data[0].Image)
	centers := genInitCenter(size, data)
	for i := 0; i < iterNum; i++ {
		groups := map[*dto.ImageVector][]*dto.ImageVector{}
		for _, center := range centers {
			groups[center] = append(groups[center], center)
		}
		for _, vector := range data {
			minDis := math.MaxFloat64
			minCenterIdx := -1
			for idx, center := range centers {
				dis := tools.Euclidean(vector.Image, center.Image)
				if minDis > dis {
					minDis = dis
					minCenterIdx = idx
				}
			}
			groups[centers[minCenterIdx]] = append(groups[centers[minCenterIdx]], vector)
		}
		newCenters := make([]*dto.ImageVector, 0)
		totalDis := 0.0
		for _, vectors := range groups {
			labelCnt := map[int]int{}
			avgCenter := &dto.ImageVector{
				Image: make([]float64, dimension),
			}
			for _, vector := range vectors {
				for idx, num := range vector.Image {
					avgCenter.Image[idx] += num
				}
				labelCnt[vector.Label]++
			}
			for idx := range avgCenter.Image {
				avgCenter.Image[idx] = avgCenter.Image[idx] / float64(len(vectors))
			}
			maxCnt, maxLabel := 0, -1
			for label, cnt := range labelCnt {
				if cnt > maxCnt {
					maxCnt = cnt
					maxLabel = label
				}
			}
			avgCenter.Label = maxLabel
			newCenters = append(newCenters, avgCenter)
			//fmt.Printf("centerLabel=%d avgCenterLabel=%d labelCnt=%+v\n ", center.Label, avgCenter.Label, labelCnt)
			// 计算样本距聚类中心距离
			for _, vector := range vectors {
				totalDis += tools.Euclidean(vector.Image, avgCenter.Image)
			}
		}
		fmt.Printf("idx=%d new center dis=%+v\n", i, totalDis)
		if isSame(centers, newCenters, dimension) {
			break
		}
		centers = newCenters
	}
	return centers
}
