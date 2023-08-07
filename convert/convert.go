package convert

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/dadongc/aggregation/dto"
)

func ConvertVector(labelFile, imageFile string, n int) []*dto.ImageVector {
	result := parseImage(labelFile, imageFile, n)
	vectors := make([]*dto.ImageVector, 0)
	for _, r := range result {
		vector := make([]float64, 0)
		for _, b := range r.Image {
			if b > 127 {
				vector = append(vector, 1)
			} else {
				vector = append(vector, 0)
			}
		}
		vectors = append(vectors, &dto.ImageVector{
			Image: vector,
			Label: int(r.Label),
		})
	}
	return vectors
}

func ConvertImage(labelFile, imageFile string, n int) {
	result := parseImage(labelFile, imageFile, n)
	rect := image.Rect(0, 0, 28, 28)
	rgba := image.NewRGBA(rect)
	for dy := 0; dy < 28; dy++ {
		for dx := 0; dx < 28; dx++ {
			rgba.Set(dy, dx, color.Gray{result[0].Image[dy+dx*28]})
		}
	}
	fIm, _ := os.Create("test.png")
	png.Encode(fIm, rgba)
}

func parseImage(labelFile, imageFile string, n int) []dto.ImageLabel {
	// 打开文件
	label, _ := os.ReadFile(labelFile)
	image, _ := os.ReadFile(imageFile)
	// 过滤初始信息
	label = label[8:]
	image = image[16:]
	// 打包image
	images := make([]dto.ImageLabel, n)
	for i := 0; i < n; i++ {
		images[i] = dto.ImageLabel{
			Image: image[:784],
			Label: label[0],
		}
		label = label[1:]
		image = image[784:]
	}
	return images
}
